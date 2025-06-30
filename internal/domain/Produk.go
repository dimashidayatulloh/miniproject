package domain

import "time"

type Produk struct {
    ID            int        `gorm:"primaryKey"`
    NamaProduk    string     `json:"nama_produk"`
    Slug          string     `json:"slug"`
    HargaReseller int        `json:"harga_reseller"`
    HargaKonsumen int        `json:"harga_konsumen"`
    Stok          int        `json:"stok"`
    Deskripsi     string     `json:"deskripsi"`
    CreatedAt     *time.Time
    UpdatedAt     *time.Time
    IdToko        int        `json:"id_toko"`
    IdCategory    int        `json:"id_category"`
}