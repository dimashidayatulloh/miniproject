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
    userService := app.NewUserService(userRepo)
    userHandler := handler.NewUserHandler(userService)

    // Inisialisasi router
    r := mux.NewRouter()

    // Endpoint tanpa middleware (public)
    r.HandleFunc("/register", userHandler.Register).Methods("POST")
    r.HandleFunc("/login", userHandler.Login).Methods("POST")

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