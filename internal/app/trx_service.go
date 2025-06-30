package app

import (
	"errors"
	"math/rand"
	"time"

	"github.com/dimashidayatulloh/miniproject/internal/domain"
	"github.com/dimashidayatulloh/miniproject/internal/repository"
)

type TrxService struct {
	repo         *repository.TrxRepository
	alamatRepo   *repository.AlamatRepository // Tambahkan dependency repository alamat
}

func NewTrxService(repo *repository.TrxRepository, alamatRepo *repository.AlamatRepository) *TrxService {
	return &TrxService{repo, alamatRepo}
}

type TrxInput struct {
	AlamatPengiriman int              `json:"alamat_pengiriman"`
	MethodBayar      string           `json:"method_bayar"`
	HargaTotal       int              `json:"harga_total"`
	Detail           []DetailTrxInput `json:"detail"`
}

type DetailTrxInput struct {
	IdLogProduk int `json:"id_log_produk"`
	IdToko      int `json:"id_toko"`
	Kuantitas   int `json:"kuantitas"`
	HargaTotal  int `json:"harga_total"`
}

func (s *TrxService) CreateTrx(userID int, input *TrxInput) error {
	// Validasi: pastikan alamat milik user
	alamat, err := s.alamatRepo.FindByID(input.AlamatPengiriman, userID)
	if err != nil || alamat == nil {
		return errors.New("alamat pengiriman tidak ditemukan")
	}

	trx := &domain.Trx{
		IdUser:           userID,
		AlamatPengiriman: input.AlamatPengiriman,
		MethodBayar:      input.MethodBayar,
		HargaTotal:       input.HargaTotal,
		KodeInvoice:      generateInvoiceCode(),
	}
	var details []domain.DetailTrx
	for _, d := range input.Detail {
		details = append(details, domain.DetailTrx{
			IdLogProduk: d.IdLogProduk,
			IdToko:      d.IdToko,
			Kuantitas:   d.Kuantitas,
			HargaTotal:  d.HargaTotal,
		})
	}
	return s.repo.Create(trx, details)
}

func (s *TrxService) GetAllTrx(userID int) ([]domain.Trx, error) {
	return s.repo.FindByUser(userID)
}

func (s *TrxService) GetTrxByID(userID int, trxID int) (*domain.Trx, []domain.DetailTrx, error) {
	trx, err := s.repo.FindByID(trxID, userID)
	if err != nil {
		return nil, nil, err
	}
	detail, err := s.repo.FindDetailsByTrx(trxID)
	return trx, detail, err
}

// helper
func generateInvoiceCode() string {
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return "INV-" + string(b)
}

func (s *TrxService) GetAllTrxPaginatedFiltered(
	userID, page, limit int,
	kodeInvoice, metode, tanggal string,
	minTotal, maxTotal int,
) ([]domain.Trx, int64, error) {
	return s.repo.FindByUserPaginatedFiltered(userID, page, limit, kodeInvoice, metode, tanggal, minTotal, maxTotal)
}