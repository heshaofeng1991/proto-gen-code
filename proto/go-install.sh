#!/bin/sh

go install github.com/gogo/protobuf@v1.3.2
go install github.com/mwitkow/go-proto-validators@v0.3.2
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators@latest
go install github.com/heshaofeng1991/protoc-gen-go-http@latest
go install github.com/gogo/protobuf/protoc-gen-gogofaster@latest
go install github.com/favadi/protoc-go-inject-tag@latest
go install github.com/gogo/protobuf/gogoproto@latest