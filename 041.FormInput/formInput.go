package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析參數
	fmt.Println(r.Form) //輸出到伺服器端的列印資訊
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //這個 w 是輸出到用戶端的
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //取得請求的方法
	if r.Method == "GET" {
		// 顯示登入介面
		t, _ := template.ParseFiles("041. FormInput/login.gtpl")
		t.Execute(w, nil)
	} else {
		// 若請求登入資料，處理登入的邏輯判斷
		err := r.ParseForm()
		if nil == err {
			fmt.Println("username:", r.Form["username"])
			fmt.Println("password:", r.Form["password"])
		} else {
			fmt.Println(err)
		}
	}
}

func main() {
	http.HandleFunc("/", sayHelloName) //設定存取的路由
	http.HandleFunc("/login", Login)   //設定存取的路由
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
