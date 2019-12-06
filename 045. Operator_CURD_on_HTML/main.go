package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/Go-SQL-Driver/MySQL"
)

func index(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@/sreetchat?charset=utf8mb4")
	checkErr(err)

	//查詢資料
	rows, err := db.Query("SELECT * FROM sreetchat.users")
	checkErr(err)

	for rows.Next() {
		var id int
		var username string
		var age string
		var sex string
		var timeStamp string
		err = rows.Scan(&id, &username, &age, &sex, &timeStamp)
		checkErr(err)

		t, _ := time.Parse("2006-01-02 15:04:05", timeStamp)
		//fmt.Printf("%v, %v, %v, %v, %v\n", id, username, age, sex, t)
		fmt.Fprintf(w, "%v, %v, %v, %v, %v\n", id, username, age, sex, t)
	}
}

func ShowUI(w http.ResponseWriter, uri string) {
	// 顯示查詢選項介面
	t, _ := template.ParseFiles("045. Operator_CURD_on_page/" + uri)
	t.Execute(w, nil)
}

func query(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ShowUI(w, "query.gtpl")
	}

	if r.Method == "POST" {
		ShowUI(w, "query.gtpl")

		// 若要求查詢，處理查詢的邏輯判斷
		err := r.ParseForm()
		if nil == err {
			//	檢查年齡欄位
			strAgeLower := r.Form.Get("ageLower")
			strAgeUpper := r.Form.Get("ageUpper")
			ageLower, err := strconv.Atoi(strAgeLower)
			ageUpper, err := strconv.Atoi(strAgeUpper)

			if nil != err {
				fmt.Println(err)
				fmt.Fprintln(w, err)
			} else {
				//建立連線
				db, err := sql.Open("mysql", "root:password@/sreetchat?charset=utf8mb4")
				checkErr(err)

				//查詢資料
				rows, err := db.Query("SELECT * FROM sreetchat.users WHERE ?<age AND age<?", ageLower, ageUpper)
				checkErr(err)

				for rows.Next() {
					var id int
					var username string
					var age string
					var sex string
					var timeStamp string
					err = rows.Scan(&id, &username, &age, &sex, &timeStamp)
					checkErr(err)

					t, _ := time.Parse("2006-01-02 15:04:05", timeStamp)
					fmt.Fprintf(w, "%v, %v, %v, %v, %v</br>", id, username, age, sex, t)
				}
			}
		}
	}
}

func add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ShowUI(w, "add.gtpl")
	}

	if r.Method == "POST" {
		ShowUI(w, "add.gtpl")

		err := r.ParseForm()
		if nil == err {
			username := r.Form.Get("username")
			age, _ := strconv.Atoi(r.Form.Get("age"))
			sex := r.Form.Get("sex")
			timestamp := time.Now()
			fmt.Printf("POST username:%v, age:%v, sex:%v, timestamp:%v\n", username, age, sex, timestamp)

			//建立連線
			db, err := sql.Open("mysql", "root:password@/sreetchat?charset=utf8mb4")
			checkErr(err)

			//插入資料
			stmt, err := db.Prepare("INSERT users SET username=?,age=?,sex=?,timestamp=?")
			checkErr(err)

			res, err := stmt.Exec(username, age, sex, timestamp)
			checkErr(err)

			id, err := res.LastInsertId()
			fmt.Println("LastInsertId():", id)
			checkErr(err)
		}
	}
}

func main() {
	//設定存取的路由
	http.HandleFunc("/", index)
	http.HandleFunc("/query", query)
	http.HandleFunc("/add", add)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func checkErr(err error) {
	if nil != err {
		fmt.Errorf("%s", err)
		panic(err)
	}
}
