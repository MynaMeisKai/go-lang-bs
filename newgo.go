package main

import (
	"fmt"
	// "reflect"
)

// Pointer

/*
 * - Astrick or Derefrenc
 & - GetPointer or Refrence
*/


func toChange(sam *string) {
    *sam = "Bad One"
}

func main() {

    // x:= 7
    // y:= &x
    // fmt.Println(x,y)
    // *y = 12
    // fmt.Println(x)
    
    // or_str := "Good One"
    // fmt.Println(or_str)

    // toChange(&or_str)
    // fmt.Println(or_str)
 
    to_run := "Duty"
    var pointer *string = &to_run
    fmt.Println(*pointer) // Duty
    fmt.Println(pointer) // O/p Address to to_run
    fmt.Println(&pointer) // O/p address of pointer
}

// two way to find TypeOf var or obj
    // a := 7444
    // fmt.Printf("%T",a)
    // fmt.Printf("%v \n",reflect.TypeOf(a))