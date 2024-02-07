package main

import "fmt"

func main() {

	// Membuat map untuk menyimpan informasi buah
	buah1 := map[string]string{
		"nama":  "Apel",
		"warna": "Merah",
	}

	buah2 := map[string]string{
		"nama":  "Jeruk",
		"warna": "Oranye",
	}

	buah3 := map[string]string{
		"nama":  "Pisang",
		"warna": "Kuning",
	}

	// Membuat slice dari map buah
	// var daftarBuah []map[string]string

	var daftarBuah = []map[string]string{
		{
			"nama":  "Durian",
			"warna": "Coklat",
		},
	}

	daftarBuah = append(daftarBuah, buah1, buah2, buah3)

	// Menampilkan informasi buah dari slice
	for _, buah := range daftarBuah {
		fmt.Printf("Nama: %s, Warna: %s\n", buah["nama"], buah["warna"])
	}
}
