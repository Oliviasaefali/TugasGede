package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Aset struct {
	Nama         string
	Harga        float64
	Kapitalisasi float64
	Simbol       string
}

type Transaksi struct {
	Jenis  string
	Aset   string
	Jumlah float64
	harga  float64
	time   time.Time
}

type AsetPribadi struct {
	Nama   string
	jumlah int
	Harga  float64
	Simbol string
}

var daftarAset []Aset
var riwayat []Transaksi
var asetku []AsetPribadi
var saldo float64 = 1000000

// Tampilan visual untuk menampilkan Headder
func tampilkanHeader(judul string) {
	panjang := (65 - len(judul)) / 2
	fmt.Println(strings.Repeat("=", 65))
	fmt.Println(strings.Repeat(" ", panjang), judul, strings.Repeat(" ", panjang))
	fmt.Println(strings.Repeat("=", 65))
}

// Tampilan visual untuk menghapus text pada terminal
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// Tampilan visual agar menekan Enter sebelum berpindah ke halaman lain
func waitForEnter() {
	fmt.Println(strings.Repeat("=", 65))
	fmt.Print("\n⦿ Tekan Enter untuk melanjutkan!")

	reader := bufio.NewReader(os.Stdin)

	reader.ReadString('\n')
	reader.ReadString('\n')
}

// Menampilkan menu utama
func tampilkanMenu() {
	clearScreen()
	tampilkanHeader("Menu Simulasi Kripto Sederhana")
	fmt.Println("1. Kelola Daftar Aset")
	fmt.Println("2. Jual Beli Aset")
	fmt.Println("3. Portfolio")
	fmt.Println("4. Lihat Riwayat Transaksi")
	fmt.Println("5. Cari Aset")
	fmt.Println("6. Urutkan Aset")
	fmt.Println("7. Tambah Saldo")
	fmt.Println("0. Keluar")
	fmt.Print("\n↳ Pilihan Anda : ")
}

// Menambahkan aset Kripto
func tambahAset() {
	var nama, simbol string
	var harga, kapital float64

	tampilkanHeader("Tambahkan Aset Kripto")
	fmt.Print("↳ Nama Aset : ")
	fmt.Scan(&nama)
	for _, aset := range daftarAset {
		if aset.Nama == nama {
			fmt.Printf("\n☒ Aset dengan nama %s sudah tersedia.\n", nama)
			waitForEnter()
			clearScreen()
			return
		}
	}
	fmt.Print("↳ Simbol Aset : ")
	fmt.Scan(&simbol)
	fmt.Print("↳ Harga : ")
	fmt.Scan(&harga)
	fmt.Print("↳ Kapitalisasi: ")
	fmt.Scan(&kapital)

	// Proses penambahkan aset Kripto ke dalam array
	daftarAset = append(daftarAset, Aset{nama, harga, kapital, simbol})

	fmt.Printf("\n☑ Aset Kripto %s Berhasil Ditambahkan.\n", nama)
	waitForEnter()
	clearScreen()
}

// Mengubah data aset Kripto
func ubahAset() {
	tampilkanHeader("Ubah Aset")
	tampilkanAset()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\n↳ Masukkan Nama Aset yang Ingin Diubah : ")
	// flush newline dari input sebelumnya
	namaInput, _ := reader.ReadString('\n')
	nama := strings.TrimSpace(namaInput)

	// Mencari data aset
	for i := range daftarAset {
		if strings.EqualFold(daftarAset[i].Nama, nama) {
			// Nama baru
			fmt.Print("↳ Nama Aset Baru (Enter jika tidak diubah) : ")
			newNama, _ := reader.ReadString('\n')
			newNama = strings.TrimSpace(newNama)
			if newNama != "" {
				daftarAset[i].Nama = newNama
			}

			// Harga baru
			fmt.Print("↳ Harga Baru (Enter jika tidak diubah) : ")
			hargaStr, _ := reader.ReadString('\n')
			hargaStr = strings.TrimSpace(hargaStr)
			if hargaStr != "" {
				if hargaBaru, err := strconv.ParseFloat(hargaStr, 64); err == nil {
					daftarAset[i].Harga = hargaBaru
				} else {
					fmt.Println("☒ Input harga tidak valid, harga tidak diubah.")
				}
			}

			// Kapitalisasi baru
			fmt.Print("↳ Kapitalisasi Baru (Enter jika tidak diubah) : ")
			kapStr, _ := reader.ReadString('\n')
			kapStr = strings.TrimSpace(kapStr)
			if kapStr != "" {
				if kapBaru, err := strconv.ParseFloat(kapStr, 64); err == nil {
					daftarAset[i].Kapitalisasi = kapBaru
				} else {
					fmt.Println("☒ Input kapitalisasi tidak valid, kapitalisasi tidak diubah.")
				}
			}

			fmt.Printf("\n☑ Data aset %s berhasil diperbarui.\n", nama)
			fmt.Println(strings.Repeat("=", 65))
			fmt.Print("\nTekan Enter untuk melanjutkan...")
			reader.ReadString('\n')
			clearScreen()
			return
		}
	}

	fmt.Printf("Aset %s tidak ditemukan.\n", nama)
	fmt.Println(strings.Repeat("=", 65))
	fmt.Print("\nTekan Enter untuk melanjutkan...")
	reader.ReadString('\n')

	clearScreen()
}

// Menghapus data aset kripto
func hapusAset() {
	var nama string
	tampilkanHeader("Hapus Aset")
	tampilkanAset()
	fmt.Print("\n↳ Masukkan Nama Aset : ")
	fmt.Scan(&nama)

	// Mencari data aset yang sesuai dengan inputan
	for i := range daftarAset {
		// Proses pencocokan nama dari daftar aset dengan nama yang diinputkan
		if strings.EqualFold(daftarAset[i].Nama, nama) {
			daftarAset = append(daftarAset[:i], daftarAset[i+1:]...)
			fmt.Printf("☑ Aset %s berhasil dihapus.\n", nama)
			waitForEnter()
			clearScreen()
			return
		}
	}
	fmt.Printf("☒ Aset %s tidak ditemukan.\n", nama)

	waitForEnter()
	clearScreen()
}

// Membeli aset yang nanti akan disimpan ke aray struct AsetPribadi
func beliAset() {
	var nama string
	var jumlah float64
	tampilkanHeader("Beli Aset")
	fmt.Printf("$ Saldo Kamu : %s\n\n", formatMataUang(saldo))
	tampilkanAset()
	fmt.Print("\n↳ Nama Aset yang Ingin Dibeli : ")
	fmt.Scan(&nama)

	// Mencari data aset yang sesuai dengan inputan
	for _, aset := range daftarAset {
		// Proses pencocokan nama dari daftar aset dengan nama yang diinputkan
		if strings.EqualFold(aset.Nama, nama) {
			fmt.Print("↳ Jumlah Pembelian : ")
			fmt.Scan(&jumlah)
			total := aset.Harga * jumlah
			if saldo >= total {
				saldo -= total
				riwayat = append(riwayat, Transaksi{"Beli", nama, jumlah, total, time.Now()})
				found := false

				// Pengecekan apakah sudah pernag beli aset yang sama sebelumnnya, jika sudah hanya mengupdate properti jumlah dan harga
				for i := range asetku {
					if strings.EqualFold(asetku[i].Nama, nama) {
						asetku[i].jumlah += int(jumlah)
						asetku[i].Harga += total
						found = true
						break
					}
				}
				// Jika tidak maka akan menambahkan aset tersebut ke AsetPribadi
				if !found {
					asetku = append(asetku, AsetPribadi{nama, int(jumlah), total, aset.Simbol})
				}

				fmt.Printf("$ Total harga : %s\n", formatMataUang(total))
				fmt.Printf("\n☑ Pembelian %s sebanyak %d berhasil.\n", nama, int(jumlah))
				fmt.Println("Sisa Saldo :", formatMataUang(saldo))
			} else {
				fmt.Printf("$ Total harga : %s\n", formatMataUang(total))
				fmt.Println("\n☒ Saldo tidak mencukupi.")
			}
			waitForEnter()
			clearScreen()
			return
		}
	}
	fmt.Println("\n☒ Aset tidak ditemukan.")
	waitForEnter()
	clearScreen()
}

// Menjual AsetPribadi
func jualAset() {
	var nama string
	var jumlah float64
	tampilkanHeader("Jual Aset")

	// Harus memiliki AsetPribadi (tidak boleh kosong)
	if len(asetku) != 0 {
		tampilkanAsetku()
		fmt.Print("\n↳ Nama Aset yang Ingin Dijual : ")
		fmt.Scan(&nama)

		// Mencari data AsetPribadi yang sesuai dengan inputan
		for i := range asetku {
			// Proses pencocokan nama dari daftar AsetPribadi dengan nama yang diinputkan
			if strings.EqualFold(asetku[i].Nama, nama) {
				fmt.Print("↳ Jumlah Penjualan : ")
				fmt.Scan(&jumlah)

				// Jumlah aset yang dijual tidak boleh lebih dari yang dimmiliki
				if int(jumlah) > asetku[i].jumlah {
					fmt.Printf("\n☒ Jumlah aset %s tidak mencukupi.\n", nama)
				} else {
					for _, aset := range daftarAset {
						if strings.EqualFold(aset.Nama, nama) {

							// Proses penambahan saldo karena aset dijual, lalu mengurangi Harga dan Jumlah dari AsetPribadi terkait
							saldo += aset.Harga * jumlah
							asetku[i].Harga -= aset.Harga * jumlah
							totalHarga := aset.Harga * jumlah

							asetku[i].jumlah -= int(jumlah)
							riwayat = append(riwayat, Transaksi{"Jual", nama, jumlah, float64(totalHarga), time.Now()})
						}
					}

					// Jika salah satu AsetPribadi dijual hingga habis, maka akan dihapus dari Array Struct AsetPribadi
					if asetku[i].jumlah == 0 {
						asetku = append(asetku[:i], asetku[i+1:]...)
					}
					fmt.Printf("\n☑ Penjualan aset %s sebanyak %d berhasil.\n", nama, int(jumlah))
				}
				waitForEnter()
				clearScreen()
				return
			}
		}
		fmt.Println("☒ Aset tidak ditemukan.")
		waitForEnter()
		clearScreen()
	}
	fmt.Println("☒ Anda tidak memiliki aset apapun.")

	waitForEnter()
	clearScreen()
}

// Melihat Riwayat Transaksi (Jual, Beli)
func lihatRiwayat() {
	tampilkanHeader("Riwayat Transaksi")

	fmt.Println("Daftar Riwayat Transaksi.\n")
	fmt.Print("| Tipe", strings.Repeat(" ", 6-len("Tipe")))
	fmt.Print("| Nama Aset", strings.Repeat(" ", 14-len("Nama Aset")))
	fmt.Print("| Jumlah", strings.Repeat(" ", 8-len("Jumlah")))
	fmt.Print("| Harga", strings.Repeat(" ", 14-len("Harga")))
	fmt.Print("| Waktu", strings.Repeat(" ", 22-len("Waktu")))
	fmt.Println()
	fmt.Println(strings.Repeat("-", 65))

	for _, r := range riwayat {
		jumlahStr := strconv.FormatFloat(r.Jumlah, 'f', 0, 64)
		hargaStr := formatMataUang(r.harga)
		waktuStr := r.time.Format("2006-01-02 15:04")

		fmt.Print("| ", r.Jenis, strings.Repeat(" ", 6-len(r.Jenis)))
		fmt.Print("| ", r.Aset, strings.Repeat(" ", 14-len(r.Aset)))
		fmt.Print("| ", r.Jumlah, strings.Repeat(" ", 8-len(jumlahStr)))
		fmt.Print("| ", hargaStr, strings.Repeat(" ", 14-len(hargaStr)))
		fmt.Println("| ", waktuStr, strings.Repeat(" ", 22-len(waktuStr)))
	}

	waitForEnter()
	clearScreen()
}

// Mencari Aset dengan metode Sequential Search
func cariSequential() {
	var nama string
	tampilkanHeader("Pencarian Aset")
	fmt.Print("↳ Nama Aset yang ingin Dicari : ")
	fmt.Scan(&nama)
	fmt.Println()
	for _, aset := range daftarAset {
		if strings.EqualFold(aset.Nama, nama) {
			fmt.Println("Nama Aset : ", aset.Nama)
			fmt.Println("Simbol : ", aset.Simbol)
			fmt.Println("Harga : ", formatMataUang(aset.Harga))
			fmt.Println("Kapitalisasi : ", formatMataUang(aset.Kapitalisasi))

			fmt.Println("\n☑ Aset Ditemukan.")
			waitForEnter()
			clearScreen()
			return
		}
	}
	fmt.Println("\n☒ Aset tidak ditemukan.")
	waitForEnter()
	clearScreen()
}

// Mencari Aset dengan metode Binary Search
func cariBinary() {
	// Binary search hanya dapat dilakukan untuk data yang sudah terurut, library sort mengurutkan secara langsung sebuah data (disini diurutkan berdasarkan alfabet secara ascending)
	sort.Slice(daftarAset, func(i, j int) bool {
		return strings.ToLower(daftarAset[i].Nama) < strings.ToLower(daftarAset[j].Nama)
	})
	var nama string
	tampilkanHeader("Pencarian Aset")
	fmt.Print("↳ Nama Aset yang  Ingin Dicari : ")
	fmt.Scan(&nama)
	fmt.Println()
	kiri, kanan := 0, len(daftarAset)-1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		comp := strings.Compare(strings.ToLower(nama), strings.ToLower(daftarAset[tengah].Nama))
		if comp == 0 {
			fmt.Println("Nama Aset : ", daftarAset[tengah].Nama)
			fmt.Println("Simbol : ", daftarAset[tengah].Simbol)
			fmt.Println("Harga : ", formatMataUang(daftarAset[tengah].Harga))
			fmt.Println("Kapitalisasi : ", formatMataUang(daftarAset[tengah].Kapitalisasi))
			fmt.Println("\n☑ Aset Ditemukan.")
			waitForEnter()
			clearScreen()
			return
		} else if comp < 0 {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	fmt.Println("\n☒ Aset tidak ditemukan.")
	waitForEnter()
	clearScreen()
}

// Mengurutkan aset secara Selection Sort
func urutkanSelection() {
	tampilkanHeader("Urutkan Aset")
	for i := 0; i < len(daftarAset)-1; i++ {
		min := i
		for j := i + 1; j < len(daftarAset); j++ {
			if daftarAset[j].Harga < daftarAset[min].Harga {
				min = j
			}
		}
		daftarAset[i], daftarAset[min] = daftarAset[min], daftarAset[i]
	}
	tampilkanAset()
	fmt.Println("\n☑ Aset diurutkan berdasarkan harga.")
	waitForEnter()
	clearScreen()
}

// Mengurutkan aset seacara Insertion Sort
func urutkanInsertion() {
	tampilkanHeader("Urutkan Aset")
	for i := 1; i < len(daftarAset); i++ {
		kunci := daftarAset[i]
		j := i - 1
		for j >= 0 && daftarAset[j].Kapitalisasi > kunci.Kapitalisasi {
			daftarAset[j+1] = daftarAset[j]
			j--
		}
		daftarAset[j+1] = kunci
	}
	tampilkanAset()
	fmt.Println("\n☑ Aset diurutkan berdasarkan kapitalisasi.")
	waitForEnter()
	clearScreen()
}

// Menampilkan seluruh aset
func tampilkanAset() {
	fmt.Print("| Nama Aset", strings.Repeat(" ", 15-len("Nama aset")))
	fmt.Print("| Simbol", strings.Repeat(" ", 7-len("Simbol")))
	fmt.Print("| Harga", strings.Repeat(" ", 14-len("Harga")))
	fmt.Print("| Kapitalisasi", strings.Repeat(" ", 24-len("Kapitalisasi")))
	fmt.Println()
	fmt.Println(strings.Repeat("-", 65))
	for _, aset := range daftarAset {
		fmt.Print("| ", aset.Nama, strings.Repeat(" ", 15-len(aset.Nama)))
		fmt.Print("| ", aset.Simbol, strings.Repeat(" ", 7-len(aset.Simbol)))
		fmt.Print("| ", formatMataUang(aset.Harga), strings.Repeat(" ", 14-len(formatMataUang(aset.Harga))))
		fmt.Print("| ", formatMataUang(aset.Kapitalisasi), strings.Repeat(" ", 24-len(formatMataUang(aset.Kapitalisasi))))
		fmt.Println()
	}
}

// Menampilkan seluruh aset pribadi
func tampilkanAsetku() {
	if len(asetku) != 0 {
		fmt.Println("\nAset Anda\n")
		fmt.Print("| Nama Aset", strings.Repeat(" ", 15-len("Nama aset")))
		fmt.Print("| Simbol", strings.Repeat(" ", 7-len("Simbol")))
		fmt.Print("| Jumlah", strings.Repeat(" ", 9-len("Jumlah")))
		fmt.Print("| Harga", strings.Repeat(" ", 14-len("Harga")))
		fmt.Println()
		fmt.Println(strings.Repeat("-", 65))
		for _, aset := range asetku {
			fmt.Print("| ", aset.Nama, strings.Repeat(" ", 15-len(aset.Nama)))
			fmt.Print("| ", aset.Simbol, strings.Repeat(" ", 7-len(aset.Simbol)))
			fmt.Print("| ", aset.jumlah, strings.Repeat(" ", 9-len(strconv.Itoa(aset.jumlah))))
			fmt.Print("| ", formatMataUang(aset.Harga), strings.Repeat(" ", 14-len(formatMataUang(aset.Harga))))
			fmt.Println()
		}
	}
}

// Menambahkan saldo
func tambahSaldo() {
	tampilkanHeader("Tambah Saldo")
	fmt.Println("$ Saldo Anda : ", formatMataUang(saldo))
	fmt.Print("\n↳ Masukkan jumlah saldo yang ingin ditambahkan : ")
	var jumlah float64

	fmt.Scan(&jumlah)
	saldo += jumlah
	fmt.Println("☑ Saldo berhasil ditambahkan.\n$ Saldo Terbaru : ", formatMataUang(saldo))

	waitForEnter()
	clearScreen()
}

// Berfungsi untuk melakukan format sebuah nilai agar lebih mudah dibaca (1000 menjadi 1,000)
func formatMataUang(angka float64) string {
	str := fmt.Sprintf("%.0f", angka)
	n := len(str)
	if n <= 3 {
		return str + " USD"
	}

	var hasil string
	sisa := n % 3
	if sisa > 0 {
		hasil = str[:sisa]
	}

	for i := sisa; i < n; i += 3 {
		if len(hasil) > 0 {
			hasil += ","
		}
		hasil += str[i : i+3]
	}
	return hasil + " USD"
}

func main() {

	daftarAset = []Aset{
		{"Bitcoin", 68000.0, 1300000000000.0, "BTC"},
		{"Ethereum", 3500.0, 420000000000.0, "ETH"},
		{"Binance Coin", 600.0, 95000000000.0, "BNB"},
		{"Solana", 180.0, 78000000000.0, "SOL"},
		{"Cardano", 5, 17000000000.0, "ADA"},
		{"XRP", 12, 32000000000.0, "XRP"},
		{"Dogecoin", 20, 18000000000.0, "DOGE"},
		{"Polkadot", 7.0, 9000000000.0, "DOT"},
	}

	for {
		tampilkanMenu()

		var pilihan int
		fmt.Scan(&pilihan)

		clearScreen()

		switch pilihan {
		case 1:
			for {
				clearScreen()
				tampilkanHeader("Kelola Daftar Aset")
				fmt.Println("1. Lihat Aset")
				fmt.Println("2. Tambah Aset")
				fmt.Println("3. Ubah Aset")
				fmt.Println("4. Hapus Aset")
				fmt.Println("0. Kembali")
				fmt.Print("\n↳ Pilihan Anda : ")

				var pilihKelola int
				fmt.Scan(&pilihKelola)

				clearScreen()

				switch pilihKelola {
				case 1:
					tampilkanHeader("Daftar Aset")
					tampilkanAset()
					waitForEnter()
					clearScreen()
					break
				case 2:
					tambahAset()
					break
				case 3:
					reader := bufio.NewReader(os.Stdin)
					reader.ReadString('\n')
					ubahAset()
					break
				case 4:
					hapusAset()
					break
				case 0:
					break
				default:
					fmt.Println("\n☒ Pilihan tidak valid")
					waitForEnter()
				}
				if pilihKelola == 1 || pilihKelola == 2 || pilihKelola == 0 || pilihKelola == 3 || pilihKelola == 4 {
					break
				}
			}
		case 2:
			for {
				clearScreen()
				tampilkanHeader("Jual Beli Aset")
				fmt.Println("1. Beli Aset")
				fmt.Println("2. Jual Aset")
				fmt.Println("0. Kembali")
				fmt.Print("\n↳ Pilihan Anda : ")

				var pilihJualBeli int
				fmt.Scan(&pilihJualBeli)

				clearScreen()

				switch pilihJualBeli {
				case 1:
					beliAset()
					break
				case 2:
					jualAset()
					break
				case 0:
					break
				default:
					fmt.Println("\n☒ Pilihan tidak valid")
					waitForEnter()
				}
				if pilihJualBeli == 1 || pilihJualBeli == 2 || pilihJualBeli == 0 {
					break
				}
			}
		case 3:
			tampilkanHeader("Portfolio Kamu")
			fmt.Printf("$ Saldo Anda : %s\n", formatMataUang(saldo))
			if len(asetku) != 0 {
			} else {
				fmt.Println("☒ Anda tidak memiliki aset apapun.")
			}
			tampilkanAsetku()
			waitForEnter()
		case 4:
			lihatRiwayat()
		case 5:
			for {
				clearScreen()
				tampilkanHeader("Cari Aset")
				fmt.Println("1. Sequential Search")
				fmt.Println("2. Binary Search")
				fmt.Println("0. Kembali")
				fmt.Print("\n↳ Pilihan Anda : ")

				var pilihCari int
				fmt.Scan(&pilihCari)

				clearScreen()

				switch pilihCari {
				case 1:
					cariSequential()
					break
				case 2:
					cariBinary()
					break
				case 0:
					break
				default:
					fmt.Println("\n☒ Pilihan tidak valid")
					waitForEnter()
				}
				if pilihCari == 1 || pilihCari == 2 || pilihCari == 0 {
					break
				}
			}
		case 6:
			for {
				clearScreen()
				tampilkanHeader("Urutkan Aset")
				fmt.Println("1. Berdasarkan Harga (Selection)")
				fmt.Println("2. Berdasarkan Kapitalisasi (Insertion)")
				fmt.Println("0. Kembali")
				fmt.Print("\n↳ Pilihan Anda : ")

				var pilihUrutan int
				fmt.Scan(&pilihUrutan)

				clearScreen()

				switch pilihUrutan {
				case 1:
					urutkanSelection()
					break
				case 2:
					urutkanInsertion()
					break
				case 0:
					break
				default:
					fmt.Println("\n☒ Pilihan tidak valid")
					waitForEnter()
				}
				if pilihUrutan == 1 || pilihUrutan == 2 || pilihUrutan == 0 {
					break
				}
			}
		case 7:
			tambahSaldo()
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi!")
			os.Exit(0)
		default:
			fmt.Println("☒ Pilihan tidak valid.")
			waitForEnter()
		}
	}
}
