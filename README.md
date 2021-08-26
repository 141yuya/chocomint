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
- ユーザー取得API(user_controller#Show)
- ユーザー作成API(user_controller#Create)
- ユーザー更新API(user_controller#Update)
- ユーザー削除API(user_controller#Delete)

# ディレクトリ構成

# 未実装
- バリデーション
- ログ出力
- ローカル環境、テスト環境、本番環境などの環境ごとの情報を持つenvの導入
- ユニットテスト
- Gormとは別のマイグレーションライブラリの導入。