package main

import (
	//"example/math"
	"fmt"
	"sync"
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

func speak(args string, ch chan <- string) {
	ch <- args // solo envia mensaje al canal
}

func work(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("working...")
}

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

	/*result := math.Add(2, 2)
	fmt.Println(result)*/
	
	/*one := make(chan string)

	two := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		one <- "one"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		two <- "two"
	}() 

	for x := 0; x < 10; x++ {
		select {
		case result := <-one:
			fmt.Println("Received from one:", result)
		case result := <-two:
			fmt.Println("Received from two:", result)
		default:
			fmt.Println("Default...")
			time.Sleep(200 * time.Millisecond)
		}
	}

	close(one)
	close(two)*/

	var wg sync.WaitGroup

	wg.Add(2)
	go work(&wg)
	go work(&wg)

	wg.Wait()
}
