package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
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

func main() {
	http.HandleFunc("/", sayhelloName)       //ルーティングの設定
	err := http.ListenAndServe(":9090", nil) //ポートの設定
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
