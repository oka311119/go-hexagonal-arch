# go-hexagonal-arch

Hexagonal Architecture (別名: Port and Adapters)のGoによる実装サンプルです。
HTTPサーバを構築し、コマンドによりアプリの操作を行います。

## フォルダ構成

- domain：ビジネスロジックを表すドメインモデルが格納される場所です。ここではそれが更に以下のサブディレクトリに分けられています：
  - entity: ビジネスの中心的存在であるエンティティを表すコードが含まれます。
  - service: エンティティ間の相互作用やビジネスロジックを表すコードが含まれます。
  - valueObject: ビジネス上の意味を持つ値を表すコードが含まれます。
  - port: ドメインが他のレイヤと通信するためのインターフェイスが含まれます。
- adapter：具体的な技術の詳細をカプセル化するアダプタが格納される場所です。これが更に以下のサブディレクトリに分けられています：
  - driver: 外部からドメインに命令を送るためのアダプタが含まれます（例：HTTPリクエストからのデータ変換など）。
  - driven: ドメインからの結果を外部に伝えるためのアダプタが含まれます（例：データベースへのデータ保存など）。

## Running the Application

プロジェクトのルートディレクトリに移動し、以下のコマンドを実行してください。
ローカル:8080ポートでHTTPサーバが起動します。

```bash Copy code
make run
```

### API Endpoints

- Create Todo: POST /create
  - Body: {"id":"1", "title":"My first task"}
- Get Todo by ID: GET /getbyid?id=1
- Get All Todos: GET /getall

ローカルサーバ起動後、以下のコマンドで動作を確認できます。

``` bash
curl -X POST -H "Content-Type: application/json" -d '{"id":"1", "title":"My first task"}' 'http://localhost:8080/create'
```

To get a todo by ID:

``` bash Copy code
Copy code
curl -X GET 'http://localhost:8080/getbyid?id=1'
```

To get all todos:

``` bash Copy code
Copy code
curl -X GET 'http://localhost:8080/getall'
```

## Testing

``` bash Copy code
go test ./...
```

## 実装したいこと

[] データベースへの接続
[] mockgenの実装
[] domain/serviceの修正
[] エラーハンドリング
[] ロギング
[] ユーザ認証
