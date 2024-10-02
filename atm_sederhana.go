package main

import (
	"fmt"
)

// Struktur data untuk menyimpan informasi akun
type Account struct {
	username string
	PIN      int
	Balance  float64
}

// Fungsi utama program
func main() {
	// Membuat akun contoh (dalam aplikasi nyata, data akun akan disimpan dalam database)
	var accounts = []Account{
		{username: "Hanoy", PIN: 2406361416, Balance: 3500000.0},
	}

	var (
		username string
		pin      int
	)

	// Meminta input nomor akun dan PIN dari pengguna
	fmt.Print("Masukkan nama akun: ")
	fmt.Scanln(&username)

	fmt.Print("Masukkan PIN: ")
	fmt.Scanln(&pin)

	// Mencari akun yang sesuai
	var foundAccount *Account
	for i := range accounts {
		if accounts[i].username == username && accounts[i].PIN == pin {
			foundAccount = &accounts[i]
			break
		}
	}

	// Jika akun ditemukan, tampilkan menu ATM
	if foundAccount != nil {
		fmt.Println("Selamat datang!")

		for {
			fmt.Println("\nPilih transaksi:")
			fmt.Println("1. Lihat Informasi Akun")
			fmt.Println("2. Lihat saldo")
			fmt.Println("3. Tarik tunai")
			fmt.Println("4. Setor tunai")
			fmt.Println("5. Riwayat transaksi")
			fmt.Println("6. Keluar")

			var choice int
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				fmt.Println("Ini Adalah Informasi Akun Anda!")
				fmt.Println("Username: Hanoy")
				fmt.Println("PIN: 2406361416")
			case 2:
				fmt.Printf("Saldo Anda: %.2f\n", foundAccount.Balance)
			case 3:
				var amount float64
				fmt.Print("Masukkan jumlah penarikan: ")
				fmt.Scanln(&amount)

				if amount > 0 && amount <= foundAccount.Balance {
					foundAccount.Balance -= amount
					fmt.Printf("Penarikan berhasil. Saldo baru: %.2f\n", foundAccount.Balance)
				} else {
					fmt.Println("Saldo tidak mencukupi atau jumlah penarikan tidak valid.")
				}
			case 4:
				var amount float64
				fmt.Print("Masukkan jumlah setoran: ")
				fmt.Scanln(&amount)

				if amount > 0 {
					foundAccount.Balance += amount
					fmt.Printf("Setoran berhasil. Saldo baru: %.2f\n", foundAccount.Balance)
				} else {
					fmt.Println("Jumlah setoran tidak valid.")
				}
			case 5:
				{
					fmt.Println("Riwayat Transaksi")
					fmt.Printf("Saldo anda saat ini tersisa: %.2f\n", foundAccount.Balance)
				}
			case 6:
				fmt.Println("Terima kasih telah menggunakan layanan kami!")
				return
			default:
				fmt.Println("Pilihan tidak valid.")
			}
		}
	} else {
		fmt.Println("Nomor akun atau PIN salah.")
	}
}
