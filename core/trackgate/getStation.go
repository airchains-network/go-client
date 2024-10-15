package trackgate

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/airchains-network/junction/x/trackgate/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
)

func ListStation(client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account) {
	queryClient := types.NewQueryClient(client.Context())

	params := &types.QueryListExtTrackStationsRequest{}

	stations, err := queryClient.ListExtTrackStations(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.MarshalIndent(stations, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	fmt.Println(string(jsonData))

}

func ListStationWithPagination(client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, offset uint64, reverse bool) {
	queryClient := types.NewQueryClient(client.Context())

	params := &types.QueryListExtTrackStationsRequest{
		Pagination: &query.PageRequest{
			Limit:   1,      // Set the limit for the number of results
			Offset:  offset,  // Use the provided offset for pagination
			Reverse: reverse, // Set the reverse flag based on the function parameter
			CountTotal: true,
		},
	}

	stations, err := queryClient.ListExtTrackStations(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.MarshalIndent(stations, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	fmt.Println(string(jsonData))
}
