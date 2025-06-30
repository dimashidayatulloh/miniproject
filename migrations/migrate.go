package migrations

import (
	"log"

	"github.com/dimashidayatulloh/miniproject/config"
	"github.com/dimashidayatulloh/miniproject/internal/domain"
)

func RunMigration() {
    db, err := config.ConnectDB()
    if err != nil {
        log.Fatal("Gagal koneksi database:", err)
    }

    err = db.AutoMigrate(
        &domain.User{},
        &domain.Category{},
        &domain.LogProduk{},
        &domain.FotoProduk{},
        &domain.Trx{},
        &domain.DetailTrx{},
        &domain.Alamat{},
        &domain.Produk{},
        &domain.Toko{},
    )
    if err != nil {
        log.Fatal("Gagal migrasi:", err)
    }

    log.Println("Migrasi sukses!")
}