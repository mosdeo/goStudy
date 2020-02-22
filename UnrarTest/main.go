package main

import (
	"fmt"

	"github.com/archiver"
	//_ "github.com/mholt/archiver"
)

func main() {
	var defaultRar = archiver.NewRar()
	defaultRar.Password = "1234"
	// rarPath := "/Users/lky/Downloads/Unknown.rar"
	rarPath := "/Users/lky/Downloads/password1234.rar"
	err := defaultRar.Unarchive(rarPath, "./")
	fmt.Println(err)
}
