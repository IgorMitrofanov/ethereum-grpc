package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"google.golang.org/grpc"

	"ethereum-grpc/internal/proto"
)

type server struct {
	proto.UnimplementedEthereumServiceServer
	client *ethclient.Client
}

func (s *server) GetAccount(ctx context.Context, req *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	accountAddress := common.HexToAddress(req.EthereumAddress)

	balance, err := s.client.BalanceAt(ctx, accountAddress, nil)
	if err != nil {
		return nil, err
	}

	nonce, err := s.client.NonceAt(ctx, accountAddress, nil)
	if err != nil {
		return nil, err
	}

	gasTokenBalance := balance.String()

	response := &proto.GetAccountResponse{
		GastokenBalance: gasTokenBalance,
		WalletNonce:     nonce,
	}

	return response, nil
}

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()

	proto.RegisterEthereumServiceServer(grpcServer, &server{client: client})

	fmt.Println("gRPC server listening on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
