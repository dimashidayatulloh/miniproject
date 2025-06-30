package domain

import "time"

type Alamat struct {
    ID           int       `gorm:"primaryKey"`
    IdUser       int       `json:"id_user"`
    JudulAlamat  string    `json:"judul_alamat"`
    NamaPenerima string    `json:"nama_penerima"`
    NoTelp       string    `json:"no_telp"`
    DetailAlamat string    `json:"detail_alamat"`
    UpdatedAt    *time.Time
    CreatedAt    *time.Time
}