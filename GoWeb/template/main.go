package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//1.先定义好一个hello.tmpl的模板文件
func sayHello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./hello.tmpl") //2.解析模板  这里注意如果用go build，可执行文件的位置一定要放对
	if err != nil {
		fmt.Printf("template parsefiles failed,err:%v\n", err)
		return
	}
	// t.Execute(w, "wangkun") //3.模板渲染
	name := "wangkun"
	t.Execute(w, name) //3.模板渲染
}

//当我们传入一个结构体对象时，我们可以根据.来访问结构体的对应字段。
//同理，当我们传入的变量是map时，也可以在模板文件中通过.根据key来取值
type UserInfo struct { //注意这里就要大小，因为在模板中要用到了
	Name   string
	Gender string
	Age    int
}

func sayHello01(w http.ResponseWriter, r *http.Request) {
	t01, err := template.ParseFiles("./hello01.tmpl")
	if err != nil {
		fmt.Printf("template parsefiles failed,err:%v\n", err)
		return
	}
	user := UserInfo{
		Name:   "王坤",
		Gender: "男",
		Age:    23,
	}
	maper := map[string]interface{}{
		"name":   "wangkun",
		"gender": "男",
	}
	lister := []string{
		"篮球",
		"游泳",
	}
	t01.Execute(w, map[string]interface{}{
		"u1": user,
		"m1": maper,
		"l1": lister,
	})
}

//在模板中使用自定义函数
func diySayHello(w http.ResponseWriter, r *http.Request) {
	//定义函数，要么有一个返回值，要么有两个返回值，第二个返回值必须是error类型
	kua := func(name string) (string, error) {
		return name + "年轻又帅气", nil
	}
	//定义模板hellodiy
	tdiy := template.New("hellodiy.tmpl") //创建一个名字是hellodiy.tmpl的模板对象
	tdiy.Funcs(template.FuncMap{
		"kua": kua,
	}) //在解析模板之前，一定要告诉模板引擎，我现在多了一个自定义的函数kua
	//修改默认标识符template.New("test").Delims("{[", "]}").ParseFiles("./t.tmpl")
	_, err := tdiy.ParseFiles("./hellodiy.tmpl")
	if err != nil {
		fmt.Printf("template parsefiles failed,err:%v\n", err)
		return
	}
	name := "wk" // 渲染模板
	tdiy.Execute(w, name)
}

//模板嵌套
func qiantaoDemo(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./t.tmpl", "./ul.tmpl") //解析模板，在解析模板时，被嵌套的模板一定要在后面解析
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	user := UserInfo{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	tmpl.Execute(w, user)
}

//模板继承
func index(w http.ResponseWriter, r *http.Request) {
	//定义模板(模板继承的方式)
	tmplindex, err := template.ParseFiles("./base.tmpl", "./index.tmpl") //解析模板,由根模板包含与被包含的关系
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	tmplindex.ExecuteTemplate(w, "index.tmpl", "wk") //渲染模板,渲染多个模板，所以需要制定

}
func base(w http.ResponseWriter, r *http.Request) {
	//定义模板
	tmplbase, err := template.ParseFiles("./base.tmpl") //解析模板
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	tmplbase.Execute(w, "wk") //渲染模板

}

func main() {
	http.HandleFunc("/", sayHello) //添加处理器，只要访问前者，执行后者函数
	http.HandleFunc("/01", sayHello01)
	http.HandleFunc("/diy", diySayHello)
	http.HandleFunc("/qiantao", qiantaoDemo)
	http.HandleFunc("/index", index)
	http.HandleFunc("/base", base)
	err := http.ListenAndServe(":9090", nil) //使用指定的监听地址和处理器启动一个http服务端
	if err != nil {
		fmt.Printf("http serve failed ,err:%v\n", err)
		return
	}
}
