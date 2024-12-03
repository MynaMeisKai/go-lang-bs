package main

import "fmt"

// Array

func main()	{
	Anum := [5]int{1,2,3,4,5}
	// fmt.Println(Anum)

	for i:=0 ; i<len(Anum) ; i++ {
		fmt.Println(Anum[i])
	}
}
