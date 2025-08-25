//
// Copyright (c) 2022 Cloudflare, Inc.
//
// Licensed under Apache 2.0 license found in the LICENSE file
// or at http://www.apache.org/licenses/LICENSE-2.0
//

package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"github.com/IBM/sarama"
	prom_bmp "github.com/bio-routing/bio-rd/metrics/bmp/adapter/prom"
	"github.com/bio-routing/bio-rd/protocols/bgp/server"
	"github.com/cloudflare/certinel/fswatcher"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

// Command line parameters
var (
	bmpListenAddr    = flag.String("bmp.listen.addr", ":5000", "BMP listening address")
	healthListenAddr = flag.String("health.listen.addr", ":8080", "Prometheus/health/readiness check listening address/port")

	// Kafka producer configuration
	kafkaCluster = flag.String("kafka.cluster", "", "Kafka cluster FQDN:port to talk to")
	kafkaTopic   = flag.String("kafka.topic", "", "Kafka topic to write messages to")

	// Kafka TLS/mTLS configuration
	KafkaEnableTLS    = flag.Bool("kafka.net.tls", false, "Enable Kafka TLS")
	KafkaEnablemTLS   = flag.Bool("kafka.net.mtls", false, "Enable Kafka mTLS authentication")
	KafkaTLSRootPath  = flag.String("kafka.net.tls.root", "", "Path of Kafka TLS root CA")
	KafkamTLSCertPath = flag.String("kafka.net.mtls.cert", "", "Path of Kafka mTLS certificate")
	KafkamTLSKeyPath  = flag.String("kafka.net.mtls.key", "", "Path of Kafka mTLS key")
)

// Prometheus metrics
var (
	messagesProcessed     = prometheus.NewCounter(prometheus.CounterOpts{Name: "messages_processed", Help: "BMP messages processed by BioBMP"})
	messagesMarshalFailed = prometheus.NewCounter(prometheus.CounterOpts{Name: "messages_marshal_failed", Help: "BMP messages failed to marshal to proto"})
	messagesSendFailed    = prometheus.NewCounter(prometheus.CounterOpts{Name: "messages_send_failed", Help: "BMP messages failed to send to kafka"})
)

// Health check
var (
	ready   = int32(0)
	healthy = int32(0)
)

func init() {
	prometheus.MustRegister(messagesProcessed)
	prometheus.MustRegister(messagesMarshalFailed)
	prometheus.MustRegister(messagesSendFailed)
}

func lookupAddrs(hostport string) ([]string, error) {
	hostname, port, err := net.SplitHostPort(hostport)
	if err != nil {
		return nil, fmt.Errorf("net.SplitHostPort failed: %v", err)
	}

	if !strings.HasPrefix(hostname, "_") {
		return []string{hostport}, nil

	}

	_, srvAddrs, err := net.LookupSRV("", "", hostname)
	if err != nil {
		return nil, fmt.Errorf("lookup or SRV record '%q' failed: %v", hostname, err)
	}

	addrs := make([]string, 0)
	for _, s := range srvAddrs {
		addrs = append(addrs, net.JoinHostPort(s.Target, port))
	}

	return addrs, nil
}

func main() {
	flag.Parse()

	log.Info("BioBMP BMP receiver bbmp2kafka starting...")

	if *kafkaCluster == "" {
		log.Fatal("No kafka.cluster service address given.")
	}

	if *kafkaTopic == "" {
		log.Fatal("No kafka.topic given.")
	}

	// Resolve server addresses
	kafkaSrvs, err := lookupAddrs(*kafkaCluster)
	if err != nil {
		log.Fatalf("failed to lookup addresses: %+v", err)
	}

	if len(kafkaSrvs) == 0 {
		log.Fatalf("no kafka server specified or srv record provided empty result")
	}

	// Health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		msg := "OK"
		if atomic.LoadInt32(&healthy) == 0 {
			msg = "NOK"
			w.WriteHeader(http.StatusServiceUnavailable)
		}

		w.Write([]byte(msg))
	})
	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		msg := "OK"
		if atomic.LoadInt32(&ready) == 0 {
			msg = "NOK"
			w.WriteHeader(http.StatusServiceUnavailable)
		}

		w.Write([]byte(msg))
	})

	// Set up Prometheus + HTTP listener
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(*healthListenAddr, nil)
		if err != nil {
			log.Fatalf("http.ListenAndServe failed: %v", err)
		}
	}()

	// Set up kafka connection
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Producer.Return.Errors = true
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}
	if *KafkaEnableTLS {
		kafkaConfig.Net.TLS.Enable = true

		caCert, err := os.ReadFile(*KafkaTLSRootPath)
		if err != nil {
			log.Fatal(err)
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		tlsConfig.RootCAs = caCertPool
	}
	if *KafkaEnablemTLS {
		certinel, err := fswatcher.New(*KafkamTLSCertPath, *KafkamTLSKeyPath)
		if err != nil {
			log.WithError(err).Fatal("failed to create certinel")
		}
		go func() {
			if err := certinel.Start(context.Background()); err != nil {
				log.WithError(err).Fatal("error on certinel loop")
			}
		}()
		tlsConfig.GetClientCertificate = certinel.GetClientCertificate
	}

	kafkaConfig.Net.TLS.Config = tlsConfig

	ssp, err := sarama.NewSyncProducer(kafkaSrvs, kafkaConfig)
	if err != nil {
		log.Fatalf("failed to set up SyncProducer: %v", err)
	}
	defer ssp.Close()

	// Set up BMP receiver
	cfg := server.BMPReceiverConfig{
		KeepalivePeriod:  time.Minute,
		AcceptAny:        true,
		IgnorePrePolicy:  false,
		IgnorePostPolicy: false,
	}

	f := &adjRIBInFactory{
		producer:    ssp,
		kafkaTopic:  *kafkaTopic,
		tokenBucket: newTokenBucket(10, time.Second),
	}
	defer f.tokenBucket.stop()

	b := server.NewBMPReceiverWithAdjRIBInFactory(cfg, f)
	prometheus.MustRegister(prom_bmp.NewCollector(b))

	atomic.StoreInt32(&ready, 1)
	atomic.StoreInt32(&healthy, 1)

	// Start BMP receiver
	err = b.Listen(*bmpListenAddr)
	if err != nil {
		log.Fatalf("BMP receiver listen failed: %v", err)
	}

	err = b.Serve()
	if err != nil {
		log.WithError(err).Fatal("error while serving BMP connections")
	}
}
