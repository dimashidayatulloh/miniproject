package domain

import "time"

type Category struct {
    ID           int        `gorm:"primaryKey" json:"id"`
    NamaCategory string     `json:"nama_category"`
    CreatedAt    *time.Time `json:"created_at"`
    UpdatedAt    *time.Time `json:"updated_at"`
}