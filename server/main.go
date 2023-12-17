package main

import (
	//"net"
	//"os"
	//"os/signal"
	//"syscall"
	//
	//pb "github.com/jankremlacek/go-bazel/proto/servicea"
	//"github.com/jankremlacek/go-bazel/shared"
	"go.uber.org/zap"
	//"google.golang.org/grpc"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
}

//
//import (
//	"context"
//	"log"
//	"net"
//
//	pb "github.com/abitofhelp/bazel7_go_grpc/proto/hello_world"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/reflection"
//)
//
//const (
//	port = ":50051"
//)
//
//// server is used to implement helloworld.GreeterServer.
//type server struct{}
//
//// SayHello implements helloworld.GreeterServer
//
//func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
//	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
//}
//func main() {
//	lis, err := net.Listen("tcp", port)
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//	s := grpc.NewServer()
//	pb.RegisterGreeterServer(s, &server{})
//	// Register reflection service on gRPC server.
//	reflection.Register(s)
//	if err := s.Serve(lis); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//}
