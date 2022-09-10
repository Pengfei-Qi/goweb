package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type User struct {
	Name   string
	Age    int
	Gender string
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//定义模板

	//解析模板
	tep, err := template.ParseFiles("./hello2.tmpl")
	if err != nil {
		fmt.Printf("parse file failed! err: %v", err)
		return
	}
	//数据写入模板

	//a. 使用结构体
	u1 := User{
		Name:   "张三",
		Age:    20,
		Gender: "男",
	}
	//tep.Execute(w, u1)

	//b. 使用map结构
	m1 := map[string]interface{}{
		"Name":   "张四",
		"Age":    23,
		"Gender": "女",
	}
	//tep.Execute(w, m1)

	habbyList := []string{
		"篮球",
		"足球",
		"乒乓球",
	}

	//3. 使用合并内容
	tep.Execute(w, map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"habby": habbyList,
	})
}

func main() {
	http.HandleFunc("/web1", sayHello)
	err := http.ListenAndServe(":9002", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed! err:%v ", err)
		return
	}
}
