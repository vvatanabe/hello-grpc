package main

import (
	"errors"

	"net"

	pb "github.com/vvatanabe/hello-grpc/cat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type CatServer struct {
}

func (s *CatServer) GetMyCat(ctx context.Context, message *pb.GetMyCatMessage) (*pb.MyCatResponse, error) {
	switch message.TargetCat {
	case "tama":
		return &pb.MyCatResponse{
			Name: "tama",
			Kind: "mainecoon",
		}, nil
	case "mike":
		return &pb.MyCatResponse{
			Name: "mike",
			Kind: "Norwegian Forest Cat",
		}, nil
	}
	return nil, errors.New("not found your cat")
}

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	catServer := &CatServer{}
	pb.RegisterCatServer(server, catServer)
	server.Serve(listenPort)
}
