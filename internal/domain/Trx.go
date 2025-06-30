package domain

import "time"

type Trx struct {
    ID               int        `gorm:"primaryKey" json:"id"`
    IdUser           int        `json:"id_user"`
    AlamatPengiriman int        `json:"alamat_pengiriman"`
    HargaTotal       int        `json:"harga_total"`
    KodeInvoice      string     `json:"kode_invoice"`
    MethodBayar      string     `json:"method_bayar"`
    UpdatedAt        *time.Time `json:"updated_at"`
    CreatedAt        *time.Time `json:"created_at"`
}

func (Trx) TableName() string {
    return "trx"
}