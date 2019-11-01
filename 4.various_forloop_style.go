package main

import(
    "fmt"
    //"reflect"
    //"strconv"
)

func main(){

    // Normal style
    for i := 0;i<10;i++{
        fmt.Print(i)
    }
    fmt.Println();

    // While style
    var i int
    for i<10{
        fmt.Print(i)
        i++
    }
    fmt.Println();

    // While style
    var j int
    for {
        fmt.Print(j)
        j++
        if j>=10{
            break
        }
    }
    fmt.Println();

    //Range style
    var arrString string = "This is a string"    
    for _, c := range arrString[0:]{
        // fmt.Println(reflect.TypeOf(c)) == int32
        // c 的型別是 rune，相當於 C++ 中的 char，直接 print 會當作 int32 印出數字
        // string(c) 才會看到對應的字元
        fmt.Print(string(c))
    }
    fmt.Println();
}