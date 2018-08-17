# hello-grpc

## Prepare

### protocol buffers

for Mac:
```
$ brew update
$ brew install protobuf
$ protoc --version
```

for linux:
```
$ curl -OL https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip
$ unzip protoc-3.6.1-linux-x86_64.zip -d protoc3
$ sudo mv protoc3/bin/* /usr/local/bin/
$ sudo mv protoc3/include/* /usr/local/include/
$ sudo chown $USER /usr/local/bin/protoc
$ sudo chown -R $USER /usr/local/include/google
$ protoc --version
```

### gRPC

```
$ go get git@github.com:grpc/grpc-go.git
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

### compile
```
$ protoc --go_out=plugins=grpc:. cat/cat.proto
```