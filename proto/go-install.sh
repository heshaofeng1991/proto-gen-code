#!/bin/sh

go install github.com/gogo/protobuf@v1.3.2
go install github.com/mwitkow/go-proto-validators@v0.3.2
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators@latest
go install github.com/heshaofeng1991/protoc-gen-go-http@latest
go install github.com/gogo/protobuf/protoc-gen-gogofaster@latest
go install github.com/favadi/protoc-go-inject-tag@v1.3.0

#for p in admin api job; do
#  if [ "$(ls ./*/$p/*.proto 2>/dev/null)" != "" ]; then
#    protoc -I=. --go_out=plugins=grpc:. --ginhttp_out=:. ./*/$p/*.proto
#    sed -i "" -E 's/(@gotags:.*)json:"[^"]*"(.*)/\1\2/' ./*/$p/proto/*.pb.go
#    sed -i "" -E 's/(@gotags:.*)form:"[^"]*"(.*)/\1\2/' ./*/$p/proto/*.pb.go
#    sed -i "" -E '/^[[:space:]]*\/\/[[:space:]]*@gotags:[[:space:]]*$/d' ./*/$p/proto/*.pb.go
#    sed -i "" -E 's/(protobuf:.*) json:"([^"]*)"/\1 json:"\2" form:"\2"/' ./*/$p/proto/*.pb.go
#    sed -i "" -E 's/,omitempty"/"/g' ./*/$p/proto/*.pb.go
#    protoc-go-inject-tag -input "./*/$p/proto/*.pb.go"
#  fi
#done