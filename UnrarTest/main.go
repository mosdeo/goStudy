package main

import (
	"fmt"

	"github.com/archiver"
	//_ "github.com/mholt/archiver"
)

func main() {
	var pwdDict = []string{"1234", "7777", "88888888"}
	var defaultRar = archiver.NewRar()
	rarPath := "/Users/lky/Downloads/預言.rar"

	for _, pwd := range pwdDict {
		defaultRar.Password = pwd
		err := defaultRar.Unarchive(rarPath, "./")
		if nil != err {
			fmt.Printf("Pwd not %s\n", pwd)
		} else {
			fmt.Printf("Pwd is %s\n", pwd)
			break
		}
	}
}
