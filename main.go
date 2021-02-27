package main

import (
	"go-user/handler"
	pb "go-user/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("go-user"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterGoUserHandler(srv.Server(), new(handler.GoUser))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
