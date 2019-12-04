package main

import (
	"database/sql"
	"fmt"

	_ "github.com/Go-SQL-Driver/MySQL"
)

func main() {
	db, err := sql.Open("mysql", "astaxie:astaxie@/test?charset=utf8")
	checkErr(err)

	//插入資料
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研發部門", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	//更新資料
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astiXXXXXupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查詢資料
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)

		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//刪除資料
	//還沒有打
}

func checkErr(err error) {
	if nil != err {
		panic(err)
	}
}
