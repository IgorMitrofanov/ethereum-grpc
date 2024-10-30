package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"ethereum-grpc/internal/proto"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewEthereumServiceClient(conn)

	req := &proto.GetAccountRequest{
		EthereumAddress: "0x71c7656ec7ab88b098defb751b7401b5f6d8976f",
		CryptoSignature: "HZ",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := client.GetAccount(ctx, req)
	if err != nil {
		log.Fatalf("could not get account: %v", err)
	}

	fmt.Printf("Gas Token Balance: %s\n", response.GastokenBalance)
	fmt.Printf("Wallet Nonce: %d\n", response.WalletNonce)
}
