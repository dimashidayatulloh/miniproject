package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/gorilla/mux"

	"github.com/dimashidayatulloh/miniproject/config"
	"github.com/dimashidayatulloh/miniproject/migrations"

	// Import repository, service, handler sesuai kebutuhan
	"github.com/dimashidayatulloh/miniproject/internal/app"
	"github.com/dimashidayatulloh/miniproject/internal/handler"
	"github.com/dimashidayatulloh/miniproject/internal/middleware"
	"github.com/dimashidayatulloh/miniproject/internal/repository"
)

func main() {
    // env
    if err := godotenv.Load(); err != nil {
        log.Println("Warning: .env file not found or not loaded!")
    }
    
    // Jalankan migrasi database
    migrations.RunMigration()

    // Koneksi ke database
    db, err := config.ConnectDB()
    if err != nil {
        log.Fatal("Gagal koneksi database:", err)
    }

    // Inisialisasi repository, service, dan handler
    userRepo := repository.NewUserRepository(db)
    tokoRepo := repository.NewTokoRepository(db)
    userService := app.NewUserService(userRepo, tokoRepo)
    tokoService := app.NewTokoService(tokoRepo)
    userHandler := handler.NewUserHandler(userService)
    tokoHandler := handler.NewTokoHandler(tokoService)
    alamatRepo := repository.NewAlamatRepository(db)
    alamatService := app.NewAlamatService(alamatRepo)
    alamatHandler := handler.NewAlamatHandler(alamatService)
    catRepo := repository.NewCategoryRepository(db)
    catService := app.NewCategoryService(catRepo)
    catHandler := handler.NewCategoryHandler(catService)
    produkRepo := repository.NewProdukRepository(db)
    produkService := app.NewProdukService(produkRepo)
    produkHandler := handler.NewProdukHandler(produkService)
    logProdukRepo := repository.NewLogProdukRepository(db)
    logProdukService := app.NewLogProdukService(logProdukRepo)
    logProdukHandler := handler.NewLogProdukHandler(logProdukService)
    trxRepo := repository.NewTrxRepository(db)
    trxService := app.NewTrxService(trxRepo)
    trxHandler := handler.NewTrxHandler(trxService)

    // Inisialisasi router
    r := mux.NewRouter()

    // Endpoint tanpa middleware (public)
    r.HandleFunc("/register", userHandler.Register).Methods("POST")
    r.HandleFunc("/login", userHandler.Login).Methods("POST")
    r.HandleFunc("/toko/me", tokoHandler.GetMyToko).Methods("GET")
    r.HandleFunc("/toko/me", tokoHandler.UpdateMyToko).Methods("PUT")

    r.HandleFunc("/alamat", alamatHandler.Create).Methods("POST")
    r.HandleFunc("/alamat", alamatHandler.GetAll).Methods("GET")
    r.HandleFunc("/alamat/{id}", alamatHandler.Update).Methods("PUT")
    r.HandleFunc("/alamat/{id}", alamatHandler.Delete).Methods("DELETE")

    r.HandleFunc("/category", catHandler.Create).Methods("POST")         // admin only
    r.HandleFunc("/category/{id}", catHandler.Update).Methods("PUT")     // admin only
    r.HandleFunc("/category/{id}", catHandler.Delete).Methods("DELETE")  // admin only
    r.HandleFunc("/category", catHandler.GetAll).Methods("GET")          // public

    r.HandleFunc("/produk", produkHandler.Create).Methods("POST")
    r.HandleFunc("/produk", produkHandler.GetAll).Methods("GET")
    r.HandleFunc("/produk/{id}", produkHandler.Update).Methods("PUT")
    r.HandleFunc("/produk/{id}", produkHandler.Delete).Methods("DELETE")
    r.HandleFunc("/produk/{id}", produkHandler.GetByID).Methods("GET")
    r.HandleFunc("/produk/toko/{id_toko}", produkHandler.GetByToko).Methods("GET")

    r.HandleFunc("/log_produk", logProdukHandler.Create).Methods("POST")
    r.HandleFunc("/log_produk", logProdukHandler.GetAll).Methods("GET")
    r.HandleFunc("/log_produk/{id}", logProdukHandler.GetByID).Methods("GET")

    r.HandleFunc("/trx", trxHandler.Create).Methods("POST")
    r.HandleFunc("/trx", trxHandler.GetAll).Methods("GET")
    r.HandleFunc("/trx/{id}", trxHandler.GetByID).Methods("GET")

    // Endpoint dengan middleware JWT (protected)
    api := r.PathPrefix("/api").Subrouter()
    api.Use(middleware.JWTAuth) // pasang middleware JWT di subrouter ini

    // Contoh endpoint protected
    api.HandleFunc("/profile", userHandler.Profile).Methods("GET")
    // Tambahkan endpoint protected lain di sini

    log.Println("Server running at :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal("Gagal menjalankan server:", err)
    }
}