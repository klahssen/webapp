protogo:
	protoc *.proto -I. -I=${GOPATH}/src/github.com/klahssen/protobuf \
	-I=${GOPATH}/src --go_out=plugins=grpc:${GOPATH}/src\
	&& protoc-go-tags --dir=../domain