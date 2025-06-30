package domain

import "time"

type Toko struct {
    ID         int       `gorm:"primaryKey"`
    IdUser     int       `json:"id_user"`
    NamaToko   string    `json:"nama_toko"`
    UrlFoto    string    `json:"url_foto"`
    UpdatedAt  *time.Time
    CreatedAt  *time.Time
}

func (Toko) TableName() string {
    return "toko"
}