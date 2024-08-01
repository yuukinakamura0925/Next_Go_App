# ベースイメージ
FROM node:18-alpine

# 作業ディレクトリを設定
WORKDIR /app

# package.json と package-lock.json をコピー
COPY package.json package-lock.json ./

# 依存関係をインストール
RUN npm install --legacy-peer-deps

# アプリケーションのソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN npm run build

# ポートを指定
EXPOSE 3000

# アプリケーションを起動
CMD ["npm", "start"]

# 本番では上記のようにbuildした.nextファイルを作成し、npm startで起動するようにする
# 開発環境では、差分を即座に反映させるたいので、ローカルでpnpm devを実行するのがよい。なので、Dockerfileもそれに合わせてvolumeでのマウントをコメントアウトしています