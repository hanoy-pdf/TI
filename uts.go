package main

import (
	"fmt"
)

var (
	username string = "Hana"
	password string = "2406361416"
	histori  []string
)

type Pengguna struct {
	Nama         string
	NPM          string
	Umur         int
	Alamat       string
	JenisKelamin string
	MakananFav   string
	MinumanFav   string
}

type Buku struct {
	NamaBuku   string
	JumlahBuku int
}

var (
	dataPengguna = Pengguna{
		Nama:         "Hana",
		NPM:          "2406361416",
		Umur:         18,
		Alamat:       "Bekasi",
		JenisKelamin: "Perempuan",
		MakananFav:   "Mie Ayam",
		MinumanFav:   "Kopi",
	}

	daftarBuku = []Buku{
		{"Pemrograman", 10},
		{"Film", 5},
		{"Printing", 20},
	}
)

func main() {
	var inputUsername, inputPassword string

	fmt.Println("=== Selamat Datang di Perpustakaan Vokasi ===")

	fmt.Print("Silakan input username: ")
	fmt.Scan(&inputUsername)

	fmt.Print("Silakan input password: ")
	fmt.Scan(&inputPassword)

	if inputUsername != username || inputPassword != password {
		fmt.Println("Username atau password salah.")
		return
	}

	fmt.Println("Login sukses!!!")

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Lihat Informasi Pengguna Program")
		fmt.Println("2. Lihat Daftar Buku")
		fmt.Println("3. Tambah Daftar Buku")
		fmt.Println("4. Tambah Peminjaman Buku")
		fmt.Println("5. Histori Peminjaman Buku")
		fmt.Println("6. Keluar Dari Program")

		var pilihan int
		fmt.Print("Input menu yang anda inginkan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println("===Informasi Pengguna Program===")
			fmt.Printf("Nama: %s\nNPM: %s\nUmur: %d\nAlamat: %s\nJenis Kelamin: %s\nMakanan Favorit: %s\nMinuman Favorit: %s\n",
				dataPengguna.Nama, dataPengguna.NPM, dataPengguna.Umur, dataPengguna.Alamat,
				dataPengguna.JenisKelamin, dataPengguna.MakananFav, dataPengguna.MinumanFav)
		case 2:
			fmt.Println("===Daftar Buku===")
			for _, buku := range daftarBuku {
				fmt.Printf("Nama Buku: %s, Jumlah: %d\n", buku.NamaBuku, buku.JumlahBuku)
			}
		case 3:
			fmt.Println("===Tambahkan Daftar Buku===")
			var namaBuku string
			var jumlahBuku int
			fmt.Print("Masukkan nama buku yang ingin ditambahkan: ")
			fmt.Scan(&namaBuku)
			fmt.Print("Masukkan jumlah buku yang ingin ditambahkan: ")
			fmt.Scan(&jumlahBuku)

			if jumlahBuku <= 0 {
				fmt.Println("Jumlah buku harus lebih besar dari 0.")
			} else {
				found := false
				for i, buku := range daftarBuku {
					if buku.NamaBuku == namaBuku {
						daftarBuku[i].JumlahBuku += jumlahBuku
						histori = append(histori, fmt.Sprintf("Ditambahkan: %s, Jumlah: %d", buku.NamaBuku, jumlahBuku))
						fmt.Println("Buku berhasil ditambahkan.")
						found = true
						break
					}
				}
				if !found {
					daftarBuku = append(daftarBuku, Buku{NamaBuku: namaBuku, JumlahBuku: jumlahBuku})
					histori = append(histori, fmt.Sprintf("Ditambahkan: %s, Jumlah: %d", namaBuku, jumlahBuku))
					fmt.Println("Buku baru berhasil ditambahkan.")
				}
			}
		case 4:
			fmt.Println("===Peminjaman Buku===")
			var namaBuku string
			var jumlahPinjam int
			fmt.Print("Masukkan nama buku yang akan dipinjam: ")
			fmt.Scan(&namaBuku)
			fmt.Print("Masukkan jumlah buku yang akan dipinjam: ")
			fmt.Scan(&jumlahPinjam)

			for i, buku := range daftarBuku {
				if buku.NamaBuku == namaBuku {
					if jumlahPinjam <= 0 {
						fmt.Println("Jumlah peminjaman harus lebih besar dari 0.")
						break
					} else if buku.JumlahBuku < jumlahPinjam {
						fmt.Println("Jumlah buku tidak mencukupi!")
						break
					} else {
						daftarBuku[i].JumlahBuku -= jumlahPinjam
						histori = append(histori, fmt.Sprintf("Dipinjam: %s, Jumlah: %d", buku.NamaBuku, jumlahPinjam))
						fmt.Println("Peminjaman buku berhasil.")
						break
					}
				}
			}
		case 5:
			fmt.Println("===Histori Peminjaman Buku===")
			for _, pinjaman := range histori {
				fmt.Println(pinjaman)
			}
		case 6:
			fmt.Println("Terima kasih telah mengunjungi Perpustakaan Vokasi!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Mohon coba lagi.")
		}
	}
}
