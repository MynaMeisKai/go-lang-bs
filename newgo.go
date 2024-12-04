package main

import "fmt"

// func to func

func samp1(newFunc func(int)int){
	fmt.Println(newFunc(9))
	
}
//Defer

func add(x,y int)(z1,z2 int){
	defer fmt.Println("Good!!")
	z1 = x+y
	z2 = x* -y
	fmt.Println("Before return")
	return

}

// functoin closure :

func dos() func()int{
	count := 0
	fmt.Println("Do's called !!")
	return func() int {
		val := count +1 
		count = count+1
		return val
	}

}

func main(){

	test := func (x int)int {
		return x * -2
	}
	samp1(test) //func --> func

	adds,prods := add(2,6)
	fmt.Printf("Sum : %d \nProduct : %d ",adds,prods) //defer 

	funs := dos()
	for i:=0 ; i<=10 ; i++{
		fmt.Println(funs())
	}
}

