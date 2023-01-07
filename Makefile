$(info $(SHELL))

GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@echo "start init command..."
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators@latest
	@go install github.com/heshaofeng1991/protoc-gen-go-http@latest
	@go install github.com/gogo/protobuf/protoc-gen-gogofaster@latest
	@go install github.com/favadi/protoc-go-inject-tag@latest
	@go install github.com/gogo/protobuf/gogoproto@latest

.PHONY: proto
proto:
	@echo "start proto command..."
	@protoc -I=. -I=${GOPATH}/pkg/mod \
 	--gogofaster_out=. --gogofaster_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --govalidators_out=. --govalidators_opt=paths=source_relative \
    --go-http_out=. --go-http_opt=paths=source_relative \
    proto/admin/admin.proto

.PHONY: inject-tag
inject-tag:
	@echo "start inject-tag command..."
	@protoc-go-inject-tag -input=${GOPATH}/src/tools/proto-gen-code/proto/admin/admin.pb.go
	@protoc-go-inject-tag -input=${GOPATH}/src/tools/proto-gen-code/proto/web/web.pb.go

.PHONY: gin
gin:
	@echo "start gin command..."
	@protoc -I=. -I=${GOPATH}/pkg/mod \
    --go_out=plugins=grpc:. \
    --ginhttp_out=:. \
    proto/web/web.proto

.PHONY: mv
mv:
	@echo "start make mv command..."
	@cd ${GOPATH}/src/tools/proto-gen-code/
	@mv ${GOPATH}/src/tools/proto-gen-code/*.pb.go ${GOPATH}/src/tools/proto-gen-code/proto/web/

.PHONY: all
all: proto gin inject-tag mv

