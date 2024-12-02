package main

import "fmt"

func main(){
	var a  int32 = 10
	b := 20.222
	fmt.Println("This is address : ",&a,&b)
	fmt.Println("This is Value : ",a,b)
	stringMain()
}

func stringMain(){
	var Name string = "Amy Stanly"
	fmt.Println("Name : ",Name)
	fmt.Println("Length of the string : ",len(Name))
}