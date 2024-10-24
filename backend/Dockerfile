# ベースイメージとして golang を使用
FROM golang:1.18 AS build

# 作業ディレクトリを設定
WORKDIR /app

# モジュールをコピー
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY ./backend/ .

# アプリケーションをビルド
RUN GOOS=linux GOARCH=amd64 go build -o main .

# 最終イメージを作成
FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/main .

# 実行コマンド
CMD ["./main"]
