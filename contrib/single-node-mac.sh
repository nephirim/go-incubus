#!/bin/sh

rm -rf ~/.incubusd

set -o errexit -o nounset

# Build genesis file incl account for passed address
incubusd init --chain-id test test
incubusd keys add validator --keyring-backend="test"
incubusd add-genesis-account $(incubusd keys show validator -a --keyring-backend="test") 100000000000000ufury
incubusd gentx validator 100000000ufury --keyring-backend="test" --chain-id test
incubusd collect-gentxs

# Set proper defaults and change ports
sed -i '' 's#"tcp://127.0.0.1:26657"#"tcp://0.0.0.0:26657"#g' ~/.incubusd/config/config.toml
sed -i '' 's/timeout_commit = "5s"/timeout_commit = "1s"/g' ~/.incubusd/config/config.toml
sed -i '' 's/timeout_propose = "3s"/timeout_propose = "1s"/g' ~/.incubusd/config/config.toml
sed -i '' 's/index_all_keys = false/index_all_keys = true/g' ~/.incubusd/config/config.toml

# Start incubus
incubusd start --pruning=nothing
