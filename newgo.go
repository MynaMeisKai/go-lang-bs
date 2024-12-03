package main

import "fmt"

// Map

func main()	{

	var mp map[string]int = map[string]int {
		"Maddy" : 14,
		"Amy":13,
		"Rain" : 18,
	}
	fmt.Println(mp)

	val,x := mp["Rn"]

	// // x return's if key is in map true / false
	// // val is value of the key ,if the is no key/false it give's 0

	fmt.Println(val,x)

	delete(mp,"Maddy")
	fmt.Println(mp)

	ms := make(map[string]int)
	ms["Sam"] = 34
	fmt.Println(ms)

	// map in map

	Superhero:= map[string]map[string]string{
		"Super Man" : {
			"RealName" : "Clark Kent",
			"City" : "Metropolis City",
		},
		"Bat Man" :  {
			"RealName" : "Bruce Wayne",
			"City" : "Gowtham City",
		},
	}

	fmt.Println( "Superhero  =>>>  ",Superhero)

	value,ok := Superhero["Bat Man"]
	fmt.Println(value["City"],ok)
}
