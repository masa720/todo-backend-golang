# Todo Backend API (Go)

Go言語とGinフレームワークを使用したTODO管理APIです。JWT認証機能付きで、ユーザーごとのTODO管理が可能です。

## 機能

- ✅ ユーザー認証（JWT）
- ✅ ユーザー登録・ログイン
- ✅ TODOのCRUD操作
- ✅ TODOの完了状態管理
- ✅ パスワードハッシュ化（bcrypt）
- ✅ データベース自動マイグレーション

## 技術スタック

- **言語**: Go 1.21.5
- **フレームワーク**: Gin
- **データベース**: MySQL 8.0
- **ORM**: GORM
- **認証**: JWT
- **パスワードハッシュ**: bcrypt
- **コンテナ**: Docker & Docker Compose

## プロジェクト構造

```
todo-backend-golang/
├── controller/          # HTTPリクエストハンドラー
│   ├── todo_controller.go
│   └── user_controller.go
├── database/            # データベース接続設定
│   └── database.go
├── docs/               # ドキュメント
│   └── login.pu
├── middleware/         # ミドルウェア
│   └── auth.go
├── model/              # データモデル
│   ├── todo.go
│   └── user.go
├── repository/         # データアクセス層
│   ├── todo_repository.go
│   └── user_repository.go
├── router/             # ルーティング設定
│   └── router.go
├── usecase/            # ビジネスロジック
│   ├── todo_usecase.go
│   └── user_usecase.go
├── main.go             # エントリーポイント
├── go.mod              # Goモジュール定義
├── go.sum              # 依存関係チェックサム
├── Dockerfile          # Docker設定
├── docker-compose.yaml # Docker Compose設定
└── .env                # 環境変数（要作成）
```

## セットアップ

### 前提条件

- Go 1.21.5以上
- Docker & Docker Compose
- MySQL 8.0（Dockerで提供）

### 1. リポジトリのクローン

```bash
git clone <repository-url>
cd todo-backend-golang
```

### 2. 環境変数ファイルの作成

プロジェクトルートに`.env`ファイルを作成：

```env
DBMS=localhost
USER=root
PASS=todo-password
PORT=3306
DBNAME=todo
```

### 3. データベースの起動

```bash
docker-compose up -d db
```

### 4. 依存関係のインストール

```bash
go mod download
```

### 5. アプリケーションの起動

#### 開発モード（Air使用）
```bash
# Airのインストール（初回のみ）
go install github.com/cosmtrek/air@latest

# 開発サーバー起動
air
```

#### 通常モード
```bash
go run main.go
```

サーバーは `http://localhost:8080` で起動します。

## API エンドポイント

### 認証不要エンドポイント

| メソッド | エンドポイント | 説明 |
|---------|---------------|------|
| POST | `/api/v1/sign_up` | ユーザー登録 |
| POST | `/api/v1/sign_in` | ログイン |

### 認証必要エンドポイント

| メソッド | エンドポイント | 説明 |
|---------|---------------|------|
| GET | `/api/v1/todos` | TODO一覧取得 |
| POST | `/api/v1/todos` | TODO作成 |
| GET | `/api/v1/todos/:id` | TODO詳細取得 |
| PUT | `/api/v1/todos/:id` | TODO更新 |
| DELETE | `/api/v1/todos/:id` | TODO削除 |
| POST | `/api/v1/todos/:id/done` | TODO完了状態切り替え |

## 使用方法

### 1. ユーザー登録

```bash
curl -X POST http://localhost:8080/api/v1/sign_up \
  -H "Content-Type: application/json" \
  -d '{
    "name": "テストユーザー",
    "mail": "test@example.com",
    "pass": "password123"
  }'
```

### 2. ログイン

```bash
curl -X POST http://localhost:8080/api/v1/sign_in \
  -H "Content-Type: application/json" \
  -d '{
    "mail": "test@example.com",
    "pass": "password123"
  }'
```

レスポンス例：
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### 3. TODO作成（認証が必要）

```bash
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <JWT_TOKEN>" \
  -d '{
    "title": "買い物に行く",
    "description": "スーパーで食材を買う"
  }'
```

### 4. TODO一覧取得

```bash
curl -X GET http://localhost:8080/api/v1/todos \
  -H "Authorization: Bearer <JWT_TOKEN>"
```

## データモデル

### User
```go
type User struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    Name      string         `json:"name"`
    Mail      string         `json:"mail"`
    Pass      string         `json:"pass"`
    CreatedAt time.Time      `json:"createdAt"`
    UpdatedAt time.Time      `json:"updatedAt"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
```

### Todo
```go
type Todo struct {
    ID          uint           `gorm:"primaryKey" json:"id"`
    Title       string         `json:"title"`
    Description string         `json:"description"`
    IsDone      bool           `json:"isDone"`
    CreatedAt   time.Time      `json:"createdAt"`
    UpdatedAt   time.Time      `json:"updatedAt"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
```

## 開発

### ホットリロード

開発時はAirを使用してホットリロードが可能です：

```bash
air
```

### テスト

```bash
go test ./...
```

### ビルド

```bash
go build -o todo-api main.go
```

## Docker

### アプリケーション全体をDockerで起動

```bash
docker-compose up -d
```

### 個別サービス

```bash
# データベースのみ
docker-compose up -d db

# アプリケーションのみ
docker-compose up -d go
```

## セキュリティ

- パスワードはbcryptでハッシュ化
- JWTトークンによる認証
- 認証が必要なエンドポイントの保護
- 環境変数による設定管理

## 注意事項

- 本番環境ではJWTシークレットキーを環境変数から取得するように変更してください
- 必ずHTTPSを使用してセキュリティを確保してください
- データベースのバックアップを定期的に取得してください

## ライセンス

MIT License

## 貢献

プルリクエストやイシューの報告を歓迎します。 