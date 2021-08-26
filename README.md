# 概要
RestAPIをクリーンアーキテクチャで作成した。

# 技術スタック
- Go 1.17
- クリーンアーキテクチャ
- Gin (フレームワーク)
- Gorm(ORマッパー)
- google/wire(DI)

# 機能一覧
- ユーザー一覧取得API(user_controller#Index)
```
curl -X GET "http://localhost:8080/users"
```

- ユーザー取得API(user_controller#Show)
```
curl -X GET "http://localhost:8080/users/1"
```

- ユーザー作成API(user_controller#Create)
```
curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"FirstName": "Yuya", "LastName": "Ishii"}' localhost:8080/users
```

- ユーザー更新API(user_controller#Update)
```
curl -i -H "Accept: application/json" -H "Content-type: application/json" -X PUT -d '{"FirstName": "Yuya2", "LastName": "Ishii2"}' localhost:8080/users/1
```

- ユーザー削除API(user_controller#Delete)
```
curl -X DELETE "http://localhost:8080/users/1"
```

# ディレクトリ構成

# 未実装
- バリデーション
- ログ出力
- ローカル環境、テスト環境、本番環境などの環境ごとの情報を持つenvの導入
- ユニットテスト
- Gormとは別のマイグレーションライブラリの導入。

# 参考
https://nrslib.com/clean-architecture/

https://www.amazon.co.jp/dp/B07FSBHS2V


