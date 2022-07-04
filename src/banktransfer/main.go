package main

import (
	"os"
	"net"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/mgr1054/myaktion-go/src/banktransfer/grpc/banktransfer"
	"github.com/mgr1054/myaktion-go/src/banktransfer/service"

)

func init() {
	// init logger
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
	level, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Info("Log level not specified, set default to: INFO")
		log.SetLevel(log.InfoLevel)
	}
	log.SetLevel(level)
}

var grpcPort = 9111
func main() {
	log.Info("Starting Banktransfer server")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen on grpc port %d: %v", grpcPort, err)
	}
	grpcServer := grpc.NewServer()
	banktransfer.RegisterBankTransferServer(grpcServer, service.NewBankTransferService())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}