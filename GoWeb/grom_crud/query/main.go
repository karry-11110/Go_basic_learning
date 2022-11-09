package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//1.定义模型
type NbaInfo struct {
	Id   string
	Name string `gorm:"default:'wangkun'"`

	Age int
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True") //连接mysql数据库
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&NbaInfo{})
	// db.Create(&NbaInfo{"35", "durant", 33})
	// db.Create(&NbaInfo{"7", "", 23})
	// db.Create(&NbaInfo{"24", "kobe", 43})
	// db.Create(&NbaInfo{"11", "klay", 30})
	// db.Create(&NbaInfo{"23", "green", 30})
	//

	//一般查询
	// var (
	// 	u1 NbaInfo
	// 	u2 NbaInfo
	// 	u3 NbaInfo
	// 	u4 []NbaInfo
	// 	u5 NbaInfo
	// )
	// db.First(&u1)       //根据主键查询第一条记录 SELECT * FROM nabinfos ORDER BY id LIMIT 1;
	// db.Take(&u2)        //随机获取一条记录 SELECT * FROM nabinfos LIMIT 1;
	// db.Last(&u3)        //根据主键获取最后一条记录  SELECT * FROM nabinfos ORDER BY id DESC LIMIT 1;
	// db.Find(&u4)        //查询所有记录   SELECT * FROM nabinfos;
	// db.First(&u5, "30") //查询指定的某条记录   SELECT * FROM nabinfos WHERE id="30";
	// fmt.Printf("u1:%v\nu2:%v\nu3:%v\nu4:%v\nu5:%v\n", u1, u2, u3, u4, u5)

	//where条件
	// var (
	// 	u1 NbaInfo
	// 	u2 []NbaInfo
	// 	u3 []NbaInfo
	// 	u4 []NbaInfo
	// 	u5 []NbaInfo
	// 	u6 []NbaInfo
	// 	//u7 []NbaInfo
	// 	u8    []NbaInfo
	// 	u66   []NbaInfo
	// 	u666  []NbaInfo
	// 	u6666 []NbaInfo
	// 	u9    []NbaInfo
	// )
	// db.Where("age = ?", 33).First(&u1)                 //查询第一条匹配到的记录  SELECT * FROM nbainfos WHERE age=33 limit 1;
	// db.Where("age = ?", 33).Find(&u2)                  // 查询所有相匹配的记录 SELECT * FROM nbainfos WHERE age=33;
	// db.Where("age <> ?", 33).Find(&u3)                 // <> SELECT * FROM nbainfos WHERE age <> 33;
	// db.Where("age IN (?)", []int{33, 30}).Find(&u4)    // IN SELECT * FROM nbainfos WHERE age in (33,30);
	// db.Where("name LIKE ?", "%ur%").Find(&u5)          // LIKE SELECT * FROM nbainfos WHERE name LIKE '%ur%'
	// db.Where("age = ? AND id > ?", 33, "31").Find(&u6) // AND SELECT * FROM nbainfos WHERE age=33 AND id>31;
	// // db.Where("updated_at > ?",)                         // Time SELECT * FROM nbainfos WHERE  updated_at > '2000-01-01 00:00:00';
	// db.Where("id BETWEEN ? AND ?", "10", "30").Find(&u8) // BETWEEN SELECT * FROM nbainfos WHERE id BETWEEN 10 AND 30;

	// //where结合struct&map查询
	// db.Where(&NbaInfo{Age: 33, Id: "31"}).Find(&u66)                    // AND SELECT * FROM nbainfos WHERE age=33 AND id=31;
	// db.Where(map[string]interface{}{"Age": 33, "Id": "31"}).Find(&u666) // AND SELECT * FROM nbainfos WHERE age=33 AND id=31;
	// db.Where(&NbaInfo{Age: 33, Id: ""}).Find(&u6666)                    //当通过结构体查询时，gorm只通过非零值字段，如果字段为0，false,''或其他零值，将不构成查询条件,可以通过指针或实现scanner、valuer接口来避免这个问题。 SELECT *FROM nbainfos WHERE AGE = 33;
	// //主键的切片放在where中
	// db.Where([]string{"30", "11"}).Find(&u9) // SELECT * FROM nbainfos WHERE id IN (30,11);

	// fmt.Printf("u1:%v\nu2:%v\nu3:%v\nu4:%v\nu5:%v\nu6:%v\nu8:%v\nu66:%v\nu666:%v\nu6666:%v\nu9:%v\n", u1, u2, u3, u4, u5, u6, u8, u66, u666, u6666, u9)

	//NOT条件
	// var (
	// 	u1 NbaInfo
	// 	u2 []NbaInfo
	// 	u3 []NbaInfo
	// 	u4 []NbaInfo
	// 	u5 NbaInfo
	// )
	// db.Not("name", "curry").First(&u1)          //SELECT * FROM nbainfos WHERE NAME <>"curry" LIMIT 1;
	// db.Not("age", []int{30, 33}).Find(&u2)      //SELECT * FROM nbainfos WHERE age NOT IN (30,33);
	// db.Not([]string{"7", "30", "11"}).Find(&u3) // SELECT * FROM nbainfos WHERE id NOT IN ("7","30","11");
	// db.Not([]string{}).Find(&u4)                //SELECT * FROM nbainfos;
	// db.Not("name = ?", "curry").First(&u5)      //SELECT * FROM nbainfos WHERE NOT(name ="curry");
	// fmt.Printf("u1:%v\nu2:%v\nu3:%v\nu4:%v\nu5:%v\n", u1, u2, u3, u4, u5)

	//OR条件
	// db.Where("name = ?", "curry").Or("name = ?", "klay").Find(&u1)                  // SELECT * FROM nbainfos WHERE name = 'curry' OR name = 'klay';
	// db.Where("name = 'curry'").Or(NbaInfo{Name: "klay"}).Find(&u2)                  // Struct  SELECT * FROM nbainfos WHERE name = 'curry' OR name = 'klay';
	// db.Where("name = 'curry'").Or(map[string]interface{}{"name": "klay"}).Find(&u3) //Map  SELECT * FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2';

	//内联条件
	//额外查询选项
	// var u NbaInfo
	// db.Set("gorm:query_option", "FOR UPDATE").First(&u, "7")
	// fmt.Println(u)
}
