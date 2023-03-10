package client

import (
	"context"
	"fmt"
	"sync"

	"groupchat/internal/model"
)

type Service struct {
	memMap *sync.Map
}

func New() *Service {
	return &Service{
		memMap: &sync.Map{},
	}
}

func (s *Service) RegisterClient(_ context.Context, clientID string, ch chan model.Message) error {
	_, loaded := s.memMap.LoadOrStore(clientID, ch)
	if loaded {
		return fmt.Errorf("RegisterClient: client already registered")
	}

	return nil
}

func (s *Service) UnregisterClient(_ context.Context, clientID string) error {
	s.memMap.Delete(clientID)

	return nil
}

func (s *Service) SendMessage(_ context.Context, msg model.Message) error {
	val, ok := s.memMap.Load(msg.ReceiverID)
	if !ok {
		return fmt.Errorf("SendMessage: client's data was not found")
	}

	ch, ok := val.(chan model.Message)
	if !ok {
		return fmt.Errorf("SendMessage: client's data corrupted")
	}

	ch <- msg

	return nil
}
