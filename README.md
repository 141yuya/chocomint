# 概要
RestAPIをクリーンアーキテクチャで作成した。

# 技術スタック
- Go 1.17
- クリーンアーキテクチャ
- フレームワーク:Gin
- ORマッパー:Gorm
- DI:google/wire
- mockライブラリ:gomock

# インストール
```
docker-compose up -d --build
```

# ディレクトリ構成
```
src
├── di
│   ├── wire.go
│   └── wire_gen.go
├── domain
│   ├── entities
│   │   └── user.go
│   └── repositories
│       └── user_repository.go
├── infrastructure
│   ├── router.go
│   └── sqlhandler.go
├── interfaces
│   ├── controllers
│   │   ├── error.go
│   │   └── user_controller.go
│   └── gateways
│       └── user_repository.go
├── main.go
└── usecases
    └── user_usecase.go
```

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


# 未実装
- バリデーション
- ログ出力
- ローカル環境、テスト環境、本番環境などの環境ごとの情報を持つenvの導入
- Gormとは別のマイグレーションライブラリの導入。

# 参考
https://nrslib.com/clean-architecture/

https://www.amazon.co.jp/dp/B07FSBHS2V


