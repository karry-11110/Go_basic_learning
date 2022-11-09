package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
type user struct {
	Id   int
	Name string
	Age  int
}
var db *sqlx.DB
//查询单行数据
func queryRow(db *sqlx.DB) {
	sqlstr := "select id,name,age from user where id=?" //查询语句信息str
	var u user                                          //定义结构体对象
	err := db.Get(&u, sqlstr, 6)                        //调用参数数据库对象，执行查询单行方法
	// 结构体user的访问限制改为全局（即大写)。 这里我们可以理解了，当db.Get(&u,sql,id)执行时，
	// sqlx包中会访问结构体u中的各字段，这时发现字段全部为小写，不可访问，即报错了。我们修改为大写即解决了问题。
	if err != nil {
		fmt.Printf("get failed ,err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", u)
	// fmt.Printf("id:%d,name:%s,age:%d\n", u.d, u.name, u.age) //输出查询信息

}

//查询多行数据
func queryMultiRow(db *sqlx.DB) {
	sqlstr := "select id,name,age from user where id>?" //查询语句信息str
	var users []user                                    //定义结构体对象切片
	err := db.Select(&users, sqlstr, 11)                //调用参数数据库对象，执行查询单行方法
	if err != nil {
		fmt.Printf("get failed ,err:%v\n", err)
		return
	}
	fmt.Printf("users:%3v\n", users) //输出查询信息
}

// //插入数据
// func insert(db *sqlx.DB) {
// 	sqlstr := "insert into user(name, age) values (?,?)" //修改myslq语句定义
// 	ret, err := db.Exec(sqlstr, "wangkun", 22)           //调用参数数据库对象执行修改方法,返回结果体
// 	if err != nil {
// 		fmt.Printf("insert failed,err:%d\n", err)
// 		return
// 	}
// 	theid, err := ret.LastInsertId() //调用结果体，输出新插入的数据
// 	if err != nil {
// 		fmt.Printf("get id failed,err:%d\n", err)
// 		return
// 	}
// 	fmt.Printf("insert success,the id is :%d\n", theid)
// }

// //更新数据
// func update(db *sqlx.DB) {
// 	sqlstr := "update user set age=? where id = ?" //修改myslq语句定义
// 	ret, err := db.Exec(sqlstr, 34, 30)            //调用参数数据库对象执行修改方法,返回结果体
// 	if err != nil {
// 		fmt.Printf("update failed,err:%d\n", err)
// 		return
// 	}
// 	n, err := ret.RowsAffected() //调用结果体，输出新插入的数据
// 	if err != nil {
// 		fmt.Printf("get n failed,err:%d\n", err)
// 		return
// 	}
// 	fmt.Printf("update success,the affacted rows is :%d\n", n)
// }

// //删除数据
// func delete(db *sqlx.DB) {
// 	sqlstr := "delete from user where name = ?" //修改myslq语句定义
// 	ret, err := db.Exec(sqlstr, "wangkun")      //调用参数数据库对象执行修改方法,返回结果体
// 	if err != nil {
// 		fmt.Printf("delete failed,err:%d\n", err)
// 		return
// 	}
// 	n, err := ret.RowsAffected() //调用结果体，输出新插入的数据
// 	if err != nil {
// 		fmt.Printf("get n failed,err:%d\n", err)
// 		return
// 	}
// 	fmt.Printf("delete success,the affacted rows is :%d\n", n)
// }

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/gorm_test"
	db, _ := sqlx.Connect("mysql", dsn)\
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	queryRow(db)
}
