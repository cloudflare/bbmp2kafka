# bbmp2kafka

`bbmk2pafka` is an Open Source BGP Monitoring Protocol (BMP, RFC7854) receiver based on the [Bio-Routing BGP/BMP library](https://github.com/bio-routing/bio-rd),
designed to forward all BMP Monitoring messages, so specifically BGP UPDATEs and WITHDRAWs, onto a [Kafka](https://kafka.apache.org/) stream as [protobuf](https://developers.google.com/protocol-buffers) messages.

It is intended to allow replacing OpenBMP in existing setups.

## Building

You need at least Go version 1.19 to build `bbmp2kafka`.

Running `make` or `make build` will build the daemon.

Running `make proto-go` will regenerate the protobuf definition for the Go language.

## Running

`bbmp2kafka` has a small number of command line parameters, with `-kafka.cluster` and `-kafka.topic` being required:

```bash
$ ./bbmp2kafka --help
Usage of ./bbmp2kafka:
  -bmp.listen.addr string
    	BMP listening address (default ":5000")
  -health.listen.addr string
    	Prometheus/health/readiness check listening address/port (default ":8080")
  -kafka.cluster string
    	Kafka cluster FQDN:port to talk to
  -kafka.topic string
    	Kafka topic to write messages to
```

### Testing

For testing you can for example run `bbmp2kafka` with a locally running Zookeper + Kafka setup.

You can download the latest version of [Apache Kafka here](https://kafka.apache.org/downloads), extract it, and start ZooKeper and Kafka.

```bash
wget https://downloads.apache.org/kafka/3.3.1/kafka_2.13-3.3.1.tgz
tar xf kafka_2.13-3.3.1.tgz
cd kafka_2.13-3.3.1
bin/zookeeper-server-start.sh config/zookeeper.properties
bin/kafka-server-start.sh config/server.properties
```

Then start `bbmp2kafka` to connect to your local Kafka instance using the topic `biobmp`

```bash
./bbmp2kafka -kafka.cluster localhost:9092 -kafka.topic biobmp
```

Now you can connect a router to port `5000` of your machine or use the example BMP file provided in the `example_data` directory and feed it into `bbmp2kafka`

```bash
nc localhost 5000 < example_data/bmp.raw 
```

## Consuming data

`bbmp2kafka` will send the BGP UPDATEs and WITHDRAWs onto the Kafka topic into ProtoBuf format marshalled.

To consume the data you need to tap into the Kafka stream and unmarshal the data like

```go
bbmpMsg := &bbmp.BBMPMessage{}
err := proto.Unmarshal(data, bbmpMsg)
if err != nil {
    return nil, fmt.Errorf("unable to unmarshal BioBMP message: %v", err)
}
```
