package protobufapi

import (
	"context"
	"log"
)

//防止被前置卡死
type Server struct {
	UnimplementedChatServiceServer
}

// 实现 SayHello
func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Received SayHello: %v", in.Body)
	return &Message{Body: "Hello " + in.Body}, nil
}

// 实现 Test
func (s *Server) Test(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Received Test: %v", in.Body)
	return &Message{Body: "Test OK: " + in.Body}, nil
}
