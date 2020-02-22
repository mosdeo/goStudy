package main

import (
	"fmt"
	"time"
)

func main() {

	var create_time string = "2020-01-08T13:41:43.399631+08:00"

	DB_TIMESTAMPTZ_FORMAT := "2006-01-02T15:04:05.000000-07:00"
	t_create, err_c := time.Parse(DB_TIMESTAMPTZ_FORMAT, create_time)
	fmt.Println(err_c)
	fmt.Println(t_create)
}
