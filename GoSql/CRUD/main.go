package main

import (
	"fmt"
	"initDB"
)

type user struct {
	id, age int
	name    string
}

func main() {

	//构造数据库连接对象
	db, _ := initDB.InitDB()

	//单行查询   func (db *DB) QueryRow(query string, args ...interface{}) *Row
	var u user                                                  //构造结构体对象，以供查询数据写入
	sqlstr := "select id, name, age from user where id=?"       //单行查询信息表示，？代表queryrow(信息表示，参数)
	err := db.QueryRow(sqlstr, 30).Scan(&u.id, &u.name, &u.age) // 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	if err != nil {
		fmt.Printf("scan failed,err:%v\n", err)
		return
	}
	fmt.Printf("id:%d,name:%s,age:%d\n", u.id, u.name, u.age) //id:30,name:curry,age:33

	//多行查询  func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
	// sqlmultistr := "select id,name,age from user where id>?"
	// rows, err := db.Query(sqlmultistr, 10) //注意和单行查询返回值不一样，rows此时存的是查询结果集，同时也要调用scan方法
	// if err != nil {
	// 	fmt.Printf("query failed,err:\n", err)
	// 	return
	// }
	// defer rows.Close() // 非常重要：关闭rows释放持有的数据库链接
	// for rows.Next() {  // 循环读取结果集中的数据
	// 	var u user //构造结构体对象，以供查询数据写入
	// 	err := rows.Scan(&u.id, &u.name, &u.age)
	// 	if err != nil {
	// 		fmt.Printf("scan failed,err:%v\n", err)
	// 		return
	// 	}
	// 	fmt.Printf("id:%d,name:%s,age:%d\n", u.id, u.name, u.age)
	// 	// id:11,name:klay,age:29
	// 	// id:23,name:green,age:29
	// 	// id:24,name:kobe,age:42
	// 	// id:30,name:curry,age:33
	// }

	// //插入数据  插入、更新和删除操作都使用Exec方法func (db *DB) Exec(query string, args ...interface{}) (Result, error)Result是对已执行的SQL命令的总结
	// sqlinsertsql := "insert into user(name,age)values(?,?)"
	// retis, err := db.Exec(sqlinsertsql, "durant", 31)
	// if err != nil {
	// 	fmt.Printf("insert failed,err:%v\n", err)
	// 	return
	// }
	// retisId, err := retis.LastInsertId() // 新插入数据的id
	// if err != nil {
	// 	fmt.Printf("get lastinsert id failed,err:%v\n", err)
	// 	return
	// }
	// fmt.Printf("insert success,the id is %d.\n", retisId)

	// //更新数据
	// sqlupdatestr := "update user set id=? where name=?"
	// retud, err := db.Exec(sqlupdatestr, 7, "durant")
	// if err != nil {
	// 	fmt.Printf("insert failed,err:%v\n", err)
	// 	return
	// }
	// n, err := retud.RowsAffected() // 操作影响的行数
	// if err != nil {
	// 	fmt.Printf("get RowsAffected failed, err:%v\n", err)
	// 	return
	// }
	// fmt.Printf("update success, affected rows:%d\n", n)

	//删除数据
	// 删除数据
	// sqldelstr := "delete from user where id = ?"
	// retdl, err := db.Exec(sqldelstr, 24)
	// if err != nil {
	// 	fmt.Printf("delete failed, err:%v\n", err)
	// 	return
	// }
	// n, err := retdl.RowsAffected() // 操作影响的行数
	// if err != nil {
	// 	fmt.Printf("get RowsAffected failed, err:%v\n", err)
	// 	return
	// }
	// fmt.Printf("delete success, affected rows:%d\n", n)
}
