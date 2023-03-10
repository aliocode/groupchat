package api

import (
	"context"
	"log"

	desc "groupchat/generated/groupchatpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) JoinGroupChat(ctx context.Context, req *desc.JoinGroupChatRequest) (*desc.JoinGroupChatResponse, error) {
	if len(req.ChatName) == 0 {
		log.Printf("failed to register group: received empty chat name")
		return nil, status.Error(codes.InvalidArgument, "chat name is empty")
	}

	if len(req.Username) == 0 {
		log.Printf("failed to register group: received empty client name")
		return nil, status.Error(codes.InvalidArgument, "client id is empty")
	}

	err := i.groupService.RegisterListener(ctx, req.ChatName, req.Username)
	if err != nil {
		log.Printf("failed to register listener: %v", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &desc.JoinGroupChatResponse{}, nil
}
