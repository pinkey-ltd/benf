package handler

import (
	"context"
	"gofr.dev/pkg/gofr/container"
	"google.golang.org/protobuf/types/known/emptypb"
	"user-service/pb"
)

type Handler struct {
	// container can be embedded into the server struct
	// to access the datasource and logger functionalities
	*container.Container

	pb.UnimplementedUserServiceServer
}

func (h *Handler) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// get the user data and handler error
	return nil, nil
}

func (h *Handler) GetUserList(ctx context.Context, in *emptypb.Empty) (*pb.GetUserListResponse, error) {

	return nil, nil
}
