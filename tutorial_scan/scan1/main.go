package main

import "fmt"

func main() {
	var input string
	var bil1 int
	var bil2 int

	fmt.Print("Masukkan nama: ")
	fmt.Scan(&input)

	fmt.Print("Masukkan bilangan 1: ")
	fmt.Scan(&bil1)

	fmt.Print("Masukkan bilangan 2: ")
	fmt.Scan(&bil2)

	fmt.Println("Anda memasukkan nama:", input)
	fmt.Println("Jumlah bilangan:", bil1+bil2)
}
