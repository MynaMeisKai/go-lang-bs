package main

import "fmt"

// Array

func main()	{
	Anum := [5]int{2,4,6,8,10}
	fmt.Println(Anum)

	for i:=0 ; i<len(Anum) ; i++ {
		fmt.Println(Anum[i])
	}

	// Slice 
	slic := Anum[:5]
	fmt.Println(slic)

	// Copy array 
	newArray := make([]int,5,10)
	copy(newArray,slic)
	fmt.Println(newArray)

	// Append
	appArray := append(newArray,100,101,103)
	fmt.Println(appArray)
}
