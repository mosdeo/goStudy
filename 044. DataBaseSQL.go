package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"reflect"
	"time"

	_ "github.com/Go-SQL-Driver/MySQL"
)

func TrueRandom(max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max)
}

func main() {
	//db, err := sql.Open("mysql", "user:password@/dbname")
	db, err := sql.Open("mysql", "root:password@/sreetchat?charset=utf8mb4")
	checkErr(err)

	//插入資料
	stmt, err := db.Prepare("INSERT users SET username=?,age=?,sex=?,timestamp=?")
	checkErr(err)

	// 執行 query 或 prepared
	res, err := stmt.Exec("LKY", TrueRandom(100), "M", time.Now())
	checkErr(err)

	id, err := res.LastInsertId()
	fmt.Println("LastInsertId():", id)
	checkErr(err)

	//更新資料
	stmt, err = db.Prepare("update users set username=? where id=?")
	checkErr(err)

	res, err = stmt.Exec("LKY_New", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查詢資料
	rows, err := db.Query("SELECT * FROM sreetchat.users")
	checkErr(err)

	for rows.Next() {
		var id int
		var username string
		var age string
		var sex string
		var iTimeStamp interface{}
		err = rows.Scan(&id, &username, &age, &sex, &iTimeStamp)
		checkErr(err)

		// 如果時間戳記不為空
		if nil != iTimeStamp {
			// 判斷時間戳記類型
			switch v := iTimeStamp.(type) {
			case []uint8:
				t, _ := time.Parse("2006-01-02 15:04:05", string(v))
				fmt.Printf("%v, %v, %v, %v, %v, %v\n", id, username, age, sex, t, reflect.TypeOf(v))
			case time.Time:
				fmt.Printf("%v, %v, %v, %v, %v, %v\n", id, username, age, sex, v, reflect.TypeOf(v))
			case string:
				//t, err := time.Parse("2006-01-02 15:04:05", v)
				t, err := time.Parse(time.Now().String(), v)
				if nil != err {
					fmt.Printf("%v, %v, %v, %v, %v, %v\n", id, username, age, sex, t, reflect.TypeOf(t))
				} else {
					fmt.Println(err)
				}
			}
		} else {
			fmt.Printf("%v, %v, %v, %v, %v\n", id, username, age, sex, "timestamp is null")
		}
	}

	//刪除資料
	//還沒有打
}

func checkErr(err error) {
	if nil != err {
		fmt.Errorf("%s", err)
		panic(err)
	}
}
