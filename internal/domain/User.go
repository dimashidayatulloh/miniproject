package domain

import "time"

type User struct {
    ID            int       `gorm:"primaryKey" json:"id"`
    Nama          string    `json:"nama"`
    KataSandi     string    `json:"kata_sandi"`
    Notelp        string    `gorm:"unique" json:"notelp"`
    Slug          string    `json:"slug"`
    TanggalLahir  *time.Time `json:"tanggal_lahir"`
    JenisKelamin  string    `json:"jenis_kelamin"`
    TempatTinggal string    `json:"tempat_tinggal"`
    Pekerjaan     string    `json:"pekerjaan"`
    Email         string    `json:"email"`
    IdProvinsi    string    `json:"id_provinsi"`
    IdKota        string    `json:"id_kota"`
    IsAdmin       *bool     `json:"isAdmin"`
    UpdatedAt     *time.Time `json:"updated_at"`
    CreatedAt     *time.Time `json:"created_at"`
}