package domain

import "time"

type DetailTrx struct {
    ID          int        `gorm:"primaryKey" json:"id"`
    IdTrx       int        `json:"id_trx"`
    IdLogProduk int        `json:"id_log_produk"`
    IdToko      int        `json:"id_toko"`
    Kuantitas   int        `json:"kuantitas"`
    HargaTotal  int        `json:"harga_total"`
    UpdatedAt   *time.Time `json:"updated_at"`
    CreatedAt   *time.Time `json:"created_at"`
}

func (DetailTrx) TableName() string {
    return "detail_trx"
}