package main

import "fmt"

// User input 

func main()	{
	
		var a string
		fmt.Print("Enter Name : ")
		fmt.Scanln(&a) 
		fmt.Println("Hello " + a)

		var b,c int 
		fmt.Print("Enter b : ")
		fmt.Scanln(&b)
		fmt.Print("Enter c : ")
		fmt.Scanln(&c)
		fmt.Println(b+c)


}
