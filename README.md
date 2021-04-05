# waco

room status server

部屋に入りうるユーザの情報と，今の在室状況を保持する．

## Usage

### End Points
| End Point           | Method   | Function             |
|---------------------|----------|----------------------|
| `/users`            | `GET`    | 全ユーザ情報取得     |
| `/users/:cardNo`    | `GET`    | ユーザ情報取得       |
|                     | `POST`   | ユーザ登録           |
|                     | `PUT`    | ユーザ情報更新       |
|                     | `DELETE` | ユーザ削除           |
| `/currents`         | `GET`    | 在室中ユーザ一覧取得 |
| `/currents/:cardNo` | `PUT`    | ユーザ入室           |
|                     | `DELETE` | ユーザ退室           |

### Params

ユーザ登録 / ユーザ情報更新 時はこれらのパラメータをbodyに含めてください

- `card_no` : 学籍番号
- `display_name` : 表示名
- `full_name` : 本名
- `pronunciation` : 呼ばれ方(ひらがな推奨)
- `email` : メールアドレス

とりあえず`card_no`以外はnull許容してます

## WebHook Rules

各種イベント毎にWebHookを複数登録できる．任意のURLに対して`GET`/`POST`リクエストが可能．

ルールは下記の場所にルールとなるyamlをおくだけ

`waco/docker/api/webhook_rules/xxx.yaml`
`waco/webhook_rules(symlink)/xxx.yaml`

```yaml
# example

- name: test-webhook-1
  event: RegisterUserEvent
  url: http://hogefuga
  method: GET
  params:
      hoge: fuga
      foo: bar
- name: test-webhook-2
  event: RegisterUserEvent
  url: http://hogefuga
  method: POST
  params:
      hoge: fuga
      foo: bar
```

### WebHookEvent

- RegisterUserEvent: ユーザ登録
- UpdateUserEvent:  ユーザ更新
- DeleteUserEvent: ユーザ削除
- PushCurrentUserEvent: 入室
- PopCurrentUserEvent: 退室

### Embedded mark

- `<card_no>`
- `<display_name>`
- `<full_name>`
- `<pronunciation>`
- `<email>`

これらを埋め込むと，WebHookEventに紐づいたユーザの情報を使うことができる．
（例えば，userAが入室したときの `PushCurrentUserEvent` によって発生するWebHookに， `<card_no>` を書くと，そこがuserAのカード番号に置き換わる）

```yaml
# lab入室時のWebHook
- name: notice-enter
  event: PushCurrentUserEvent
  url: http://nenech.eleuth/accessing/enter
  method: GET
  params:
      message: ":rocket: Hi *<display_name>*"
```

`<display_name>` の部分が入室したuserのdisplay_nameに起き変わる


## Setup

### local
```
docker-compose up --build -d
```

### k8s
refs: [./.k8s/](./.k8s/)

