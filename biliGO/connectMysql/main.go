package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = "3306"
	user     = "root"
	password = "root"
	database = "test"
)

func main() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
	var err error
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(db)
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("Connected to MySQL")

	// 查询一笔
	a, err := getOne(2)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(a)
	a.userName = "lisi01"
	a.userAge = 21

	// 查询多行
	many, err := getMany(0)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(many)

	// 修改
	err = a.update()
	if err != nil {
		log.Fatalln(err.Error())
	}
	a1, _ := getOne(2)
	fmt.Println(a1)
}
