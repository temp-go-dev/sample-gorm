package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// gormConnect DBへの接続を取得
func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "user"
	PASS := "password"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "sampledb"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}

	// テーブル名を個別解決
	db.SingularTable(true)

	// ログを有効化（SQLが見れる）
	db.LogMode(true)
	return db
}

func main() {
	db := gormConnect()
	defer db.Close()

	// 全レコード取得
	var todos []Todo
	db.Find(&todos)

	fmt.Println("---------------------")
	fmt.Println(todos)
	fmt.Println("---------------------")
}

// Todo 構造体
type Todo struct {
	ID          string
	UserID      string `gorm:"column:user_id"`
	Start       time.Time
	End         time.Time
	Title       string
	Description string
	Priority    int
}

// Hoge Entity
type Hoge struct {
	ID   int
	Name string
}

// TableName todoテーブル名解決
func (Todo) TableName() string {
	return "todo"
}

// TableName hogesテーブル名解決
func (Hoge) TableName() string {
	return "hoges"
}
