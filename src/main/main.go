package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	_ "github.com/go-sql-driver/mysql"
)

var port = "9090" // 定义端口

func home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./src/views/login.tpl")
	log.Println(t.Execute(w, nil))
}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path", r.URL.Path)
	fmt.Fprintf(w, "register")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path", r.URL.Path)
	fmt.Fprintf(w, "login")
}

func regRouter() {
	http.HandleFunc("/", home)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
}

func main() {
	// 设置访问的路由
	regRouter()
	fmt.Println("server start " + port)
	err := http.ListenAndServe(":"+port, nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
