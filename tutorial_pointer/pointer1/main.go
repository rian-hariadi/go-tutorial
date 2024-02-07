package main

import "fmt"

func main() {
	var ptr *int
	var num = 42

	ptr = &num
	*ptr = 12

	fmt.Println(ptr)
	fmt.Println(num)

}
