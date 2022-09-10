package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func hello(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	tem, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("parse file failed, err:%v", err)
		return
	}
	//渲染模板
	msg := "小王子"
	tem.Execute(w, msg)
}
func home(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	tem, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Printf("parse file failed, err:%v", err)
		return
	}
	//渲染模板
	msg := "小王子"
	tem.Execute(w, msg)
}

func hello2(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	tem, err := template.ParseFiles("./templates/base.tmpl", "./templates/index2.tmpl")
	if err != nil {
		fmt.Printf("parse file failed, err:%v", err)
		return
	}
	//渲染模板
	msg := "张三"
	tem.ExecuteTemplate(w, "index2.tmpl", msg)
}

func home2(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	tem, err := template.ParseFiles("./templates/base.tmpl", "./templates/home2.tmpl")
	if err != nil {
		fmt.Printf("parse file failed, err:%v", err)
		return
	}
	//渲染模板
	msg := "李四"
	tem.ExecuteTemplate(w, "home2.tmpl", msg)
}

func main() {
	http.HandleFunc("/index", hello)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", hello2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("parse file failed, err:%v", err)
		return
	}
}
