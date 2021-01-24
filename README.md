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
