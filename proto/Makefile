protogo:
	protoc *.proto -I. -I=${GOPATH}/src/github.com/klahssen/protobuf \
	-I=${GOPATH}/src --go_out=plugins=grpc:${GOPATH}/src\
	&& protoc-go-tags --dir=../pkg/domain

gogoproto:
	protoc *.proto -I. -I=${GOPATH}/src/github.com/klahssen/protobuf \
	-I=${GOPATH}/src -I=$(GOPATH)/src/github.com/gogo/protobuf/protobuf --gogofaster_out=plugins=grpc:${GOPATH}/src