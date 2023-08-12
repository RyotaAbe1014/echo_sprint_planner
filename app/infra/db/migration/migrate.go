package main

import (
	db "echo_sprint_planner/app/infra/db"
	model "echo_sprint_planner/app/infra/db/models"
	"fmt"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("マイグレーション成功")
	defer db.CloseDB(dbConn)
	// マイグレーションを実行する
	// 第二引数にマイグレーションしたいテーブルを渡すことができる
	// 順番に注意すること(外部キー制約など)

	// 追加: PostgreSQL に拡張機能 "uuid-ossp" を追加する
	// これによって uuid_generate_v4() が使えるようになる
	dbConn.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	dbConn.AutoMigrate(&model.User{}, &model.Sprint{}, &model.ProductBacklog{})
}
