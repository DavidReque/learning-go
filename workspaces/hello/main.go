package main

import (
	"errors"
	"fmt"
	"hello/errorsgo"
)

//"hello/methods"
//"hello/interfaces"
//"hello/utils"

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

	//a := [5]int{20, 15, 5, 30, 25}

	//s := a[1:4]

	// Output: Array: [20 15 5 30 25], Length: 5, Capacity: 5
	//fmt.Printf("Array: %v, Length: %d, Capacity: %d\n", a, len(a), cap(a))

	// Output: Slice [15 5 30], Length: 3, Capacity: 4
	//fmt.Printf("Slice: %v, Length: %d, Capacity: %d", s, len(s), cap(s))

	/*var m = map[string]methods.User{
		"1": {"John"},
		"2": {"Doe"},
	}

	//delete(m, "1")

	fmt.Println(m["1"])

	for k, v := range m {
		fmt.Println("Key:", k, "Value:", v)

	}*/

/*
	m := interfaces.Mobile{"Apple"}
	l := interfaces.Laptop{"Intel"}
	t := interfaces.Toaster{2}
	k := interfaces.Kettle{"50%"}

	s := interfaces.Socket{}
	
	s.Plug(m, 10)
	s.Plug(l, 50)
	s.Plug(t, 30)
	s.Plug(k, 25)
*/

	result, err := errorsgo.Divide(4, 0)

	if err != nil {
		var divErr errorsgo.DivisionError

		switch {
			case errors.As(err, &divErr):
				fmt.Println("DivisionError:", divErr.Code, divErr.Msg)
			default: 
				fmt.Println("Error:", err.Error())
		}

		return
	}

	fmt.Println("Result:", result)
}
