# waco

room status server

## Config

`waco/docker/api/webhook_rules/xxx.yaml`

書き方は他のyamlを参照

### WebHookEvent

- RegisterUserEvent: ユーザ登録
- DeleteUserEvent: ユーザ削除
- PushCurrentUserEvent: 入室
- PopCurrentUserEvent: 退室

### Embedded mark

- `<card_no>`
- `<display_name>`
- `<full_name>`
- `<email>`

これらを埋め込むと，WebHookEventに紐づいたユーザの情報を使うことができる．
（例えば，userAが入室したときの `PushCurrentUserEvent` によって発生するWebHookに， `<card_no>` を書くと，そこがuserAのカード番号に置き換わる）


## usage

### local
```
docker-compose up --build -d
```

### k8s service
dnsname: `waco.eleuth`

