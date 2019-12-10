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
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	w := c.Writer

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
	t, _ := template.ParseFiles("046. Operator_CURD_on_HTML_by_Gin/" + uri)
	t.Execute(w, nil)
}

func Read(c *gin.Context) {
	r := c.Request
	w := c.Writer

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

func Create(c *gin.Context) {
	r := c.Request
	w := c.Writer

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

func Update(c *gin.Context) {
	r := c.Request
	w := c.Writer

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

func Delete(c *gin.Context) {
	r := c.Request
	w := c.Writer

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
	router := gin.Default()

	//設定存取的路由
	//找像是原生先進路由，再判斷何種 http method 的
	router.GET("/", index)
	router.GET("/create", func(c *gin.Context) { ShowUI(c.Writer, "create.gtpl") })
	router.GET("/read", func(c *gin.Context) { ShowUI(c.Writer, "read.gtpl") })
	router.GET("/update", func(c *gin.Context) { ShowUI(c.Writer, "update.gtpl") })
	router.GET("/delete", func(c *gin.Context) { ShowUI(c.Writer, "delete.gtpl") })
	router.POST("/create", Create)
	router.POST("/read", Read)
	router.POST("/update", Update)
	router.POST("/delete", Delete)

	err := router.Run(":9090")
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
