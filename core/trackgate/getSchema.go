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

func ListSchema(client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, stationId string) {
	queryClient := types.NewQueryClient(client.Context())

	params := &types.QueryListSchemasRequest{
		ExtTrackStationId: stationId,
	}

	schemas, err := queryClient.ListSchemas(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.MarshalIndent(schemas, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	fmt.Println(string(jsonData))
}

func ListSchemaWithPagination(client cosmosclient.Client, ctx context.Context, stationId string, offset uint64, limit uint64, reverse bool) {
	queryClient := types.NewQueryClient(client.Context())

	params := &types.QueryListSchemasRequest{
		ExtTrackStationId: stationId,
		Pagination: &query.PageRequest{
			Offset:     offset,
			Limit:      limit,
			CountTotal: true,
			Reverse:    reverse,
		},
	}

	schemas, err := queryClient.ListSchemas(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.MarshalIndent(schemas, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	fmt.Println(string(jsonData))
}

func RetriveSchemaKey(client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, stationId string, version string) string {
	queryClient := types.NewQueryClient(client.Context())

	params := &types.QueryRetrieveSchemaKeyRequest{
		ExtTrackStationId: stationId,
		SchemaVersion:     version,
	}

	schema, err := queryClient.RetrieveSchemaKey(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.MarshalIndent(schema, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	fmt.Println(string(jsonData))
	 
	return schema.SchemaKey
}

