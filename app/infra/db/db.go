package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB { // *gorm.DBを返す関数
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}

	// DB接続
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	// 第二引数にgorm.Config{}を渡すことで、デフォルトの設定を使用することができる
	// 今回は、gormのデフォルトの設定を使用する
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("DB接続成功")
	return db
}

// DB接続をクローズする
func CloseDB(db *gorm.DB) { // *gorm.DBを引数に取る関数
	dbSQL, _ := db.DB()
	if err := dbSQL.Close(); err != nil {
		log.Fatalln(err)
	}
}
