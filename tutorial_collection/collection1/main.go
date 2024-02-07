package main

import "fmt"

func main() {
	// Array
	var angka [5]int = [5]int{1, 2, 3, 4, 5}

	// Slice
	var numbers = []int{1, 2, 3, 4, 5}

	// Map
	var person = map[string]string{
		"nama":      "John",
		"usia":      "30",
		"pekerjaan": "Developer",
	}

	// Slice dari Map
	var people = []map[string]string{
		{"nama": "Alice", "usia": "25"},
		{"nama": "Bob", "usia": "30"},
		{"nama": "Charlie", "usia": "35"},
	}

	// Channel
	var ch = make(chan int)

	// Struct
	type Person struct {
		Nama      string
		Usia      int
		Pekerjaan string
	}

	var orang Person

	fmt.Println("Array:", angka)
	fmt.Println("Slice:", numbers)
	fmt.Println("Map:", person)
	fmt.Println("Slice dari Map:", people)
	fmt.Println("Channel:", ch)
	fmt.Println("Struct:", orang)

	fmt.Println("=========================================")
	testChannel()

}

func testChannel() {
	pesan := make(chan string)

	go func() {
		pesan <- "Hello, Golang!"
	}()

	pesanDiterima := <-pesan
	fmt.Println(pesanDiterima)

}
