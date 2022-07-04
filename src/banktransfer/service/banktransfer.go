package service

import (
	"context"
	"time"


	log "github.com/sirupsen/logrus"
	"github.com/mgr1054/myaktion-go/src/banktransfer/grpc/banktransfer"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BankTransferService struct {
	banktransfer.BankTransferServer
	counter int32
}

func NewBankTransferService() *BankTransferService {
	return &BankTransferService{counter: 1}
}

func (s *BankTransferService) TransferMoney(_ context.Context, transaction *banktransfer.Transaction) (*emptypb.Empty, error) {
	log.Infof("Received transaction: %v", transaction)
	return &emptypb.Empty{}, nil
}

func (s *BankTransferService) ProcessTransactions(stream banktransfer.BankTransfer_ProcessTransactionsServer) error {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	return func() error {
		for {
				select {
				case <-stream.Context().Done():
					log.Info("Watching transactions cancelled from the client side")
					return nil

				case _ = <-ticker.C: transaction := &banktransfer.Transaction{Id: s.counter, Amount: 20}
					entry := log.WithField("transaction", transaction)
					entry.Info("Sending transaction")
					if err := stream.Send(transaction); err != nil {
							entry.WithError(err).Error("Error sending transaction")
							return err
					}
					entry.Info("Transaction sent. Waiting for processing response")
					response, err := stream.Recv()
					if err != nil {
							entry.WithError(err).Error("Error receiving processing response")
							return err
					}
					if response.Id != s.counter {
							// NOTE: this is just a guard and not happening as transaction is local per connection
							entry.Error("Received processing response of a different transaction")
					} else {
							entry.Info("Processing response received")
							s.counter++
					}	
				}
		}
	}()
}

