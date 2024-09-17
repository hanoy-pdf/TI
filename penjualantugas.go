package main

import "fmt"

func main() {

	//deklarasi variabel
	var nama_customer, nama_barang string
	var harga float32 = 25000
	var quantity int32
	var hasil_discount, total_harga float32

	//input
	fmt.Print("Input Nama Customer: ")
	fmt.Scanf("%s\n", &nama_customer)

	fmt.Print("Input Nama Barang: ")
	fmt.Scanf("%s\n", &nama_barang)

	fmt.Print("Input quantity barang: ")
	fmt.Scanf("%d\n", &quantity)

	// kondisi diskon, kalau lebih dari 5 dapat 10%, selain itu 2%
	switch quantity {
	case 2:
		hasil_discount = 0.4
	default:
		hasil_discount = 0.08
	}

	//proses
	sub_total := float32(quantity) * harga
	total_harga = sub_total - (hasil_discount * sub_total)

	// tampilkan hasil harga
	fmt.Println("hasil_discount: ", hasil_discount)
	fmt.Println("quantity: ", quantity)
	fmt.Println("harga: ", harga)
	fmt.Println("Total harga adalah: ", total_harga)

}
