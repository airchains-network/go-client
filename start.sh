#!/bin/bash

rm -rf data
# rm -rf accounts

echo "Starting Bomber"
go run main.go --numStations 2 --numSchemas 2 --numEngagements 2
