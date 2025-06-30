package domain

import "time"

type LogProduk struct {
    ID            int        `gorm:"primaryKey" json:"id"`
    IdProduk      int        `json:"id_produk"`
    NamaProduk    string     `json:"nama_produk"`
    Slug          string     `json:"slug"`
    HargaReseller int        `json:"harga_reseller"`
    HargaKonsumen int        `json:"harga_konsumen"`
    Stok          int        `json:"stok"`
    Deskripsi     string     `json:"deskripsi"`
    CreatedAt     *time.Time `json:"created_at"`
    UpdatedAt     *time.Time `json:"updated_at"`
    IdToko        int        `json:"id_toko"`
    IdCategory    int        `json:"id_category"`
}

func (LogProduk) TableName() string {
    return "log_produk"
}