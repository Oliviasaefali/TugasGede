package main

import (
	"fmt"
	"sort"
	"strings"
)

type Crypto struct {
	Nama  string
	Harga float64
}

var daftarCrypto []Crypto

func main() {
	for {
		fmt.Println("\n=== Menu Aplikasi Crypto ===")
		fmt.Println("1. Tambah Crypto")
		fmt.Println("2. Lihat Daftar Crypto")
		fmt.Println("3. Cari Crypto (berdasarkan Nama)")
		fmt.Println("4. Urutkan Crypto (berdasarkan Harga)")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu (1-5): ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahCrypto()
		case 2:
			lihatDaftar()
		case 3:
			cariCrypto()
		case 4:
			urutkanCrypto()
		case 5:
			fmt.Println("Keluar dari aplikasi. Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahCrypto() {
	var nama string
	var harga float64

	fmt.Print("Masukkan Nama Crypto: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan Harga Crypto (USD): ")
	fmt.Scanln(&harga)

	c := Crypto{Nama: nama, Harga: harga}
	daftarCrypto = append(daftarCrypto, c)
	fmt.Println("Data crypto berhasil ditambahkan.")
}

func lihatDaftar() {
	if len(daftarCrypto) == 0 {
		fmt.Println("Daftar crypto kosong.")
		return
	}
	fmt.Println("\n--- Daftar Crypto ---")
	for i, c := range daftarCrypto {
		fmt.Printf("%d. %s - $%.2f\n", i+1, c.Nama, c.Harga)
	}
}

func cariCrypto() {
	var keyword string
	fmt.Print("Masukkan nama crypto yang dicari: ")
	fmt.Scanln(&keyword)

	found := false
	for _, c := range daftarCrypto {
		if strings.EqualFold(c.Nama, keyword) {
			fmt.Printf("Ditemukan: %s - $%.2f\n", c.Nama, c.Harga)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Crypto tidak ditemukan.")
	}
}

func urutkanCrypto() {
	sort.Slice(daftarCrypto, func(i, j int) bool {
		return daftarCrypto[i].Harga > daftarCrypto[j].Harga
	})
	fmt.Println("Data berhasil diurutkan berdasarkan harga tertinggi.")
}
