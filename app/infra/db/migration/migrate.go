package main

import (
	db "echo_sprint_planner/app/infra/db"
	model "echo_sprint_planner/app/infra/db/model"
	"fmt"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("マイグレーション成功")
	defer db.CloseDB(dbConn)
	// マイグレーションを実行する
	// 第二引数にマイグレーションしたいテーブルを渡すことができる
	// 順番に注意すること(外部キー制約など)
	dbConn.AutoMigrate(&model.User{}, &model.Sprint{})
}
