# ベースイメージとして Node.js を使用
FROM node:18 AS build

# 作業ディレクトリを設定
WORKDIR /app

# package.json と package-lock.json をコピー
COPY package*.json ./

# 依存関係をインストール
RUN npm install

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN npm run build

# Nginxを使用して静的ファイルを配信
FROM nginx:alpine

# Nginxのデフォルトの設定を削除
RUN rm /etc/nginx/conf.d/default.conf

# カスタム設定ファイルをコピー
COPY nginx.conf /etc/nginx/conf.d

# ビルドされた静的ファイルをNginxの公開ディレクトリにコピー
COPY --from=build /app/dist /usr/share/nginx/html

# Nginxを実行
CMD ["nginx", "-g", "daemon off;"]
