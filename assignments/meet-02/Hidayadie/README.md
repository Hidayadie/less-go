# CLI Kasir Sederhana

Mini project ini melatih:

- variable
- input dengan `fmt.Scanln`
- output dengan `fmt.Println` dan `fmt.Printf`
- operator aritmatika
- conditional
- loop
- function
- slice
- struct
- error handling

## Cara Run

```bash
go run ./mini-project/cli-kasir
```

## Fitur

- Input nama barang.
- Input harga barang.
- Input jumlah barang.
- Hitung subtotal.
- Hitung total belanja.
- Beri diskon 10% jika total minimal Rp100.000.
- Validasi jumlah item, harga, dan jumlah barang harus lebih dari 0.
- Tampilkan error jika input tidak valid dan kembali ke menu.
- Reset keranjang.
- Keluar dari aplikasi lewat menu.

## Contoh Input

```text
=== CLI Kasir Sederhana ===
1. Tambah barang
2. Lihat keranjang
3. Cetak struk
4. Reset keranjang
5. Keluar
Pilih menu: 1
Nama barang: Buku
Harga barang: 50000
Jumlah barang: 2
Pilih menu: 1
Nama barang: Pulpen
Harga barang: 5000
Jumlah barang: 3
Pilih menu: 3
```

## Cek

```bash
go test ./mini-project/cli-kasir
```
