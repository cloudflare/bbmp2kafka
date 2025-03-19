//
// Copyright (c) 2022 Cloudflare, Inc.
//
// Licensed under Apache 2.0 license found in the LICENSE file
// or at http://www.apache.org/licenses/LICENSE-2.0
//

package main

import (
	"sync/atomic"

	"github.com/cloudflare/bbmp2kafka/protos/bbmp"

	"github.com/IBM/sarama"
	"google.golang.org/protobuf/proto"

	"github.com/bio-routing/bio-rd/net"
	"github.com/bio-routing/bio-rd/route"
	"github.com/bio-routing/bio-rd/routingtable"
	"github.com/bio-routing/bio-rd/routingtable/filter"

	log "github.com/sirupsen/logrus"
)

type adjRIBInFactory struct {
	producer    sarama.SyncProducer
	kafkaTopic  string
	tokenBucket *tokenBucket
}

type adjRIBin struct {
	sessionAttrs routingtable.SessionAttrs
	producer     sarama.SyncProducer
	kafkaTopic   string
	tokenBucket  *tokenBucket
}

func (a *adjRIBInFactory) New(exportFilterChain filter.Chain, contributingASNs *routingtable.ContributingASNs, sessionAttrs routingtable.SessionAttrs) routingtable.AdjRIBIn {
	return &adjRIBin{
		sessionAttrs: sessionAttrs,
		producer:     a.producer,
		kafkaTopic:   a.kafkaTopic,
		tokenBucket:  a.tokenBucket,
	}
}

func (a *adjRIBin) createBBMPUnicastMonitoringMessage(pfx *net.Prefix, path *route.Path, announcement bool) []byte {
	bbmpMsg := bbmp.BBMPUnicastMonitoringMessage{
		RouterIp:      a.sessionAttrs.RouterIP.ToProto(),
		LocalBpgIp:    a.sessionAttrs.LocalIP.ToProto(),
		NeighborBgpIp: a.sessionAttrs.PeerIP.ToProto(),
		LocalAs:       a.sessionAttrs.LocalASN,
		RemoteAs:      a.sessionAttrs.PeerASN,
		Announcement:  announcement,
		BgpPath:       path.BGPPath.ToProto(),
		Pfx:           pfx.ToProto(),
		Timestamp:     path.LTime,
	}

	msg := bbmp.BBMPMessage{
		MessageType:                  bbmp.BBMPMessage_RouteMonitoringMessage,
		BbmpUnicastMonitoringMessage: &bbmpMsg,
	}

	msgBytes, err := proto.Marshal(&msg)
	if err != nil {
		messagesMarshalFailed.Inc()
		if a.tokenBucket.getToken() {
			log.Errorf("failed to marshal BBMPMessage: %v", err)
		}

		return nil
	}

	return msgBytes
}

func (a *adjRIBin) sendMessage(msg []byte) {
	_, _, err := a.producer.SendMessage(&sarama.ProducerMessage{
		Topic: a.kafkaTopic,
		Value: sarama.ByteEncoder(msg),
	})
	if err != nil {
		atomic.StoreInt32(&healthy, 0)

		messagesSendFailed.Inc()
		if a.tokenBucket.getToken() {
			log.Errorf("could not send message: %v", err)
		}

		return
	}

	atomic.StoreInt32(&healthy, 1)
}

func (a *adjRIBin) AddPath(pfx *net.Prefix, path *route.Path) error {
	messagesProcessed.Inc()
	msg := a.createBBMPUnicastMonitoringMessage(pfx, path, true)
	_ = msg

	a.sendMessage(msg)

	return nil
}

func (a *adjRIBin) RemovePath(pfx *net.Prefix, path *route.Path) bool {
	messagesProcessed.Inc()
	msg := a.createBBMPUnicastMonitoringMessage(pfx, path, false)
	_ = msg

	a.sendMessage(msg)

	return true
}

/*
 * Only here to fulfill the Interface
 */

func (a *adjRIBin) ReplaceFilterChain(filter.Chain) {}

func (a *adjRIBin) Dump() []*route.Route {
	return nil
}

func (a *adjRIBin) Register(client routingtable.RouteTableClient) {}

func (a *adjRIBin) Unregister(client routingtable.RouteTableClient) {}

func (a *adjRIBin) Flush() {}

func (a *adjRIBin) RouteCount() int64 {
	return 0
}

func (a *adjRIBin) ClientCount() uint64 {
	return 0
}
