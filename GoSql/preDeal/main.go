package main

// 预处理执行过程：
// 把SQL语句分成两部分，命令部分与数据部分。
// 先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理。
// 然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换。
// MySQL服务端执行完整的SQL语句并将结果返回给客户端。
// why
// 优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
// 避免SQL注入问题。//我们任何时候都不应该自己拼接SQL语句！
import (
	"database/sql"
	"fmt"
	"initDB"

	_ "github.com/go-sql-driver/mysql" //init()导入驱动，但不使用
)

type user struct {
	id, age int
	name    string
}

//预处理查询
//查找
func preQuery(db *sql.DB) {
	sqlstr := "select id,name,age from user where id>?"
	stmt, err := db.Prepare(sqlstr) // 先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理。
	if err != nil {
		fmt.Printf("prepare failed,err:%v\n", err)
		return
	}
	defer stmt.Close() //一定要注意别忘记关闭数据库连接
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed,err:&v\n", err)
		return
	}
	defer rows.Close() // 非常重要：关闭rows释放持有的数据库链接

	for rows.Next() { // 循环读取结果集中的数据
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed ,err:%v\n", err)
			return
		}
		fmt.Printf("id:%d,name:%s,age:%d\n", u.id, u.name, u.age)
	}

}

//插入
func preInsert(db *sql.DB) {
	sqlstr := "insert into user values(?,?,?)"
	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("prepare failed ,err:%v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(33, "lilard", 33)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	fmt.Println("Insert success")

}

//删除
func preDelete(db *sql.DB) {
	sqlstr := "delete from user where id=?"
	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("prepare failed ,err:%v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(32)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return
	}
	fmt.Println("delete success")

}

//更新
//删除
func preUpdate(db *sql.DB) {
	sqlstr := "update user set age=? where id=?"
	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("prepare failed ,err:%v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(66, 33)
	if err != nil {
		fmt.Printf("update failed,err:%v\n", err)
		return
	}
	fmt.Println("update success")

}
func main() {
	db, _ := initDB.InitDB()
	preQuery(db)
	preInsert(db)
	preDelete(db)
	preUpdate(db)
}
