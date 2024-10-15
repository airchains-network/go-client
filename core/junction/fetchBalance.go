package junction

import (
	"context"
	"log"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
)

// CheckBalance retrieves the balance of a specified account and checks if it holds any tokens of the specified denomination.
func CheckBalance(ctx context.Context, accountAddress string, client cosmosclient.Client) (bool, int64, error) {
	pageRequest := &query.PageRequest{}

	balances, err := client.BankBalances(ctx, accountAddress, pageRequest)
	if err != nil {
		log.Printf("Error querying bank balances: %v", err)
		return false, 0, err
	}

	for _, balance := range balances {
		if balance.Denom == "amf" {
			return true, balance.Amount.Int64(), nil
		}
	}

	return false, 0, nil
}
