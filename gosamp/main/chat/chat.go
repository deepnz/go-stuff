package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s * Server) Hello(cntx context.Context, message *Message) (
	*Message, error) {
		log.Printf("Received message successfully from Client: %s", message.Body)
		return &Message{Body: "Hello From the other Server!"}, nil
}
