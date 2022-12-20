# テスト用データベースに同じテーブルをマイグレーションする
mysql -uroot -proot earthquake-alert-test < "/docker-entrypoint-initdb.d/001_schema.sql"
