package main

import (
	"fmt"
)

// Struct

type Symbol struct{
    xVal int32
    yVal int32
}

func toC(cVal *Symbol,opt int )  {

    var chg int32
    fmt.Println("Enter your'e Number : ")
    fmt.Scanln(&chg)

    if opt == 1 {
        cVal.xVal = chg 
    } else if opt == 2 {
        cVal.yVal = chg
    }else{
        fmt.Println("--- Canceled ---")
    }
}

func main(){
    var  xx int32 
    var yy int32
    var opt int
    stor := make(map[string]map[string]int32)

    fmt.Println("Enter any Num1 xx :  ") 
    fmt.Scanln(&xx)
    fmt.Println("Enter any Num2 yy :  ")  
    fmt.Scanln(&yy)
    stor["Before"] = map[string]int32{
                "xx" : xx,
                "yy" : yy,
    }
    s1 := &Symbol{xx,yy}
    fmt.Printf("Num1 xx : %d \nNum2 yy : %d\n",xx,yy)

    fmt.Println("Do you want to change Num's [Yes]- 1 [No] - 0 : ")
    fmt.Scanln(&opt)

    if opt == 1 {
        fmt.Println("You choose Yes !! \nEnter option to change [Num1]- 1 [Num2] - 2 : ") 
        fmt.Scanln(&opt)
        if opt == 1 {
            toC(s1,opt)
        }else if  opt == 2 {
            toC(s1,opt)
        }else {
            fmt.Println(" --- Canceled --- you opt non ")
        }
    }else{
        fmt.Println(" --- Canceled --- you opt 0 or any")
    }
    
    stor["After"] = map[string]int32{
        "xx" : s1.xVal,
        "yy" : s1.yVal,
        }
        
    fmt.Println("Before val : ",stor["Before"])
    fmt.Println("After val : ",stor["After"])
}




// type Symbol struct {

//     x int32
//     y float64 
     
// }

// func main() {
//     var s1 Symbol = Symbol{9,0.34}
//     fmt.Println(s1.x)
//     fmt.Println(s1.y)
// }