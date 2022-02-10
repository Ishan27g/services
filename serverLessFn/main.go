package main

import (
	"fmt"

	"serverLessFn/handler"
	pb "serverLessFn/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("serverlessfn"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterServerLessFnHandler(srv.Server(), new(handler.ServerLessFn))
	fmt.Println("ok")
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
