package main

import "fmt"

func sam(){
	fmt.Println("W-W-W")
}

func add(a,b int) int{ // 1 st int is datatyp 2nd int is return value
	return a+b
}

func sum1(x,y int) (z1,z2 int){ // z1,z2 are labels
	return x+y , x*y
}

func main(){

	sam() // simple fun call
	add1 := add       // passing Arguments in func
	fmt.Println("Sum : ", add1(1,2))
	
	// multiple output func 
	sums,pro:=sum1(7,2)
	fmt.Println("Ans : ",sums , pro)

	// Anonymous Function 
	samp1 := func (){
		fmt.Println("samp-1")
	}
	samp1()

	samp2 := func(x int)int {
		return x*-2
	}(2)
	fmt.Println(samp2)


}
