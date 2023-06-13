<!-- 
order: 5
-->

# Events

The merkledrop module emits the following events:
## EventCreate

| Type            | Attribute Key | Attribute Value  |
| :-------------- | :------------ | :--------------- |
| message         | action        | `/incubus.merkledrop.v1beta1.MsgCreate` |
| incubus.merkledrop.v1beta1.EventCreate | owner        | {owner}         |
| incubus.merkledrop.v1beta1.EventCreate | merkledrop_id        | {merkledrop_id}         |

## EventClaim

| Type            | Attribute Key | Attribute Value  |
| :-------------- | :------------ | :--------------- |
| message         | action        | `/incubus.merkledrop.v1beta1.MsgClaim` |
| incubus.merkledrop.v1beta1.EventClaim | merkledrop_id        | {merkledrop_id}         |
| incubus.merkledrop.v1beta1.EventClaim | index        | {index}         |
| incubus.merkledrop.v1beta1.EventClaim | coin        | {coin}         |

## EventWithdraw

| Type                     | Attribute Key | Attribute Value   |
| :----------------------- | :------------ | :---------------- |
| incubus.merkledrop.v1beta1.EventWithdraw | merkledrop_id        | {merkledrop_id}         |
| incubus.merkledrop.v1beta1.EventWithdraw | coin        | {coin}         |