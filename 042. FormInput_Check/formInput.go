package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
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
		t, _ := template.ParseFiles("042. FormInput_Check/login.gtpl")
		t.Execute(w, nil)
	} else {
		// 若要求登入，處理登入的邏輯判斷
		err := r.ParseForm()
		if nil == err {
			//開始各項欄位檢查
			if "" == r.Form.Get("username") {
				msg := "username欄位不存在"
				fmt.Println(msg)
				fmt.Fprintln(w, msg)
			}
			if 0 == len(r.Form["username"][0]) {
				msg := "沒有輸入 username"
				fmt.Println(msg)
				fmt.Fprintln(w, msg)
			}

			//	檢查年齡欄位
			getAge, err := strconv.Atoi(r.Form.Get("age"))
			if nil != err {
				fmt.Println(err)
				fmt.Fprintln(w, err)
			} else {
				msg := fmt.Sprintf("年齡是:%d\n", getAge)
				fmt.Println(msg)
				fmt.Fprintln(w, msg)
			}

			// 檢查是否為中文字
			if m, _ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]+$", r.Form.Get("realname")); !m {
				msg := fmt.Sprintf("中文實名欄位 %v 含有非中文\n", r.Form.Get("realname"))
				fmt.Println(msg)
				fmt.Fprintln(w, msg)
			}

			// 檢查是否為英文

			// 檢查電子郵件地址

			// 檢查手機號碼

			// 多選項檢查器
			var MultiSelectionCheck = func(selectionName string, selectionTitleText string, legalList map[string]struct{}) {
				// 先檢查欄位是否存在？
				_, ok := legalList[r.Form.Get(selectionName)]
				if ok {
					//如果選項存在，拋出已選的選項
					msg := fmt.Sprintf("%s選了 %v\n", selectionTitleText, r.Form[selectionName])
					fmt.Println(msg)
					fmt.Fprintln(w, msg)
				} else {
					msg := fmt.Sprintf("選項 %v 不在選單中\n", r.Form.Get(selectionName))
					fmt.Println(msg)
					fmt.Fprintln(w, msg)
				}
			}

			// 檢查下拉式功能選單
			dropDownList := map[string]struct{}{
				"apple":  struct{}{},
				"pear":   struct{}{},
				"banana": struct{}{},
			}
			MultiSelectionCheck("fruit", "下拉式功能選單-水果", dropDownList)
			// _, ok := list[r.Form.Get("fruit")]
			// if ok {
			// 	msg := fmt.Sprintf("下拉式選單選了 %v\n", r.Form.Get("fruit"))
			// 	fmt.Println(msg)
			// 	fmt.Fprintln(w, msg)
			// } else {
			// 	msg := fmt.Sprintf("選項 %v 不在選單中\n", r.Form.Get("fruit"))
			// 	fmt.Println(msg)
			// 	fmt.Fprintln(w, msg)
			// }

			// 必須選項按鈕
			radioBoxList := map[string]struct{}{
				"M": struct{}{},
				"F": struct{}{},
			}
			MultiSelectionCheck("gender", "必須選項按鈕-性別", radioBoxList)
			// _, ok := list[r.Form.Get("gender")]
			// if ok {
			// 	msg := fmt.Sprintf("選項按鈕選了 %v\n", r.Form.Get("gender"))
			// 	fmt.Println(msg)
			// 	fmt.Fprintln(w, msg)
			// } else {
			// 	msg := fmt.Sprintf("選項 %v 不在選單中\n", r.Form.Get("gender"))
			// 	fmt.Println(msg)
			// 	fmt.Fprintln(w, msg)
			// }

			// 核取按鈕
			checkBoxList := map[string]struct{}{
				"interest0": struct{}{},
				"interest1": struct{}{},
				"interest2": struct{}{},
			}
			fmt.Fprintf(w, "r.Form[\"interest\"]=%s", r.Form["interest"])
			MultiSelectionCheck("interest", "核取按鈕-興趣", checkBoxList)
			// _, ok := list[r.Form.Get("interest")]
			// if ok {
			// 	msg := fmt.Sprintf("核取按鈕選了 %v\n", r.Form.Get("interest"))
			// 	fmt.Println(msg)
			// 	fmt.Fprintln(w, msg)
			// } else {
			// 	msg := fmt.Sprintf("選項 %v 不在選單中\n", r.Form.Get("interest"))
			// 	fmt.Println(msg)
			// 	fmt.Fprintln(w, msg)
			// }

			// 日期和時間

			// 身分證字號

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
