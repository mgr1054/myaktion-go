package main

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/mgr1054/myaktion-go/myaktion-go/src/myaktion/client"
	"github.com/mgr1054/myaktion-go/myaktion-go/src/myaktion/client/banktransfer"
	"github.com/mgr1054/myaktion-go/myaktion-go/src/myaktion/service"
)

func monitortransactions() {
	for {
		connectandmonitor()
		time.Sleep(time.Second)
	}
}
func connectandmonitor() {
	conn, err := client.GetBankTransferConnection()
	if err != nil {
		log.WithError(err).Fatal("error connecting to the banktransfer service")
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	banktransferClient := banktransfer.NewBankTransferClient(conn)
	watcher, err := banktransferClient.ProcessTransactions(ctx)
	if err != nil {
		log.WithError(err).Fatal("error watching transactions")
	}
	log.Info("Successfully connected to banktransfer service for processing transactions")
	for {
		transaction, err := watcher.Recv()
		if err != nil {
			if _, deadline := ctx.Deadline(); deadline {
				log.Info("deadline reached. reconnect client")
				break
			}
			log.WithError(err).Error("error receiving transaction")
			continue
		}
		entry := log.WithField("transaction", transaction)
		entry.Info("Received transaction")
		err = service.MarkDonation(uint(transaction.DonationId))
		if err != nil {
			entry.WithError(err).Error("error changing donation status")
			continue
		}
		entry.Info("Sending processing response")
		err = watcher.Send(&banktransfer.ProcessingResponse{Id: transaction.Id})
		if err != nil {
			entry.WithError(err).Error("error sending processing response")
			continue
		}
		entry.Info("Processing response sent")
	}
}
