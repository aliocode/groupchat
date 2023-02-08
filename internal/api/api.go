package api

import (
	desc "groupchat/generated/groupchatpb"
)

type Implementation struct {
	desc.UnimplementedMessageServiceServer
	clientService clientService
	groupService  groupService
}

// NewAPI return new instance of Implementation.
func NewAPI(
	clientService clientService,
	groupService groupService,
) *Implementation {
	return &Implementation{
		clientService: clientService,
		groupService:  groupService,
	}
}
