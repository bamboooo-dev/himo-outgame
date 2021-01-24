# HI-MO

## requirements

- Go: 1.15.6
- protoc: >= 3.14.0

## gRPC 関連のツールをインストール

```
make install-grpc
```

## ローカル環境で build

```
make build
```

## docker を使った環境構築

### ビルド

```
docker build -t himo-outgame .
```

### コンテナの立ち上げ

```
docker run -d -p 5502:5502 himo-outgame bin/outgame
```

### docker-compose による開発環境

`outgame-grpcui` サービスが `outgame:5502` に接続確認するので、順番に立ち上げる必要がある  
`fullstorydev/grpcui` イメージには sh すら入ってないので `wait-for-it.sh` がすぐ使えず、とりあえず妥協

```
docker-compose up outgame
docker-compose up outgame-grpcui
```

http://localhost:10080 で gRPC の接続テストができる

## gRPC 接続の GUI による確認

```
go install github.com/fullstorydev/grpcui/cmd/grpcui
grpcui -plaintext localhost:5502
```

## ディレクトリ構成

```
tree
.
├── README.md
├── cmd
│   └── outgame
│       └── main.go
├── docs # ドキュメント
├── go.mod
├── grpc # gRPC の proto ファイルを置くところ
├── internal # このプロジェクト内でのみ使用されるパッケージ
│   ├── domain # ドメインオブジェクトを置く
│   ├── interface # クリーンアーキテクチャの interface(DBやリクエストを扱う層)
│   │   └── handler # リクエストを受け取ってレスポンスを返すものを置く
│   ├── registry # DI コンテナを置く場所
│   └── usecase # クリーンアーキテクチャの usecase (アプリにリクエストしたときどう動いてほしいかのロジックを書くところ)
├── pkg # 他のプロジェクトでも使用できるパッケージを置くところ
└── tools # import を明示的にしていないが go mod tidy で消されてほしくないライブラリを宣言しておくところ
    └── tools.go
```
