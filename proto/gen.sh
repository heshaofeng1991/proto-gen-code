#!/bin/sh

vergte() {
    [ "$2" = "`printf "$1\n$2\n" | sort -V | head -n1`" ]
}

for p in admin web; do
  echo $p
  echo ls ./*/$p/*.proto 2>/dev/null
  if [ "$(ls ./*/$p/*.proto 2>/dev/null)" != "" ]; then
    make all
    sed -i "" -E 's/(@gotags:.*)json:"[^"]*"(.*)/\1\2/' ./*/$p/*.pb.go
    sed -i "" -E 's/(@gotags:.*)form:"[^"]*"(.*)/\1\2/' ./*/$p/*.pb.go
    sed -i "" -E '/^[[:space:]]*\/\/[[:space:]]*@gotags:[[:space:]]*$/d' ./*/$p/*.pb.go
    sed -i "" -E 's/(protobuf:.*) json:"([^"]*)"/\1 json:"\2" form:"\2"/' ./*/$p/*.pb.go
    sed -i "" -E 's/,omitempty"/"/g' ./*/$p/*.pb.go
    protoc-go-inject-tag -input "./*/$p/*.pb.go"
  fi
done