# 最新のGoイメージを使用
FROM golang:latest

# コンテナ内の作業ディレクトリを設定
WORKDIR /app

# ローカルファイルをコンテナにコピー
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Goアプリケーションをビルド
RUN go build -o main ./cmd/main.go

# ポート8080を公開
EXPOSE 8080

# Goアプリケーションを実行
CMD ["./main"]
