package main

import (
	"fmt"
	"time"
)

const arraySize = 1000000

type Data struct {
	Array [arraySize]int
}

func modifyValue(data Data) {
	for i := 0; i < arraySize; i++ {
		data.Array[i] = data.Array[i] * 2
	}
}

func modifyPointer(data *Data) {
	for i := 0; i < arraySize; i++ {
		data.Array[i] = data.Array[i] * 2
	}
}

func main() {
	// Membuat variabel struct Data
	data := Data{
		Array: [arraySize]int{},
	}

	// Mengukur waktu untuk mengubah nilai dengan salinan
	startTime := time.Now()
	modifyValue(data)
	durationValue := time.Since(startTime)

	// Mengukur waktu untuk mengubah nilai dengan pointer
	startTime = time.Now()
	modifyPointer(&data)
	durationPointer := time.Since(startTime)

	fmt.Printf("Waktu untuk mengubah nilai dengan salinan: %v\n", durationValue)
	fmt.Printf("Waktu untuk mengubah nilai dengan pointer: %v\n", durationPointer)
}
