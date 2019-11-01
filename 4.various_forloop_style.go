package main

import(
    "fmt"
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
}