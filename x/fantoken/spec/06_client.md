<!-- 
order: 6
-->

# Client

## Transactions

The `transactions` commands allow users to `issue`, `mint`, `burn`, `disable minting`, `transfer minting and editing capabilities` for _fan tokens_.

```bash=
incubusd tx fantoken --help
```
### issue

```bash=
incubusd tx fantoken issue \
    --name "fantoken name" \
    --symbol "bitangel" \
    --max-supply 100000000000 \
    --uri "ipfs://...." \
    --from <key-name> -b block --chain-id <chain-id> --fees <fee>
```

### mint

```bash=
incubusd tx fantoken mint [amount][denom] \
    --recipient <address> \
    --from <key-name> -b block --chain-id <chain-id> --fees <fee>
```

### burn

```bash=
incubusd tx fantoken burn [amount][denom] \
    --from <key-name> -b block --chain-id <chain-id> --fees <fee>
```

### set-authority

```bash=
incubusd tx fantoken set-authority [denom] \
    --new-authority <address> \
    --from <key-name> -b block --chain-id <chain-id> --fees <fee>
```

### set-minter

```bash=
incubusd tx fantoken set-minter [denom] \
    --new-minter <address> \
    --from <key-name> -b block --chain-id <chain-id> --fees <fee>
```

### set-uri

```bash=
incubusd tx fantoken set-uri [denom] \
    --uri <uri> \
    --from <key-name> -b block --chain-id <chain-id> --fees <fee>
```

### disable-mint

```bash=
incubusd tx fantoken disable-mint [denom] \
    --from <key-name> -b block --chain-id <chain-id> --fees <fee>
```

## Query

The `query` commands allow users to query the `fantoken` module.

```bash=
incubusd q fantoken --help
```

### denom

```bash=
incubusd q fantoken denom <denom>
```

### authority

```bash=
incubusd q fantoken authority <address>
```

### params

```bash=
incubusd q fantoken params
```