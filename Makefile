#!/usr/bin/make

build:
	go build

proto-go:
	protoc -Ivendor/github.com/bio-routing/bio-rd/ --proto_path=protos --go_out=protos/ --go_opt=paths=source_relative bbmp/bbmp.proto

