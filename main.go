package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/mkolchurin/bad_msg_handler/proto/req"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type MessageHandler struct {
	pb.UnimplementedMessageHandlerServer
}

func (s *MessageHandler) reqHandler(ctx context.Context, in *pb.PhraseRequest) (*pb.PhraseReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.PhraseReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMessageHandlerServer(s, &MessageHandler{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
