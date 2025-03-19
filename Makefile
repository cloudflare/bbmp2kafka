.PHONY: clean
clean:
	go clean

.PHONY: build
build:
	go build -mod=vendor

.PHONY: vet
vet:
	go vet -mod=vendor github.com/cloudflare/bbmp2kafka/...

.PHONY: test
test: vet
	go test -v -mod=vendor -race  ./...

proto-go:
	protoc -Ivendor/github.com/bio-routing/bio-rd/ --proto_path=protos --go_out=protos/ --go_opt=paths=source_relative bbmp/bbmp.proto

