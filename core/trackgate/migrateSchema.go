package trackgate

import (
	"context"
	"fmt"
	"log"

	"github.com/airchains-network/junction/x/trackgate/types"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
)

func MigrateSchema(client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, stationId string, operator string, schemaKey string) {

	msgSchemaMigrate := &types.MsgMigrateSchema{
		Operator:          operator,
		ExtTrackStationId: stationId,
		NewSchemaKey:      schemaKey,
	}

	txResp, err := client.BroadcastTx(ctx, account, msgSchemaMigrate)
	if err != nil {
		fmt.Println("txResp above")
		fmt.Println(txResp)
		log.Fatal(err)
	}
	fmt.Println("txResp below")
	fmt.Println(txResp)
}
