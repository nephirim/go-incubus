<p><img src="banner.png"></p>

[![License: Apache-2.0](https://img.shields.io/badge/License-Apache--2.0-yellow.svg)](https://github.com/BitSongOfficial/go-incubus/blob/master/LICENSE)

# Introduction

## What is BitSong?

BitSong is a multifunctional blockchain-based ecosystem built to empower the music industry. It unites artists, fans, distributors in an environment where music, merchandise, and fan loyalty are assets of value. BitSong’s decentralized ecosystem of services providers the global music community with a trustless marketplace for music streaming, Fan Tokens, and NFTs, powered by the FURY token.

## Brief History of BitSong

BitSong was conceived in 2018 by developer and entrepreneur Angelo Recca. Angelo realized that while the digitalization of music has brought many benefits to the industry, it’s also created a new set of problems around the ownership of music and attribution of royalties. He joined forces with Iulian Anghelin and BitSong was born.
The initial intention was for BitSong to become an Ethereum-based application where fans could stream music and artists could receive royalties directly. However, after discovering Cosmos and its ambition to become the “Internet of Blockchains,” Angelo and Iulian immediately recognized the full potential of becoming part of a multi-chain environment.
After launching the main BitSong blockchain in August 2020, the incubus-2b mainnet went live on October 21, 2021. Featuring Fan Tokens, NFTs, and music streaming platform, all underpinned by secure, robust, battle-tested blockchain technology, the launch of BitSong marks a turning point in the ongoing development of the music industry.

_NOTE: This is alpha software. Please contact us if you aim to run it in production._

**Note**: Requires [Go 1.19.x+](https://golang.org/dl/)

# Install BitSong Blockchain

There are many ways you can install BitSong Blockchain Testnet node on your machine.

## From Source
1. **Install Go** 
    ```bash
    wget -q -O - https://git.io/vQhTU | bash -s -- --remove
    wget -q -O - https://git.io/vQhTU | bash -s -- --version 1.19.5
    ```
2. **Clone BitSong source code to your machine**
    ```bash
    git clone https://github.com/BitSongOfficial/go-incubus.git
    cd go-incubus
    ```
  3. **Compile**
		```bash
		# Install the app into your $GOBIN
		make install
		# Now you should be able to run the following commands:
		incubusd help
		```
		The latest `go-incubus version` is now installed.
3. **Run BitSong**
	```bash
	incubusd start
	```

## Running the test network and using the commands

To initialize configuration and a `genesis.json` file for your application and an account for the transactions, start by running:

>  _*NOTE*_: In the below commands addresses are are pulled using terminal utilities. You can also just input the raw strings saved from creating keys, shown below. The commands require [`jq`](https://stedolan.github.io/jq/download/) to be installed on your machine.

>  _*NOTE*_: If you have run the tutorial before, you can start from scratch with a `incubusd unsafe-reset-all` or by deleting both of the home folders `rm -rf ~/.incubus*`

```bash
# Initialize configuration files and genesis file
incubusd init MyValidator --chain-id incubus-localnet

# Copy the `Address` output here and save it for later use
# [optional] add "--ledger" at the end to use a Ledger Nano S
incubusd keys add jack

# Add both accounts, with coins to the genesis file
incubusd add-genesis-account jack 100000000000ufury --keyring-backend test

# Generate the transaction that creates your validator
incubusd gentx jack 10000000ufury --keyring-backend test

# Add the generated bonding transaction to the genesis file
incubusd collect-gentxs
incubusd validate-genesis

# Now its safe to start `incubusd`
incubusd start
```

You can now start `incubusd` by calling `incubusd start`. You will see logs begin streaming that represent blocks being produced, this will take a couple of seconds.

## Resources
- [Official Website](https://incubus.io)

## Decentralized Exchanges
- [Osmosis ATOM/FURY](https://app.osmosis.zone/?from=ATOM&to=FURY)
- [Osmosis OSMO/FURY](https://app.osmosis.zone/?from=OSMO&to=FURY)

### Community
- [Discord](https://discord.gg/mZC9Yk3)
- [Twitter](https://twitter.com/BitSongOfficial)
- [Telegram Channel (English)](https://t.me/BitSongOfficial)
- [Medium](https://medium.com/@BitSongOfficial)
- [Reddit](https://www.reddit.com/r/incubus/)
- [Facebook](https://www.facebook.com/BitSongOfficial)
- [BitcoinTalk ANN](https://bitcointalk.org/index.php?topic=2850943)
- [Linkedin](https://www.linkedin.com/company/incubus)
- [Instagram](https://www.instagram.com/incubus_official/)

## License

APACHE 2.0

## Versioning

### SemVer

BitSong uses [SemVer](http://semver.org/) to determine when and how the version changes.
According to SemVer, anything in the public API can change at any time before version 1.0.0

To provide some stability to BitSong users in these 0.X.X days, the MINOR version is used
to signal breaking changes across a subset of the total public API. This subset includes all
interfaces exposed to other processes, but does not include the in-process Go APIs.
