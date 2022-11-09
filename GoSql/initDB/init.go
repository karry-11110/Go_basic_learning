package initDB //按惯例和文件夹名相同

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //init()导入驱动，但不使用
)

//定义一个全局对象db
var db *sql.DB //sql.DB是表示连接的数据库对象（结构体实例），它保存了连接数据库相关的所有信息。它内部维护着一个具有零到多个底层连接的连接池，它可以安全地被多个goroutine同时使用

//定义一个初始化数据库函数
func InitDB() (db *sql.DB, err error) {

	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True" //连接信息
	db, err = sql.Open("mysql", dsn)                                                 //给全局变量赋值，不要用简短变量,一定要导入驱动，否则mysql识别不了
	if err != nil {
		return db, err
	}
	err = db.Ping() // 尝试与数据库建立连接（校验dsn是否正确）
	if err != nil {
		return db, err
	}
	return db, nil //没有错误的话说明连接成功，返回一个nil代表没有错误

}

// func main() {
// 	_, err := InitDB()
// 	if err != nil {
// 		fmt.Printf("init db failed,err:%v\n", err)
// 		return
// 	}

// }
