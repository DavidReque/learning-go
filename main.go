package main

import (
	"example/math"
	"fmt"
	//"github.com/rs/zerolog/log"
	//abcd "example/custom"
)

/*type Person struct {
	FirstName, LastName  string
	Age       int
}

type Point struct {
	X, Y float64
}*/

func main() {
	//log.Print(abcd.Value)
	//a := 10

	//var p *int = &a

	//fmt.Println("before", a)
	//fmt.Println("address:", p)

	//*p = 20
	//fmt.Println("after:", a)

	//var p1 = Person{"a", "b", 20}
	//var p2 = Person{"a", "b", 20}

	//fmt.Println(p1 == p2)

	/*p1 := Point{1, 2}
	p2 := p1 // Copy of p1 is assigned to p2

	p2.X = 2

	fmt.Println(p1) // Output: {1 2}
	fmt.Println(p2) // Output: {2 2}*/

	result := math.Add(2, 2)
	fmt.Println(result)
}
