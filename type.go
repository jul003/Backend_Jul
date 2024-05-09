package backendjulgo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Gadget struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Merek       string             `bson:"merek,omitempty" json:"merek,omitempty"`
	Model       string             `bson:"model,omitempty" json:"model,omitempty"`
	Tipe        string             `bson:"tipe,omitempty" json:"tipe,omitempty"`
	Harga       float64            `bson:"harga,omitempty" json:"harga,omitempty"`
	Warna       string             `bson:"warna,omitempty" json:"warna,omitempty"`
	TahunRilis  int                `bson:"tahun_rilis,omitempty" json:"tahun_rilis,omitempty"`
	Keterangan  string             `bson:"keterangan,omitempty" json:"keterangan,omitempty"`
	Spesifikasi []string           `bson:"spesifikasi,omitempty" json:"spesifikasi,omitempty"`
	Stok        int                `bson:"stok,omitempty" json:"stok,omitempty"`
	Garansi     int                `bson:"garansi,omitempty" json:"garansi,omitempty"`
}

type Pengguna struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Email        string             `bson:"email,omitempty" json:"email,omitempty"`
	Alamat       string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	NomorTelepon string             `bson:"nomor_telepon,omitempty" json:"nomor_telepon,omitempty"`
	TipePengguna string             `bson:"tipe_pengguna,omitempty" json:"tipe_pengguna,omitempty"`
	Status       string             `bson:"status,omitempty" json:"status,omitempty"`
	JenisKelamin string             `bson:"jenis_kelamin,omitempty" json:"jenis_kelamin,omitempty"`
	Umur         int                `bson:"umur,omitempty" json:"umur,omitempty"`
	Transaksi    []Transaksi        `bson:"transaksi,omitempty" json:"transaksi,omitempty"`
}

type Transaksi struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Tanggal        primitive.DateTime `bson:"tanggal,omitempty" json:"tanggal,omitempty"`
	TotalHarga     float64            `bson:"total_harga,omitempty" json:"total_harga,omitempty"`
	Pembayaran     string             `bson:"pembayaran,omitempty" json:"pembayaran,omitempty"`
	Status         string             `bson:"status,omitempty" json:"status,omitempty"`
	Gadget         Gadget             `bson:"gadget,omitempty" json:"gadget,omitempty"`
	MetodePengiriman string             `bson:"metode_pengiriman,omitempty" json:"metode_pengiriman,omitempty"`
	AlamatPengiriman string             `bson:"alamat_pengiriman,omitempty" json:"alamat_pengiriman,omitempty"`
	DurasiPengiriman int                `bson:"durasi_pengiriman,omitempty" json:"durasi_pengiriman,omitempty"`
	BiayaPengiriman  float64            `bson:"biaya_pengiriman,omitempty" json:"biaya_pengiriman,omitempty"`
}


