<!-- 
order: 7
-->

# Client

## Transactions

The `transactions` commands allow users to `create` and `claim` for _merkledrops_.

```bash=
incubusd tx merkledrop --help
```
### create

```bash=
incubusd tx merkledrop create [account-file] [output-file] \
	--denom=ufury \
	--start-height=1 \
	--end-height=10 \
	--from=<key-name> -b block --chain-id <chain-id>
```

### claim

```bash=
incubusd tx merkledrop claim [merkledrop-id] \
	--proofs=[proofs-list] \
	--amount=[amount-to-claim] \
	--index=[level-index]
	--from=<key-name> -b block --chain-id <chain-id>
```

## Query

The `query` commands allow users to query the _merkledrop_ module.

```bash=
incubusd q merkledrop --help
```

### detail by id

```bash=
incubusd q merkledrop detail [id]
```

### if index and id have been claimed

```bash=
incubusd q merkledrop index-claimed [id] [index]
```

### params

```bash=
incubusd q merkledrop params
```