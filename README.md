# Go Expense Tracker CLI

Aplikasi pencatat keuangan simple (CLI/Terminal Based).

### Fitur Utama:
- **Add**: Input 1 baris (Format: Judul Jumlah Kategori Tanggal Tipe).
  - *Detail About Description*: Judul bisa lebih dari satu kata (e.g. `Beli Kopi Susu`).
  - *Auto Format              : Tanggal otomatis dinormalkan ke format `DD/MM/YYYY`.
  - *Auto Lowercase*          : Tipe transaksi otomatis diubah menjadi huruf kecil (expense/income).
- **Read**: 
  - Lihat Data & Ringkasan (Total & Saldo).
- **Update** & **Delete** berdasarkan ID.

### Contoh Input:
```text
Beli Kopi Susu 25000 Makan 11-03-2026 Expense
```

### Cara Menjalankan:
```bash
go run main.go