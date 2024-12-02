package main

import "fmt"

// formate print function

func main()	{
	
	a:= "Amy Stanly"
	b:= 221
	c:= 4.214653

	fmt.Printf("%T\n",a) // Type of Data is Stored
	fmt.Printf("%d %T \n",b,b) // Value and the  Type
	fmt.Printf("%.3f %T \n",c,c) // Value with given no of floating points  and the  Type
	fmt.Printf("%b \n",8) // 32 16 8 4 2 1 binary value 
	fmt.Printf("%x",42) // hex value 
}
