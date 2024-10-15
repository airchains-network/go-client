package trackgate

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/rand"

	"github.com/airchains-network/junction-go-client/components"
	"github.com/airchains-network/junction/x/trackgate/types"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
)

func InitStation(addr string, client cosmosclient.Client, ctx context.Context, account cosmosaccount.Account, stationName string) string {
	stationId := uuid.New().String()

	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	// Check if the data directory exists, if not create it
	dataDir := filepath.Join(cwd, "data")
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
			log.Fatalf("Failed to create data directory: %v", err)
		}
	}

	// Write the station information details to a file
	filePath := fmt.Sprintf("%s/stationInfoDetails_%s.json", dataDir, stationId)
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	daTypes := []string{"celestia", "eigen", "avail"}
	stationTypes := []string{"evm", "cosmos", "svm"}
	rand.Seed(uint64(time.Now().UnixNano()))
	selectedDAType := daTypes[rand.Intn(len(daTypes))]
	selectedStationType := stationTypes[rand.Intn(len(stationTypes))]

	stationInfo := types.StationInfoDetails{
		StationName: stationName,
		Type:        selectedStationType,
		FheEnabled:  false,
		Operators:   []string{addr, components.GenerateRandomBech32Address(), components.GenerateRandomBech32Address(), components.GenerateRandomBech32Address()},
		SequencerDetails: types.SequencerDetails{
			Name:    components.GenerateRandomString(6),
			Version: fmt.Sprintf("%d.%d.%d", rand.Intn(9), rand.Intn(9), rand.Intn(9)),
		},
		DADetails: types.DADetails{
			Name:    components.GenerateRandomString(6),
			Type:    selectedDAType,
			Version: fmt.Sprintf("%d.%d.%d", rand.Intn(9), rand.Intn(9), rand.Intn(9)),
		},
		ProverDetails: types.ProverDetails{
			Name:    components.GenerateRandomString(6),
			Version: fmt.Sprintf("%d.%d.%d", rand.Intn(9), rand.Intn(9), rand.Intn(9)),
		},
	}

	if err := encoder.Encode(stationInfo); err != nil {
		log.Fatalf("Failed to encode station information: %v", err)
	}

	// Read the station information details from the file
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	msg := &types.MsgInitStation{
		Submitter:   addr,
		StationId:   stationId,
		StationInfo: fileContent,
		Operators:   stationInfo.Operators,
	}

	// Broadcast a transaction from account with the message
	txResp, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		log.Fatalf("Failed to broadcast transaction: %v", err)
	}

	// Print response from broadcasting a transaction
	fmt.Println("MsgCreatePost:\n\n", txResp)

	// Instantiate a query client for the `trackgate` blockchain
	queryClient := types.NewQueryClient(client.Context())

	queryResp, err := queryClient.GetExtTrackStation(ctx, &types.QueryGetExtTrackStationRequest{Id: stationId})
	if err != nil {
		log.Fatalf("Failed to query station: %v", err)
	}

	data := queryResp.Station
	fmt.Printf("Operator: %v\n", data.Operators)
	fmt.Printf("LatestPod: %v\n", data.LatestPod)
	fmt.Printf("LatestAcknowledgementHash: %v\n", data.LatestAcknowledgementHash)
	fmt.Printf("Name: %v\n", data.Name)
	fmt.Printf("Id: %v\n", data.Id)
	fmt.Printf("StationType: %v\n", data.StationType)
	fmt.Printf("FheEnabled: %v\n", data.FheEnabled)

	var seqData types.SequencerDetails
	if err := json.Unmarshal(data.SequencerDetails, &seqData); err != nil {
		log.Fatalf("Failed to unmarshal sequencer details: %v", err)
	}
	var daData types.DADetails
	if err := json.Unmarshal(data.DaDetails, &daData); err != nil {
		log.Fatalf("Failed to unmarshal DA details: %v", err)
	}
	var proverData types.ProverDetails
	if err := json.Unmarshal(data.ProverDetails, &proverData); err != nil {
		log.Fatalf("Failed to unmarshal prover details: %v", err)
	}

	seqDataInd, err := json.MarshalIndent(seqData, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal sequencer details: %v", err)
	}
	daDataInd, err := json.MarshalIndent(daData, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal DA details: %v", err)
	}
	proverDataInd, err := json.MarshalIndent(proverData, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal prover details: %v", err)
	}

	fmt.Printf("SequencerDetails: %s\n", string(seqDataInd))
	fmt.Printf("DaDetails: %s\n", string(daDataInd))
	fmt.Printf("ProverDetails: %s\n", string(proverDataInd))

	return stationId
}
