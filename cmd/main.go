package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/dmikhr/chat-server/pkg/chat_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatV1Server
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Create - создание нового чата для пользователей req.Usernames
// возвращает id чата
func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("UserId: %v | Chat name: %v", req.GetUserid(), req.GetName())
	return &desc.CreateResponse{}, nil
}

// Delete - удаление чата по id
func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Chat to delete id: %d", req.GetId())
	return &emptypb.Empty{}, nil
}

// SendMessage - отправка сообщения
// From - автор сообщения, Text - текст сообщения, Timestamp - время отправки
func (s *server) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Message from: %s | text: %s | time: %v", req.GetFrom(), req.GetText(), req.GetTimestamp())
	return &emptypb.Empty{}, nil
}
