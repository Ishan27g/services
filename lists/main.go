package main

import (
	"github.com/micro/services/pkg/tracing"
	"lists/handler"
	pb "lists/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("lists"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterListsHandler(srv.Server(), new(handler.Lists))

	traceCloser := tracing.SetupOpentracing("db")
	defer traceCloser.Close()

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
