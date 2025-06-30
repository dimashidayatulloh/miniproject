package domain

import "time"

type FotoProduk struct {
    ID        int        `gorm:"primaryKey" json:"id"`
    IdProduk  int        `json:"id_produk"`
    URL       string     `json:"url"`
    UpdatedAt *time.Time `json:"updated_at"`
    CreatedAt *time.Time `json:"created_at"`
}