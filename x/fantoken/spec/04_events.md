<!-- 
order: 4
-->

# Events

The fantoken module emits the following events:
## EventIssue

| Type            | Attribute Key | Attribute Value  |
| :-------------- | :------------ | :--------------- |
| message         | action        | `/incubus.fantoken.v1beta1.MsgIssue` |
| incubus.fantoken.v1beta1.EventIssue | denom        | {denom}         |

## EventDisableMint

| Type            | Attribute Key | Attribute Value  |
| :-------------- | :------------ | :--------------- |
| message         | action        | `/incubus.fantoken.v1beta1.MsgDisableMint` |
| incubus.fantoken.v1beta1.EventDisableMint | denom        | {denom}         |

## EventMint

| Type                     | Attribute Key | Attribute Value   |
| :----------------------- | :------------ | :---------------- |
| message         | action        | `/incubus.fantoken.v1beta1.MsgMint` |
| incubus.fantoken.v1beta1.EventMint | recipient        | {recipient}         |
| incubus.fantoken.v1beta1.EventMint | coin        | {coin}         |

## EventBurn

| Type           | Attribute Key | Attribute Value    |
| :------------- | :------------ | :----------------- |
| message         | action        | `/incubus.fantoken.v1beta1.MsgBurn` |
| incubus.fantoken.v1beta1.EventBurn | sender        | {sender}         |
| incubus.fantoken.v1beta1.EventBurn | coin        | {coin}         |

## EventSetAuthority

| Type           | Attribute Key | Attribute Value |
| :------------- | :------------ | :-------------- |
| message         | action        | `/incubus.fantoken.v1beta1.MsgSetAuthority` |
| incubus.fantoken.v1beta1.EventTransferAuthority | denom        | {denom}         |
| incubus.fantoken.v1beta1.EventTransferAuthority | old_authority        | {old_authority}         |
| incubus.fantoken.v1beta1.EventTransferAuthority | new_authority        | {new_authority}         |

## EventSetMinter

| Type           | Attribute Key | Attribute Value |
| :------------- | :------------ | :-------------- |
| message         | action        | `/incubus.fantoken.v1beta1.MsgSetMinter` |
| incubus.fantoken.v1beta1.EventTransferMinter | denom        | {denom}         |
| incubus.fantoken.v1beta1.EventTransferMinter | old_minter        | {old_minter}         |
| incubus.fantoken.v1beta1.EventTransferMinter | new_authority        | {new_minter}         |

## EventSetUri

| Type            | Attribute Key | Attribute Value  |
| :-------------- | :------------ | :--------------- |
| message         | action        | `/incubus.fantoken.v1beta1.MsgSetUri` |
| incubus.fantoken.v1beta1.EventSetUri | denom        | {denom}         |
