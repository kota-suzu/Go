FROM golang:1.16

WORKDIR /app

# go.mod と go.sum ファイルを先にコピー
COPY go.mod go.sum ./

# 必要な依存関係をインストール
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main .

CMD ["./main"]
