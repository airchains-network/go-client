package trackgate

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/airchains-network/junction-go-client/components"
	"github.com/airchains-network/junction/x/trackgate/types"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
)

func ReadSchemaVersionFromFile() (string, bool) {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filePath := cwd + "/data/schemaVersion.txt"
	dat, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(filePath)
			if err != nil {
				return "", false
			}
			defer file.Close()

			_, err = file.WriteString("1.0.0")
			if err != nil {
				return "", false
			}
			return "1.0.0", true
		}
		return "", false
	}
	return string(dat), true
}

func InitSchema(client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, stationId string, address string) {
	// Marshal the schema to JSON
	var SchemaV1 = types.SchemaDef{
		Fields: map[string]interface{}{
			"espresso_tx_response_v_1": map[string]interface{}{
				"transaction": map[string]interface{}{
					"namespace": "int",
					"payload":   "string",
				},
				"hash":  "string",
				"index": "int",
				"proof": map[string]interface{}{
					"tx_index":        "string",
					"payload_num_txs": "string",
					"payload_proof_num_txs": map[string]interface{}{
						"proofs":       "string",
						"prefix_bytes": "string",
						"suffix_bytes": "string",
					},
					"payload_tx_table_entries": "string",
					"payload_proof_tx_table_entries": map[string]interface{}{
						"proofs":       "string",
						"prefix_bytes": "string",
						"suffix_bytes": "string",
					},
					"payload_proof_tx": map[string]interface{}{
						"proofs":       "string",
						"prefix_bytes": "string",
						"suffix_bytes": "string",
					},
				},
				"block_hash":   "string",
				"block_height": "int",
			},
			"station_id": "string",
		},
	}
	
	schemaBytes, err := json.Marshal(SchemaV1)
	if err != nil {
		log.Fatalf("failed to marshal schema: %v", err)
	}

	schemaVersion, ok := ReadSchemaVersionFromFile()
	if !ok {
		log.Fatalf("failed to read schema version from file")
	}


	// Increment the schema version for the next initialization
	schemaVersion = components.IncrementVersion(schemaVersion)

	// Create a new message for schema creation
	msgSchema := &types.MsgSchemaCreation{
		Creator:           address,
		ExtTrackStationId: stationId,
		Version:           schemaVersion,
		Schema:            schemaBytes,
	}

	fmt.Println("msg: ", msgSchema)

	// Broadcast the schema creation transaction
	txResp, err := client.BroadcastTx(ctx, account, msgSchema)
	if err != nil {
		log.Fatalf("failed to broadcast transaction: %v", err)
	}

	// Print response from broadcasting a transaction
	fmt.Print("MsgSchemaCreation:\n\n")
	fmt.Println(txResp)

	// Uncomment if you want to print the message in JSON format
	msgJSON, err := json.MarshalIndent(msgSchema, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting JSON: %v", err)
	}
	fmt.Println(msgJSON)
}
