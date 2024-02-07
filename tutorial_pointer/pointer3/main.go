package main

import (
	"fmt"
	"time"
)

func main() {
	WithPointer()
	fmt.Println("---------------------------------------")
	WithoutPointer()
}

func WithPointer() {
	fmt.Println("With Pointer")
	var originalValue int = 42
	var pointerToValue *int = &originalValue

	startTime := time.Now()

	// Menggandakan nilai melalui pointer
	for i := 1; i < 1000000; i++ {
		*pointerToValue += 2
	}

	elapsedTime := time.Since(startTime)

	fmt.Printf("Waktu yang diperlukan untuk menggunakan pointer: %v\n", elapsedTime)
	fmt.Printf("Nilai akhir melalui pointer: %d\n", originalValue)

}

func WithoutPointer() {
	fmt.Println("Without Pointer")
	var originalValue int = 42

	startTime := time.Now()

	// Menggandakan nilai tanpa menggunakan pointer
	for i := 1; i < 1000000; i++ {
		originalValue += 2
	}

	elapsedTime := time.Since(startTime)

	fmt.Printf("Waktu yang diperlukan tanpa menggunakan pointer: %v\n", elapsedTime)
	fmt.Printf("Nilai akhir tanpa menggunakan pointer: %d\n", originalValue)
}
