package cli

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/airchains-network/junction-go-client/core/trackgate"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
	"github.com/urfave/cli/v2"
)

func Bomber(ctx context.Context, client cosmosclient.Client, account cosmosaccount.Account, addr string, numStations int, numSchemas int, numEngagements int) {
	for i := 0; i < numStations; i++ {
		stationName := fmt.Sprintf("testing-station-%d", i+1)
		stationId := trackgate.InitStation(addr, client, ctx, account, stationName)
		for j := 0; j < numSchemas; j++ {
			trackgate.InitSchema(client, ctx, account, stationId, addr)
			for k := 0; k < numEngagements; k++ {
				trackgate.InitSchemaEngagement(client, ctx, account, stationId, addr)
			}
		}
		// Reset pod number to 1 after each station
		err := trackgate.ResetPodNumberToFile()
		if err != nil {
			log.Fatalf("failed to reset pod number to file: %v", err)
		}
	}
}

func Trackgate(ctx context.Context, client cosmosclient.Client, account cosmosaccount.Account, addr string) {
	app := &cli.App{
		Name:  "Bomber CLI",
		Usage: "A CLI app to initialize stations, schemas, and engagements",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "numStations",
				Usage:    "Number of stations to initialize",
				Required: true,
			},
			&cli.IntFlag{
				Name:     "numSchemas",
				Usage:    "Number of schemas to initialize per station",
				Required: true,
			},
			&cli.IntFlag{
				Name:     "numEngagements",
				Usage:    "Number of engagements to initialize per schema",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			numStations := c.Int("numStations")
			numSchemas := c.Int("numSchemas")
			numEngagements := c.Int("numEngagements")

			Bomber(ctx, client, account, addr, numStations, numSchemas, numEngagements)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
