package main

import "fmt"

// Immutable's int,float,string,bool,array

func main(){

    x := 10
    y := x
    fmt.Println(x,y)
    y = 20
    fmt.Println(x,y)

    var xx[3]int = [3]int {3,4,5}
    yy := xx
    fmt.Println(xx)
    yy[0] = 12
    fmt.Println(xx,yy)

    a :="Str"
    b := a
    b = "samp"
    fmt.Println(a,b)
}

 