# hello-grpc

```
# protocol buffers
$ brew update
$ brew install protobuf
$ protoc --version
libprotoc 3.5.1

# grpc 
$ go get git@github.com:grpc/grpc-go.git
$ go get -u github.com/golang/protobuf/protoc-gen-go

# compile
$ protoc --go_out=plugins=grpc:. cat/cat.proto
```