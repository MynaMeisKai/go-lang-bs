package main

import "fmt"

func main(){
	var a  int32 = 10
	b := 20
	fmt.Println("This is address : ",&a,&b)
	fmt.Println("This is Value : ",a,b)
}