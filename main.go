package main

import "fmt"

func showMenus() {
	fmt.Println("=== Expense Tracker CLI ===")
	fmt.Println("1. Tambah Pengeluaran")
	fmt.Println("2. Lihat Pengeluaran")
	fmt.Println("3. Hapus Pengeluaran")
	fmt.Println("4. Keluar")
	fmt.Print("Pilih menu: ")
}

func main() {
	var choice int

	for {
		showMenus()
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Tambah Pengeluaran")
		case 2:
			fmt.Println("Lihat Pengeluaran")
		case 3:
			fmt.Println("Hapus Pengeluaran")
		case 4:
			fmt.Println("Keluar")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
