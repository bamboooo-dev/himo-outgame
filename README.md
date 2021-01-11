# HI-MO

## requirements

- Go: 1.15.6
- protoc: >= 3.14.0

## gRPC 関連のツールをインストール

```
make install-grpc
```

## build

```
make build
```

## ディレクトリ構成

```
├── README.md
├── cmd
│   └── outgame
│       └── main.go
├── docs
├── go.mod
├── internal
│   ├── domain
│   ├── interface
│   ├── registry
│   └── usecase
└── tools
    └── tools.go
```
