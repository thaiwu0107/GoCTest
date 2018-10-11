package main

import (
	Fmt "fmt"
	"log"
	"net"
	"runtime"

	"git.com.ggttoo44/src/grpcinit"
	Init "git.com.ggttoo44/src/init"
	"google.golang.org/grpc"
	Reflection "google.golang.org/grpc/reflection"
)

const (
	port = "localhost:50051"
)

// 這是呼叫GRPC
func main() {
	runtime.GOMAXPROCS(8)
	Init.Start()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// 註冊現有的GRPC Server
	grpcinit.RegisterAllServer(s)
	// Register reflection service on gRPC server.
	Reflection.Register(s)
	Fmt.Println("RegisterAllServer Done Server On Port: ", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
