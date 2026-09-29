package main

import "testing"

func TestHitungTotalDanDiskon(t *testing.T) {
	items := []Item{
		{Nama: "Buku", Harga: 50000, Jumlah: 2, Subtotal: hitungSubtotal(50000, 2)},
		{Nama: "Pulpen", Harga: 5000, Jumlah: 3, Subtotal: hitungSubtotal(5000, 3)},
	}

	total := hitungTotal(items)
	if total != 115000 {
		t.Fatalf("total = %d, want 115000", total)
	}

	diskon := hitungDiskon(total)
	if diskon != 11500 {
		t.Fatalf("diskon = %d, want 11500", diskon)
	}
}
