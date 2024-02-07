package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("source : https://www.youtube.com/watch?v=sTFJtxJXkaY")
	i, j := 42, 270
	// add1, add2 := "", ""
	log.Println("Start")

	p := &i
	fmt.Println("Value of *p or i:", *p)
	fmt.Println("Memory address of p:", p)

	*p = 21
	fmt.Println("Value of i:", i)
	add1 := &i
	address1 := fmt.Sprintf("%p", add1)

	p = &j
	*p = *p / 3
	fmt.Println("value of j:", j)
	fmt.Println("value of i:", i)
	fmt.Println("Memory address of p:", p)
	add2 := &j
	address2 := fmt.Sprintf("%p", add2)

	fmt.Println("Memory address i in string:", address1)
	fmt.Println("Memory address i in string:", address2)

	if address1 != address2 {
		fmt.Println("DIFFERENCE MEMORY ADDRESS")
	} else {
		fmt.Println("SAME!!")
	}

}
