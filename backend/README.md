### seedの追加
  バックエンドコンテナ内で
  ```sh
    go run ./scripts/seed.go
  ```

### カラムの追加（例　users）
  ent/schema/user.go
  Fieldsに新しいカラムを追加

  コードの生成: スキーマファイルを編集した後、entのコードを再生成します
  ```sh
    go generate ./ent
  ```

  マイグレーションの生成: 新しいマイグレーションを生成します
  ```sh
    go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema
  ```

  マイグレーションの適用: データベースにマイグレーションを適用します
  ```sh
    go run -mod=mod entgo.io/ent/cmd/ent migrate exec

### カラムの削除（例　users）
  ent/schema/user.go
    Fieldsに削除したいカラムを消す

  コードの生成: スキーマファイルを編集した後、entのコードを再生成します。
  ```sh
    go generate ./ent
  ```

  マイグレーションの生成: 新しいマイグレーションを生成します
  ```sh
    go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema
  ```

  マイグレーションの適用: データベースにマイグレーションを適用します
  ```sh
    go run -mod=mod entgo.io/ent/cmd/ent migrate exec

### テーブルの追加（例　users）
  新しいエンティティのスキーマファイルを作成: 新しいテーブルに対応するエンティティのスキーマファイルを作成します。

   コードの生成: スキーマファイルを編集した後、entのコードを再生成します。
  ```sh
    go generate ./ent
  ```

  マイグレーションの生成: 新しいマイグレーションを生成します
  ```sh
    go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema
  ```

  マイグレーションの適用: データベースにマイグレーションを適用します
  ```sh
    go run -mod=mod entgo.io/ent/cmd/ent migrate exec