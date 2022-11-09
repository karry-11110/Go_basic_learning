package main

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

// gorm.Model 定义
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// 将 `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`字段注入到`User`模型中
type User struct {
	gorm.Model
	Name string
}

// 不使用gorm.Model，自行定义模型
type User1 struct {
	ID   int
	Name string
}

type User2 struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

//内置约束
// 主键：GORM 默认会使用名为ID的字段作为表的主键，或者使用tag:`gorm:"primary_key"`
// 表名：1.表名默认就是结构体名称的复数 2.禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`db.SingularTable(true)
//      3.通过Table()指定表名，使用User结构体创建名为`deleted_users`的表	 db.Table("deleted_users").CreateTable(&User{})
// 列名：1.结构体tag指定列名 `gorm:"column:beast_id"`  2.列名由字段名称进行下划线分割来生成
// 时间戳跟踪：
