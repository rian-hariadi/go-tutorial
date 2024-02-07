package main

import "fmt"

func main() {
	var num int = 42
	var ptr *int // Deklarasi variabel pointer

	fmt.Println("Nilai awal num", num)
	fmt.Println("Nilai ptr awal", ptr)

	ptr = &num // Inisialisasi pointer dengan alamat memori variabel num
	*ptr = 12

	fmt.Println("Nilai variabel:", num)
	fmt.Println("Alamat memori variabel:", &num)
	fmt.Println("Nilai yang disimpan di alamat memori yang ditunjuk oleh pointer:", *ptr)
	fmt.Println("Alamat memori yang ditunjuk oleh pointer:", ptr)
}
