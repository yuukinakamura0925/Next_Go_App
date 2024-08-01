# 最新のGoイメージを使用
FROM golang:latest

# コンテナ内の作業ディレクトリを設定
WORKDIR /app

# ローカルファイルをコンテナにコピー
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# airをインストール
RUN go install github.com/air-verse/air@latest

# バイナリをビルド
RUN go build -o main ./cmd/main.go

# ポート8080を公開
EXPOSE 8080

# airを実行
CMD ["air"]
