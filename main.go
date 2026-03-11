package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type transaction struct {
	id          int    // Id unique transaction
	amount      int    // nominal of money
	description string // description of transaction
	category    string // category of transaction
	date        string // date of transaction
	types       string // type of transaction (income or expense)
}

func readLine(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func centerText(s string, width int) string {
	padding := width - len(s)
	if padding <= 0 {
		return s
	}
	left := padding / 2
	right := padding - left
	return strings.Repeat(" ", left) + s + strings.Repeat(" ", right)
}

func normalizeDate(dateStr string) string {
	layouts := []string{
		"02/01/2006",
		"2/1/2006",
		"02-01-2006",
		"2-1-2006",
		"2006-01-02",
		"02/01/06",
		"2/1/06",
		"02012006",
		"20060102",
		"02.01.2006",
		"2.1.2006",
		"2006.01.02",
		"2.1.06",
		"02.01.06",
	}

	for _, layout := range layouts {
		t, err := time.Parse(layout, dateStr)
		if err == nil {
			return t.Format("02/01/2006")
		}
	}
	return dateStr
}

func clearScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

var transactions []transaction
var nextId int = 1
var lengtestId, lengtestAmount, lengtestDesc, lengtestCat, lengtestDate, lengtestTypes int = 0, 0, 0, 0, 0, 0

var reader = bufio.NewReader(os.Stdin)

func showMenus() {
	fmt.Println("=== Expense Tracker CLI ===")
	fmt.Println("1. Tambah Pengeluaran")
	fmt.Println("2. Lihat Pengeluaran")
	fmt.Println("3. Hapus Pengeluaran")
	fmt.Println("4. Keluar")
	fmt.Print("Pilih menu: ")
}

func add() {
	input := readLine("Masukkan Transaksi (Format: Judul Jumlah Kategori Tanggal Tipe): ")
	if input == "" {
		fmt.Println("Input tidak boleh kosong")
		return
	}

	// Split input berdasarkan spasi
	parts := strings.Fields(input)

	// Minimal harus ada 5 bagian: Judul(min 1), Jumlah, Kategori, Tanggal, Tipe
	if len(parts) < 5 {
		fmt.Println("Format salah! Harus ada: Judul Jumlah Kategori Tanggal Tipe")
		fmt.Println("Contoh: Kabel USB 20000 Elektronik 20/03/2026 Expense")
		return
	}

	// Ambil 4 bagian terakhir
	n := len(parts)
	types := parts[n-1]
	rawDate := parts[n-2]
	category := parts[n-3]
	amountStr := parts[n-4]

	// Normalisasi tanggal
	date := normalizeDate(rawDate)

	// Sisanya di depan digabung kembali sebagai Judul (Description)
	title := strings.Join(parts[:n-4], " ")

	// Konversi amount
	amount, err := strconv.Atoi(amountStr)
	if err != nil || amount <= 0 {
		fmt.Println("Jumlah (Amount) harus berupa angka positif!")
		return
	}

	transaction := transaction{
		id:          nextId,
		amount:      amount,
		description: title,
		category:    category,
		date:        date,
		types:       types,
	}
	transactions = append(transactions, transaction)
	nextId++
	fmt.Println("Transaksi berhasil ditambahkan!")
}

func view() {
	if len(transactions) == 0 {
		fmt.Println("Belum ada transaksi.")
		return
	}

	// Reset dan tentukan lebar minimal berdasarkan nama kolom
	lengtestId = 2
	lengtestDesc = 5
	lengtestAmount = 6
	lengtestCat = 8
	lengtestDate = 7
	lengtestTypes = 4

	// Hitung lebar maksimal
	for _, t := range transactions {
		idStr := strconv.Itoa(t.id)
		if len(idStr) > lengtestId {
			lengtestId = len(idStr)
		}
		amountStr := strconv.Itoa(t.amount)
		if len(amountStr) > lengtestAmount {
			lengtestAmount = len(amountStr)
		}
		if len(t.description) > lengtestDesc {
			lengtestDesc = len(t.description)
		}
		if len(t.category) > lengtestCat {
			lengtestCat = len(t.category)
		}
		if len(t.date) > lengtestDate {
			lengtestDate = len(t.date)
		}
		if len(t.types) > lengtestTypes {
			lengtestTypes = len(t.types)
		}
	}

	lineLen := lengtestId + lengtestDesc + lengtestAmount + lengtestCat + lengtestDate + lengtestTypes + 19
	fmt.Println(strings.Repeat("=", lineLen))

	// Cetak Header (Center)
	fmt.Printf("| %s | %s | %s | %s | %s | %s |\n",
		centerText("ID", lengtestId),
		centerText("Judul", lengtestDesc),
		centerText("Jumlah", lengtestAmount),
		centerText("Kategori", lengtestCat),
		centerText("Tanggal", lengtestDate),
		centerText("Tipe", lengtestTypes))

	// Garis pemisah
	fmt.Println(strings.Repeat("-", lineLen))

	// Baris data
	for _, t := range transactions {
		fmt.Printf("| %s | %s | %s | %s | %s | %s |\n",
			centerText(strconv.Itoa(t.id), lengtestId),
			centerText(t.description, lengtestDesc),
			centerText(strconv.Itoa(t.amount), lengtestAmount),
			centerText(t.category, lengtestCat),
			centerText(t.date, lengtestDate),
			centerText(t.types, lengtestTypes))
	}
	fmt.Println(strings.Repeat("=", lineLen))
}

func main() {
	fmt.Print("Wilujeng Sumping! ^_^\n\n")
	showMenus()
	for {
		choiceStr := readLine("")
		choice, _ := strconv.Atoi(choiceStr)

		switch choice {
		case 1:
			clearScreen()
			add()
			showMenus()
		case 2:
			clearScreen()
			view()
			showMenus()
		case 3:
			clearScreen()
			fmt.Println("Hapus Pengeluaran (Belum diimplementasi)")
		case 4:
			clearScreen()
			fmt.Println("Terima Kasih telah mencoba! ^_^")
			return
		default:
			clearScreen()
			fmt.Println("Pilihan tidak valid")
		}
	}
}
