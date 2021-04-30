package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello roaris!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //メソッドの取得
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {
		//フォームに入力されたデータを表示
		r.ParseForm() //これをしないと、フォームのデータに対して操作をすることはできない
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", sayhelloName) //ルーティングの設定
	http.HandleFunc("/login", login)   //ルーティングの設定
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
