package main

import (
	"flag"
	"log"
	"strings"

	"github.com/xinliangnote/go-gin-api/cmd/init/db/mysql"
)

var (
	dbAddr string
	dbUser string
	dbPass string
	dbName string
)

func init() {
	addr := flag.String("addr", "", "请输入 db 地址，例如：127.0.0.1:3306\n")
	user := flag.String("user", "", "请输入 db 用户名\n")
	pass := flag.String("pass", "", "请输入 db 密码\n")
	name := flag.String("name", "", "请输入 db 名称\n")

	flag.Parse()

	dbAddr = *addr
	dbUser = *user
	dbPass = *pass
	dbName = strings.ToLower(*name)
}

func main() {
	// 初始化 DB
	db, err := mysql.New(dbAddr, dbUser, dbPass, dbName)
	if err != nil {
		log.Fatal("new db err: ", err.Error())
	}

	defer func() {
		if err := db.DbClose(); err != nil {
			log.Fatal("db close err: ", err.Error())
		}
	}()

	// 创建 user_demo 表
	err = db.GetDb().Exec(mysql.CreateUserDemoTableSql()).Error
	if err != nil {
		log.Fatal("create user_demo table err: ", err.Error())
	}

	log.Println("create user_demo table success")
}
