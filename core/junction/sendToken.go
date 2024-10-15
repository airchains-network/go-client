package junction

import (
	// "context"
	// "fmt"
	// "log"

	// cosmosTypes "github.com/cosmos/cosmos-sdk/types"
	// cosmosBankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	// "github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	// "github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
	// "github.com/airchains-network/junction-go-client/components"
)

// func SendToken(reqAmount int64, toAddress string, ctx context.Context, account cosmosaccount.Account, adminAddress string) (success bool, message string, txhash string, err error) {

// 	accountPath := "./accounts"
// 	addressPrefix := "air"

// 	// Create a Cosmos client instance
// 	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix), cosmosclient.WithNodeAddress(components.JunctionRPC), cosmosclient.WithHome(accountPath))
// 	if err != nil {
// 		log.Fatalf("Failed to create Cosmos client: %v", err)
// 	}

// 	// Check if the admin has enough balance to send tokens
// 	hasBalance, balance, err := CheckBalance(ctx, adminAddress, client)
// 	if err != nil {
// 		return false, "Failed to check admin balance", "", err
// 	}
// 	fmt.Printf("Admin balance: %v\n", balance)

// 	// No balance in admin wallet
// 	if !hasBalance {
// 		return false, "No balance in admin wallet", "", nil
// 	}

// 	// Admin has less than 10 tokens (not enough balance)
// 	if balance < 10 {
// 		return false, "Admin does not have enough tokens", "", nil
// 	}

// 	msg := &cosmosBankTypes.MsgSend{
// 		FromAddress: adminAddress,
// 		ToAddress:   toAddress,
// 		Amount:      cosmosTypes.NewCoins(cosmosTypes.NewInt64Coin("amf", reqAmount)),
// 	}

// 	txResp, err := client.BroadcastTx(ctx, account, msg)
// 	if err != nil {
// 		return false, "Error in transaction", "", err
// 	}

// 	return true, "Transaction successful", txResp.TxHash, nil
// }
