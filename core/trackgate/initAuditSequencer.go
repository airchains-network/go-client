package trackgate

import (
	"context"
	"fmt"

	"github.com/airchains-network/junction/x/trackgate/types"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
)

func InitAuditSequencer(client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, stationId string, operator string) {
	msgAuditSequencer := &types.MsgAuditSequencer{
		Verifier: operator,
		SequencerChecks: []types.ExtSequencerCheck{
			{
				PodNumber:          "1",
				ExtTrackStationId:  stationId,
				Namespace:          "sequencer-namespace-1",
				VerificationStatus: true,
			},
		},
	}
	
	txResp, err := client.BroadcastTx(ctx, account, msgAuditSequencer)
	if err != nil {
		fmt.Println("Error broadcasting transaction:", err)
		return
	}

	fmt.Println("Transaction response:", txResp)
}
