#!/bin/sh

printf  "⚡ Enabling  GO111MODULE  support \n"
export GO111MODULE=on
printf "GO111MODULE=$GO111MODULE \n\n"
echo "GOPATH is $GOPATH"

go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
export PATH=$PATH:$GOPATH/bin
rm -rf api/swagger/
mkdir api/swagger/
for filename in $(find ./api/proto -name '*.proto' );do \
    echo "Generating for ${filename}"
    protoc  \
    -I. \
    -I$(go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway) \
    -I$(go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway)/third_party/googleapis \
    -I${GOPATH}/src \
    --go_out=plugins=grpc:${GOPATH}/src \
    --grpc-gateway_out=logtostderr=true:${GOPATH}/src \
    --swagger_out=logtostderr=true:api/swagger/ \
${filename};done

