package main

import (
	"database/sql"
	"fmt"
	"initDB"

	_ "github.com/go-sql-driver/mysql" //init()导入驱动，但不使用
)

// 开始事务func (db *DB) Begin() (*Tx, error)
// 提交事务func (tx *Tx) Commit() error
// 回滚事务func (tx *Tx) Rollback() error
type user struct {
	id, age int
	name    string
}

func transaction(db *sql.DB) {
	tx, err := db.Begin() //开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlstrinsert := "insert into user values(?,?,?)"
	_, err = tx.Exec(sqlstrinsert, 9, "lebron", 38)
	if err != nil {
		tx.Rollback() //回滚
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	fmt.Println("Insert success")

	sqlstrdelete := "delete from user where id=?"
	_, err = tx.Exec(sqlstrdelete, 30)
	if err != nil {
		tx.Rollback() //回滚事务
		fmt.Printf("delete failed,err:%v\n", err)
		return
	}
	fmt.Println("delete success")

	tx.Commit() //提交事务
	fmt.Println("事务成功提交了")
}
func main() {
	db, _ := initDB.InitDB()
	transaction(db)
}
