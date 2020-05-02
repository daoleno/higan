#!/bin/bash

KEY="bana"
CHAINID="higan-test"
MONIKER="higan"

# remove existing chain environment, data and 
rm -rf ~/.higan*

make install

higancli config keyring-backend test

# if bana exists it should be deleted
higancli keys add $KEY

higand init $MONIKER --chain-id $CHAINID

# Set up config for CLI
higancli config chain-id $CHAINID
higancli config output json
higancli config indent true
higancli config trust-node true

# Allocate genesis accounts (cosmos formatted addresses)
higand add-genesis-account $(higancli keys show $KEY -a) 1000000000000000000stake

# Sign genesis transaction
higand gentx --name $KEY --keyring-backend test

# Collect genesis tx
higand collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
higand validate-genesis

# Command to run the rest server in a different terminal/window
echo -e '\n\nRun this rest-server command in a different terminal/window:'
echo -e "higancli rest-server --laddr \"tcp://localhost:8545\" --unlock-key $KEY --chain-id $CHAINID\n\n"

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
higand start --pruning=nothing --rpc.unsafe --log_level "main:info,state:info,mempool:info"

