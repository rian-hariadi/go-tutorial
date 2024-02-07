package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
	int
}

type owner struct {
	name string
}

func something() float32 {
	return 2.3
}

func useGasoline(m int) gasEngine {
	var a gasEngine

	a.mpg = 1
	a.gallons = 23
	a.owner.name = "Dono"
	a.int = m

	fmt.Println(a.owner.name)
	return a

}

func main() {
	var myEngine gasEngine = gasEngine{25, 15, owner{"Rian"}, 12}
	fmt.Print(myEngine)
	test1 := something()

	fmt.Println(test1)
	as := useGasoline(555)
	fmt.Println(as)
}
