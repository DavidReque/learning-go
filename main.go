package main

import (
	"fmt"
	//"github.com/rs/zerolog/log"
	//abcd "example/custom"
)

func main() {
	//log.Print(abcd.Value)
	a := 10

	var p *int = &a

	fmt.Println("before", a)
	fmt.Println("address:", p)

	*p = 20
	fmt.Println("after:", a)
}
