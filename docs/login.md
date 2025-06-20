# ログイン機能の使用方法

## エンドポイント

### 1. ユーザー登録 (Sign Up)
**POST** `/api/v1/sign_up`

**リクエストボディ:**
```json
{
  "name": "ユーザー名",
  "mail": "user@example.com",
  "pass": "パスワード"
}
```

**レスポンス:**
```json
{
  "message": "User created successfully"
}
```

### 2. ログイン (Sign In)
**POST** `/api/v1/sign_in`

**リクエストボディ:**
```json
{
  "mail": "user@example.com",
  "pass": "パスワード"
}
```

**レスポンス:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

## 認証が必要なエンドポイント

以下のエンドポイントを使用する際は、リクエストヘッダーにJWTトークンを含める必要があります：

**Authorization Header:**
```
Authorization: Bearer <JWT_TOKEN>
```

### 認証が必要なエンドポイント一覧:
- `GET /api/v1/todos` - TODO一覧取得
- `POST /api/v1/todos` - TODO作成
- `GET /api/v1/todos/:id` - TODO詳細取得
- `PUT /api/v1/todos/:id` - TODO更新
- `DELETE /api/v1/todos/:id` - TODO削除
- `POST /api/v1/todos/:id/done` - TODO完了状態切り替え

## 使用例

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

## セキュリティ注意事項

1. **JWTシークレットキー**: 本番環境では環境変数から取得するように変更してください
2. **パスワード**: bcryptでハッシュ化されて保存されます
3. **トークン有効期限**: 24時間に設定されています
4. **HTTPS**: 本番環境では必ずHTTPSを使用してください
