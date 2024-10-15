package trackgate

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/airchains-network/junction/x/trackgate/types"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
)

func GetTrackEngagement(client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, stationId string, podNumber uint64) {
	queryClient := types.NewQueryClient(client.Context())

	params := &types.QueryGetTrackEngagementRequest{
		ExtTrackStationId: stationId,
		PodNumber:         podNumber,
	}

	schemas, err := queryClient.GetTrackEngagement(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.MarshalIndent(schemas, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	fmt.Println(string(jsonData))
}

func ListTrackEngagementwithPagination(client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, stationId string, offset uint64, limit uint64, order string) {
	queryClient := types.NewQueryClient(client.Context())

	params := &types.QueryListTrackEngagementsRequest{
		ExtTrackStationId: stationId,
		Pagination: &types.TrackgatePaginationRequest{
			Offset: offset,
			Limit:  limit,
			Order:  order,
		},
	}

	schemas, err := queryClient.ListTrackEngagements(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.MarshalIndent(schemas, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	fmt.Println(string(jsonData))
}

