<!-- 参考文献 -->
https://qiita.com/takakou/items/a01af0515f49e90bd05c


# Next Go App

このプロジェクトは、Next.jsを使用したフロントエンドとGoを使用したバックエンドを持つWebアプリケーションです。また、データ保存にはPostgreSQLデータベースを使用しています。

## はじめに

### 必要条件

- Docker
- Docker Compose

### インストール

1. リポジトリをクローンします:

  ```bash
   git clone https://github.com/yuukinakamura0925/Next_Go_App.git
   cd Next_Go_App
  ```
2. Dockerコンテナを起動します:

```bash
docker-compose up --build
```

3. ブラウザでhttp://localhost:3000（フロントエンド）およびhttp://localhost:8080（バックエンドAPI）にアクセスします。

# プロジェクト構成
frontend/: Next.jsフロントエンドアプリケーションが含まれています。
backend/: Goバックエンドアプリケーションが含まれています。
db/: PostgreSQLデータベース。

# Docker構成
## サービス
- backend:


  - ビルドコンテキスト: ./backend/
  - Dockerfile: Dockerfile
  - ポート: 8080:8080
  - ボリューム: ./backend:/app
  - 依存関係: db
  - ベースイメージ: golang:latest

- frontend:

  - ビルドコンテキスト: ./frontend/
  - Dockerfile: Dockerfile
  - ボリューム: ./frontend:/app
  - ポート: 3000:3000
  - ベースイメージ: node:18-alpine

- db:

  - イメージ: postgres:15-alpine
  - ポート: 5432:5432
  - 環境変数:
  - POSTGRES_PASSWORD=password
  - POSTGRES_DB=db
  - ボリューム: db-store:/var/lib/postgresql/data
  - ボリューム
  - db-store: PostgreSQLデータベースデータの永続化ストレージ。
  - Dockerfile
  - バックエンドのDockerfile
  - Dockerfile
  - コードをコピーする

```
# 最新のGoイメージを使用
FROM golang:latest

# コンテナ内の作業ディレクトリを設定
WORKDIR /app

# ローカルファイルをコンテナにコピー
COPY . .

# 必要なGoパッケージをダウンロード
RUN go mod download

# Goアプリケーションをビルド
RUN go build -o main .

# ポート8080を公開
EXPOSE 8080

# Goアプリケーションを実行
CMD ["./main"]
フロントエンドのDockerfile
Dockerfile
コードをコピーする
# Node.js 18-alpineイメージを使用
FROM node:18-alpine

# コンテナ内の作業ディレクトリを設定
WORKDIR /app

# キャッシュ利用で効率化するためpackage.jsonとpackage-lock.jsonをコピー
COPY package.json package-lock.json ./

# ローカルファイルをコンテナにコピー
COPY . .

# 依存関係をインストール
RUN npm install

# Next.jsアプリケーションをビルド
RUN npm run build

# ポート3000を公開
EXPOSE 3000

# Next.jsアプリケーションを開発モードで起動
CMD ["npm","run","dev"]
```