package main

import "fmt"

// for loop

func main()	{

	a := "Sample Foright "

	for i := 0 ; i< len(a) ;i++ {
		fmt.Println(string(a[i]))
	}
	for _,char := range a {
		fmt.Println(string(char))
	}
	loop1()
	loop2()
	loop3()
	loop4()
}

func loop1(){
	for i:=1 ; i<=10 ; i ++{
		fmt.Println(i)
	}
	
}
func loop2(){
	i:=1
	for i<=10{
		fmt.Println(i)
		i++
	}
	
}
func loop3(){
	for i:=1 ; i<5 ; i++{
		for j:=0 ; j<i ; j++{
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func loop4(){
	for i:=5 ; i>=0 ; i--{
		for j:=0 ; j<i ; j++{
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
