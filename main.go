package main

import (
	"context"
	"fmt"
	"log"

	"github.com/airchains-network/junction-go-client/cli"
	"github.com/airchains-network/junction-go-client/components"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
	junction "github.com/airchains-network/junction-go-client/core/junction"
)

func main() {
	addressPrefix := "air"
	accountPath := "./accounts"
	accountName := "charlie"
	ctx := context.Background()

	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix), cosmosclient.WithNodeAddress(components.JunctionRPC), cosmosclient.WithHome(accountPath), cosmosclient.WithGas("auto"), cosmosclient.WithFees("1000amf"))
	if err != nil {
		log.Fatal(err)
	}
	exists, _ := components.CheckIfAccountExists(accountName, client, addressPrefix, accountPath)
	if !exists {
		components.CreateAccount(accountName, accountPath)
	}

	
	account, err := client.Account(accountName)
	if err != nil {
		fmt.Println("Error getting account")
	}
	addr, err := account.Address(addressPrefix)
	if err != nil {
		fmt.Println("Error getting account address")
	}
	
	hasBalance, balance, err := junction.CheckBalance(ctx, addr, client)
	if err != nil {
		log.Fatalf("Error checking balance: %v", err)
	}
	if !hasBalance || balance == 0 {
		log.Fatal("Account balance is zero. Please add balance to the account.")
	}
	
	cli.Trackgate(ctx, client, account, addr)
}
