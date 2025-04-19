package main

import (
	"fmt"
	//"hello/methods"
	//"hello/utils"
)

func main() {
	//fmt.Println(utils.Reverse("Hello, world!"))

	//c := methods.Car{"Tesla", 2023}
	//fmt.Println("is latest:", c.IsLatest())

	//c.UpdateName("Toyota")
	//fmt.Println("car name:", c)

	arr := [2][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}
	
	for i, e := range arr {
		fmt.Println("index:", i, "element:", e)
	}
}
