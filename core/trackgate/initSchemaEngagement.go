package trackgate

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/airchains-network/junction-go-client/components"
	internal "github.com/airchains-network/junction-go-client/types"
	"github.com/airchains-network/junction/x/trackgate/types"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
)

const (
	dataDir = "data"
	podFile = "pod_number.txt"
)

func InitSchemaEngagement(client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, stationId string, operator string) {
	podNumber, err := readPodNumberFromFile()
	if err != nil {
		log.Fatalf("failed to read pod number from file: %v", err)
	}

	namespace, err := strconv.Atoi(components.GenerateRandomNumber())
	if err != nil {
		log.Fatalf("failed to convert namespace to int: %v", err)
	}

	acknowledgementHash, err := components.GenerateTransactionHash()
	if err != nil {
		log.Fatalf("failed to generate transaction hash: %v", err)
	}

	index, err := strconv.Atoi(components.GenerateRandomNumber())
	if err != nil {
		log.Fatalf("failed to convert index to int: %v", err)
	}

	blockHeight, err := strconv.Atoi(components.GenerateRandomNumber())
	if err != nil {
		log.Fatalf("failed to convert block height to int: %v", err)
	}

	espressoStruct := createEspressoStruct(namespace, acknowledgementHash, index, blockHeight, stationId)

	espressoJsonBytes, err := json.Marshal(espressoStruct)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	address, err := components.GenerateEthereumAddress()
	if err != nil {
		log.Fatalf("failed to generate Ethereum address: %v", err)
	}

	schemaVersion, ok := ReadSchemaVersionFromFile()
	if !ok {
		log.Fatalf("failed to read schema version from file")
	}

	sequencerDetails := createSequencerDetails(schemaVersion, namespace, address)

	sequencerDetailsBytes, err := json.Marshal(sequencerDetails)
	if err != nil {
		log.Fatalf("failed to marshal sequencer details: %v", err)
	}

	msgSchemaEngage := &types.MsgSchemaEngage{
		Operator:            operator,
		ExtTrackStationId:   stationId,
		SchemaObject:        espressoJsonBytes,
		AcknowledgementHash: acknowledgementHash,
		PodNumber:           podNumber,
		SequencerDetails:    sequencerDetailsBytes,
	}

	txResp, err := client.BroadcastTx(ctx, account, msgSchemaEngage)
	if err != nil {
		log.Fatalf("failed to broadcast transaction: %v", err)
	}
	fmt.Println("Transaction response:", txResp)

	podNumber++
	err = writePodNumberToFile(podNumber)
	if err != nil {
		log.Fatalf("failed to write pod number to file: %v", err)
	}
}

func createEspressoStruct(namespace int, acknowledgementHash string, index int, blockHeight int, stationId string) internal.EspressoSchemaV1 {
	return internal.EspressoSchemaV1{
		EspressoTxResponseV1: internal.EspressoTxResponseV1{
			Transaction: internal.Transaction{
				Namespace: namespace,
				Payload:   components.GenerateRandomString(10),
			},
			Hash:  acknowledgementHash,
			Index: index,
			Proof: internal.Proof{
				TxIndex:       components.GenerateRandomString(10),
				PayloadNumTxs: components.GenerateRandomString(10),
				PayloadProofNumTxs: internal.ProofDetails{
					Proofs:      components.GenerateRandomString(10),
					PrefixBytes: components.GenerateRandomString(5),
					SuffixBytes: components.GenerateRandomString(5),
				},
				PayloadTxTableEntries: components.GenerateRandomString(10),
				PayloadProofTxTableEntries: internal.ProofDetails{
					Proofs:      components.GenerateRandomString(10),
					PrefixBytes: components.GenerateRandomString(5),
					SuffixBytes: components.GenerateRandomString(5),
				},
				PayloadProofTx: internal.ProofDetails{
					Proofs:      components.GenerateRandomString(10),
					PrefixBytes: components.GenerateRandomString(5),
					SuffixBytes: components.GenerateRandomString(5),
				},
			},
			BlockHash:   components.GenerateRandomString(10),
			BlockHeight: blockHeight,
		},
		StationId: stationId,
	}
}

func createSequencerDetails(schemaVersion string, namespace int, address string) types.SequencerDetails {
	return types.SequencerDetails{
		Name:      "espresso",
		Version:   schemaVersion,
		NameSpace: fmt.Sprintf("%d", namespace),
		Address:   address,
	}
}

func readPodNumberFromFile() (uint64, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return 0, err
	}
	filePath := filepath.Join(cwd, dataDir, podFile)

	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.WriteFile(filePath, []byte("1"), 0644); err != nil {
				return 0, err
			}
			return 1, nil
		}
		return 0, err
	}

	podNumber, err := strconv.ParseUint(string(data), 10, 64)
	if err != nil {
		return 0, err
	}
	return podNumber, nil
}

func writePodNumberToFile(podNumber uint64) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	filePath := filepath.Join(cwd, dataDir, podFile)
	return os.WriteFile(filePath, []byte(strconv.FormatUint(podNumber, 10)), 0644)
}

func ResetPodNumberToFile() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	filePath := filepath.Join(cwd, dataDir, podFile)
	return os.WriteFile(filePath, []byte("1"), 0644)
}
