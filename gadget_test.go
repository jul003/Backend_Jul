package backendjulgo_test

import (
	"os"
	"testing"
	"github.com/jul003/BackEnd_Jul" 
)

func TestInsertAndGetAll(t *testing.T) {
	
	os.Setenv("MONGOSTRING", "mongodb://localhost:27017")
	defer os.Unsetenv("MONGOSTRING")

	
	testGadget := backendjulgo.Gadget{
		Merek:       "Asus",
		Model:       "ROG",
		Tipe:        "Strix",
		Harga:       110000000,
		Warna:       "Blue canon",
		TahunRilis:  2024,
		Keterangan:  "Test Keterangan",
		Spesifikasi: []string{"Spec 1", "Spec 2"},
		Stok:        10,
		Garansi:     1,
	}

	testPengguna := backendjulgo.Pengguna{
		Nama:         "Test Nama",
		Email:        "test@example.com",
		Alamat:       "Test Alamat",
		NomorTelepon: "1234567890",
		TipePengguna: "Test Tipe Pengguna",
		Status:       "Aktif",
		JenisKelamin: "Laki-laki",
		Umur:         30,
	}

	testTransaksi := backendjulgo.Transaksi{
		Tanggal:        1234567890,
		TotalHarga:     500.75,
		Pembayaran:     "Kartu Kredit",
		Status:         "Selesai",
		Gadget:         testGadget,
		MetodePengiriman: "Kurir",
		AlamatPengiriman: "Test Alamat Pengiriman",
		DurasiPengiriman: 3,
		BiayaPengiriman: 20.50,
	}

	
	gadgetID, err := backendjulgo.GadgetInsertOne(testGadget)
	if err != nil {
		t.Fatalf("Failed to insert gadget: %v", err)
	}

	penggunaID, err := backendjulgo.PenggunaInsertOne(testPengguna)
	if err != nil {
		t.Fatalf("Failed to insert pengguna: %v", err)
	}

	transaksiID, err := backendjulgo.TransaksiInsertOne(testTransaksi)
	if err != nil {
		t.Fatalf("Failed to insert transaksi: %v", err)
	}

	
	gadgets, err := backendjulgo.GadgetGetAll()
	if err != nil {
		t.Errorf("Failed to get gadgets: %v", err)
	}
	if len(gadgets) == 0 {
		t.Error("Expected gadgets, got none")
	}

	penggunas, err := backendjulgo.PenggunaGetAll()
	if err != nil {
		t.Errorf("Failed to get penggunas: %v", err)
	}
	if len(penggunas) == 0 {
		t.Error("Expected penggunas, got none")
	}

	transaksis, err := backendjulgo.TransaksiGetAll()
	if err != nil {
		t.Errorf("Failed to get transaksis: %v", err)
	}
	if len(transaksis) == 0 {
		t.Error("Expected transaksis, got none")
	}

	
	var found bool
	for _, gadget := range gadgets {
		if gadget.ID == gadgetID {
			found = true
			break
		}
	}
	if !found {
		t.Error("Test gadget not found")
	}

	found = false
	for _, pengguna := range penggunas {
		if pengguna.ID == penggunaID {
			found = true
			break
		}
	}
	if !found {
		t.Error("Test pengguna not found")
	}

	found = false
	for _, transaksi := range transaksis {
		if transaksi.ID == transaksiID {
			found = true
			break
		}
	}
	if !found {
		t.Error("Test transaksi not found")
	}
}
