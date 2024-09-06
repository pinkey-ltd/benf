package main

import (
	"gofr.dev/pkg/gofr"
	"user-service/internal/handler"
	"user-service/pb"
)

func main() {
	service := gofr.New()
	pb.RegisterUserServiceServer(service, &handler.Handler{})

	service.Run()
}
