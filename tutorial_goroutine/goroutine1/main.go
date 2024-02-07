package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int = 0
var mu sync.Mutex // Mutex untuk sinkronisasi akses ke counter

func main() {
	// Membuat tiga goroutine yang berjalan secara konkuren
	go printNumbers("Go ++++++++++++++++++++++  1")
	go printNumbers("Go ----------------------  2")
	go printNumbers("Go xxxxxxxxxxxxxxxxxxxxxx  3")

	// Menunggu sebentar agar goroutine selesai
	time.Sleep(50 * time.Second)

	fmt.Println("Program selesai.")
}

func printNumbers(prefix string) {
	for i := 1; i <= 50; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%s: %d\n", prefix, i)
		printCounter(prefix)
	}
}

func printCounter(prefix string) {
	// For testing, try to comment lock
	mu.Lock()

	for i := 1; i <= 4; i++ {
		// Menggunakan Lock dan Unlock untuk melindungi akses ke counter

		counter++
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("%s  ===>(Counter: %d) : %d   \n", prefix, counter, i)

	}
	mu.Unlock()
}
