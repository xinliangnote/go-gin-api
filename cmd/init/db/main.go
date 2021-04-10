package main

import (
	"flag"
	"fmt"
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

	// 开启事务
	tx := db.GetDb().Begin()

	// 创建 user_demo 表
	err = tx.Exec(mysql.CreateUserDemoTableSql()).Error
	if err != nil {
		tx.Rollback()
		log.Fatal("create user_demo table err: ", err.Error())
	}
	fmt.Println("create user_demo table success")

	// 创建 authorized 表
	err = tx.Exec(mysql.CreateAuthorizedTableSql()).Error
	if err != nil {
		tx.Rollback()
		log.Fatal("create authorized table err: ", err.Error())
	}
	fmt.Println("create authorized table success")

	err = tx.Exec(mysql.CreateAuthorizedTableDataSql()).Error
	if err != nil {
		tx.Rollback()
		log.Fatal("create authorized table data err: ", err.Error())
	}
	fmt.Println("create authorized table data success")

	// 创建 authorized_api 表
	err = tx.Exec(mysql.CreateAuthorizedAPITableSql()).Error
	if err != nil {
		tx.Rollback()
		log.Fatal("create authorized_api table err: ", err.Error())
	}
	fmt.Println("create authorized_api table success")

	err = tx.Exec(mysql.CreateAuthorizedAPITableDataSql()).Error
	if err != nil {
		tx.Rollback()
		log.Fatal("create authorized_api table data err: ", err.Error())
	}
	fmt.Println("create authorized_api table data success")

	// 创建 admin 表
	err = tx.Exec(mysql.CreateAdminTableSql()).Error
	if err != nil {
		tx.Rollback()
		log.Fatal("create admin table err: ", err.Error())
	}
	fmt.Println("create admin table success")

	err = tx.Exec(mysql.CreateAdminTableDataSql()).Error
	if err != nil {
		tx.Rollback()
		log.Fatal("create admin table data err: ", err.Error())
	}
	fmt.Println("create admin table data success")

	// 完成事务
	tx.Commit()

}
