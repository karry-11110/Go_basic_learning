package main

import "fmt"

type student struct {
	id, age int
	name    string
	score   float32
}

var studNum = 0 //全局变量计数
//依据结构体写一个构造方法
func newStu(id int, name string, age int, score float32) *student { //这里结构体构造方法才是构造对象的正确打开方式
	return &student{
		id:    id,
		name:  name,
		age:   age,
		score: score,
	}
}

func add(a *[]student) { //考虑到需要进行编辑和删除，id变量一定要有顺序,//这种把students传进来的做法对修改切片没作用，引用型不应该有作用吗？？？？？？
	//涉及到变量作用域的知识点 因此要把students设置为全局变量,或者把引用型变量的指针传进来，只有值传递
	fmt.Println("请输入你要添加几个学生的信息：")
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("输入无效，err:", err)
	}
	var id = studNum
	for i := 0; i < num; i++ {
		id = studNum
		studNum += 1
		var name string
		var age int
		var score float32

		fmt.Print("请输入学生姓名:")
		_, err = fmt.Scan(&name) //注意这里就不能用简短：=符号了不能重复用
		if err != nil {
			fmt.Println("输入无效，err:", err)
		}
		fmt.Print("请输入学生年龄:")
		_, err = fmt.Scan(&age)
		if err != nil {
			fmt.Println("输入无效，err:", err)
		}
		fmt.Print("请输入学生分数:")
		_, err = fmt.Scan(&score)
		if err != nil {
			fmt.Println("输入无效，err:", err)
		}
		newStu(id, name, age, score)
		(*a) = append((*a), *newStu(id, name, age, score))
	}
}

func edit(a *[]student) {
	var i int
	fmt.Println("请输入学生id：")
	_, err := fmt.Scan(&i)
	if err != nil {
		fmt.Println("输入错误！！ERROR:", err)
	}
	for _, v := range *a {
		if v.id == i {
			fmt.Print("请输入学生姓名：")
			_, err := fmt.Scan(&(*a)[i].name)
			if err != nil {
				fmt.Println("输入错误！！ERROR:", err)
			}
			fmt.Print("请输入学生年龄：")
			_, err = fmt.Scan(&(*a)[i].age)
			if err != nil {
				fmt.Println("输入错误！！ERROR:", err)
			}
			fmt.Print("请输入学生成绩：")
			_, err = fmt.Scan(&(*a)[i].score)
			if err != nil {
				fmt.Println("输入错误！！ERROR:", err)
			}
		}
	}

}

func delete(a *[]student) { // 输入学生id进行删除
	var i int
	fmt.Println("请输入学生id：")
	_, err := fmt.Scan(&i)
	if err != nil {
		fmt.Println("输入错误！！ERROR:", err)
	}
	for _, v := range *a {
		if v.id == i {
			(*a) = append((*a)[:i], (*a)[i+1:]...)
		}
	}
}

func main() {
	var students []student

	fmt.Println("***欢迎来到学生管理系统！！！***")

	for {
		fmt.Println("请选择你需要执行的功能：1.查看学生信息 2.添加学生信息 3.编辑学生信息 4.删除学生信息 其他：退出系统")
		var button int
		_, err := fmt.Scan(&button) //输入函数这里一定要定义清楚，否则会自动执行default
		if err != nil {
			fmt.Println("输入有误！")
		}
		switch button {
		case 1:
			fmt.Println(students) //展示学生列表
		case 2:
			add(&students) //添加学生信息
		case 3:
			edit(&students) //编辑学生信息
		case 4:
			delete(&students) //删除学生信息
		// case 5:
		// 	goto breakTag
		default:
			goto breakTag //这是必须要查询的
			// fmt.Println("输入有误")
		}
	}
breakTag:
	fmt.Println("查询结束！！！")

}
