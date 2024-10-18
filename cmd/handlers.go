package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dmikhr/chat-server/internal/data"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/dmikhr/chat-server/pkg/chat_v1"
)

// CreateChat - создание нового чата для пользователей
// возвращает id чата
func (s *server) CreateChat(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("UserId: %v | Chat name: %v", req.GetUserid(), req.GetName())
	chatID, err := s.models.Chat.Create(&data.Chat{Name: req.GetName(),
		Users:     req.GetUserid(),
		CreatedAt: time.Now()})
	if err != nil {
		fmt.Println("Error while creating chat:", err)
		return nil, err
	}
	return &desc.CreateResponse{Id: chatID}, nil
}

// DeleteChat - удаление чата по id
func (s *server) DeleteChat(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Chat to delete id: %d", req.GetId())
	err := s.models.Chat.Delete(req.GetId())
	if err != nil {
		fmt.Println("Error deleting chat:", err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// SendMessageToChat - отправка сообщения
// From - автор сообщения, Text - текст сообщения, Timestamp - время отправки
func (s *server) SendMessageToChat(_ context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Message from: %s | text: %s | time: %v", req.GetFrom(), req.GetText(), req.GetTimestamp())
	err := s.models.Message.Send(&data.Message{From: req.GetFrom(),
		Msg:       req.GetText(),
		CreatedAt: data.TimestamppbToTime(req.GetTimestamp())})
	if err != nil {
		fmt.Println("Error while sending message:", err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
