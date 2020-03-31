# Start higand

# Init higand
higand init higan-test --chain-id=higan-test

# Create a new recorder
higancli keys add bana

# Add genesis account, with coins to the genesis file
higand add-genesis-account $(higancli keys show bana -a) 100000000stake,100000000banacoin

# Generate the transaction that creates your validator
higand gentx --name bana

# Add the generated bonding transaction to the genesis file
higand collect-gentxs

# Start higand
higand start