package main

import "fmt"

func main()	{

	x := 10
	fmt.Println("Old Value :>>> ",x)
	changeValue(&x)
	fmt.Println("New Value :>>> ",x)
}

func changeValue(x *int)	{
	*x = 20
}
