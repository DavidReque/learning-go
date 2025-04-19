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

	/*arr := [2][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}

	for i, e := range arr {
		fmt.Println("index:", i, "element:", e)
	}*/

	a := [5]int{20, 15, 5, 30, 25}

	s := a[1:4]

	// Output: Array: [20 15 5 30 25], Length: 5, Capacity: 5
	fmt.Printf("Array: %v, Length: %d, Capacity: %d\n", a, len(a), cap(a))

	// Output: Slice [15 5 30], Length: 3, Capacity: 4
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d", s, len(s), cap(s))
}
