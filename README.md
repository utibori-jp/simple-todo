# Simple-Todoアプリケーション

## 概要
このプロジェクトは、シンプルなTodoアプリケーションです。ユーザーはタスクを追加、編集、削除することができます。バックエンドはGo言語で構築されており、フロントエンドはVue.jsを使用しています。データはPostgreSQLに保存され、Dockerを使用して開発環境を構築しています。

## 機能
- タスクの追加
- タスクの表示
- タスクの編集
- タスクの削除

## 技術スタック
- **バックエンド**: Go, Gorilla Mux, PostgreSQL
- **フロントエンド**: Vue.js, Vue Router
- **スタイル**: CSS
- **コンテナ化**: Docker, Docker Compose
- **Webサーバー**: Nginx

## 環境構築
このアプリケーションをローカル環境で実行するには、DockerとDocker Composeが必要です。

1. リポジトリをクローンします。
   ```bash
   git clone https://github.com/yourusername/todo-app.git
   cd todo-app
   ```
2. Dockerイメージをビルドして、コンテナを起動します。
   ```bash
   docker-compose up --build
   ```
3. ブラウザで``http://localhost:8080``にアクセスしてアプリケーションを使用します。

## APIエンドポイント
- GET /todos:すべてのタスクを取得
- POST /todos:新しいタスクを追加
- PUT /todos/{id}:指定したIDのタスクを更新
- DELETE /todos/{id}:指定したIDのタスクを削除

## repository構成
```csharp
todo-app/
├── backend/              # バックエンドのコード
│   ├── main.go           # エントリーポイント
│   ├── handlers/         # APIハンドラー
│   │   └── todo.go       # Todoに関するハンドラー
│   ├── models/           # Todoモデル
│   │   └── todo.go       # Todoモデル定義
│   ├── database/         # データベース接続関連
│   │   └── db.go         # データベース接続処理
│   └── Dockerfile        # バックエンドのDockerfile
├── frontend/             # フロントエンドのコード
│   ├── src/              # Vue.jsのソースコード
│   │   ├── components/   # Vueコンポーネント
│   │   ├── views/        # ビュー
│   │   ├── App.vue       # アプリケーションのメインコンポーネント
│   │   └── main.js       # エントリーポイント
│   ├── public/           # 公開リソース（index.htmlなど）
│   │   └── index.html    # アプリケーションのHTMLエントリーポイント
│   ├── nginx.conf        # Nginxの設定ファイル
│   ├── package.json      # プロジェクトの依存関係とスクリプトを管理するファイル
│   └── Dockerfile        # フロントエンドのDockerfile
├── db/                   # データベース設定（Docker Compose用）
│   └── init.sql          # 初期データのSQLスクリプト
├── docker-compose.yml    # Docker Compose設定ファイル
├── go.mod                # go.mod
├── go.sum                # go.sum
└── README.md             # プロジェクトの説明書
```
