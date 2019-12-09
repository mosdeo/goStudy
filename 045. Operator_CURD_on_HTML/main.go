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
	t, _ := template.ParseFiles("045. Operator_CURD_on_HTML/" + uri)
	t.Execute(w, nil)
}

func Read(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ShowUI(w, "read.gtpl")
	}

	if r.Method == "POST" {
		ShowUI(w, "read.gtpl")

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

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ShowUI(w, "create.gtpl")
	}

	if r.Method == "POST" {
		ShowUI(w, "create.gtpl")

		err := r.ParseForm()
		if nil == err {
			//獲取欄位上輸入的資料
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

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ShowUI(w, "update.gtpl")
	}

	if r.Method == "POST" {
		ShowUI(w, "update.gtpl")

		err := r.ParseForm()
		if nil == err {
			var id = r.Form.Get("id")
			var new_username = r.Form.Get("new_username")

			db, err := sql.Open("mysql", "root:password@/sreetchat?charset=utf8mb4")
			checkErr(err)

			stmt, err := db.Prepare("Update users SET username=? WHERE id=?")
			checkErr(err)

			result, err := stmt.Exec(new_username, id)
			checkErr(err)
			fmt.Fprintf(w, "Update result=%s\n", result)

			affect, err := result.RowsAffected()
			checkErr(err)

			fmt.Println(affect)
		}
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ShowUI(w, "delete.gtpl")
	}

	if r.Method == "POST" {
		ShowUI(w, "delete.gtpl")

		err := r.ParseForm()
		if nil == err {
			var id = r.Form.Get("id")
			db, err := sql.Open("mysql", "root:password@/sreetchat?charset=utf8mb4")
			checkErr(err)

			stmt, err := db.Prepare("DELETE FROM users WHERE id=" + id)
			checkErr(err)

			result, err := stmt.Exec()
			checkErr(err)
			fmt.Fprintf(w, "Delete result=%s\n", result)

			affect, err := result.RowsAffected()
			checkErr(err)

			fmt.Println(affect)
		}
	}
}

func main() {
	//設定存取的路由
	http.HandleFunc("/", index)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/read", Read)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)

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
