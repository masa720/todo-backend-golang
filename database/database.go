package database

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/masa720/todo-backend-golang/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		// .env ファイルが見つからない場合や読み込みに失敗した場合の処理
		panic("Error loading .env file")
	}

	DBMS := os.Getenv("DBMS")
	USER := os.Getenv("USER")
	PASS := os.Getenv("PASS")
	PORT := os.Getenv("PORT")
	DBNAME := os.Getenv("DBNAME")
	CONNECT := USER + ":" + PASS + "@tcp(" + DBMS + ":" + PORT + ")/" + DBNAME + "?charset=utf8mb4&parseTime=true&loc=Local"

	dialector := mysql.Open(CONNECT)
	option := &gorm.Config{}

	if err = dbConnect(dialector, option, 10); err != nil {
		// panic(err)
		log.Fatalln(err)
	}

	autoMigration()

	return db
}

func dbConnect(dialector gorm.Dialector, config gorm.Option, count uint) (err error) {
	// countで指定した回数リトライする
	for count > 1 {
		if db, err = gorm.Open(dialector, config); err != nil {
			time.Sleep(time.Second * 2)
			count--
			log.Printf("retry... count:%v\n", count)
			continue
		}
		break
	}
	// エラーを返す
	return err
}

func autoMigration() {
	db.AutoMigrate(&model.Todo{}, &model.User{})
}
