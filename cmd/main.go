package main

import (
	"fmt"
	"log"
	"net/http"
	database "perpustakaan/config"
	"perpustakaan/internal/interface/controllers"
	"perpustakaan/internal/interface/router"
	"perpustakaan/internal/repository"
	"perpustakaan/internal/usecase"
	"perpustakaan/internal/interface/middleware"
)

func main() {
	db := database.ConnectDB()

	// Repository
	anggotaRepo := repository.NewAnggotaRepository(db)
	kategoriRepo := repository.NewKategoriBukuRepository(db)
	penulisRepo := repository.NewPenulisRepository(db)
	penerbitRepo := repository.NewPenerbitRepository(db)
	petugasRepo := repository.NewPetugasRepository(db)
	bukuRepo := repository.NewBukuRepository(db)
	dendaRepo := repository.NewDendaRepository(db)
	detailPeminjamanRepo := repository.NewDetailPeminjamanRepository(db)
	logAktivitasRepo := repository.NewLogAktivitasRepository(db)
	peminjamanRepo := repository.NewPeminjamanRepository(db)
	pengembalianRepo := repository.NewPengembalianRepository(db)
	reservasiRepo := repository.NewReservasiRepository(db)

	// Usecase
	anggotaUsecase := usecase.NewAnggotaUsecase(anggotaRepo)
	kategoriUsecase := usecase.NewKategoriBukuUsecase(kategoriRepo)
	penulisUsecase := usecase.NewPenulisUsecase(penulisRepo)
	penerbitUsecase := usecase.NewPenerbitUsecase(penerbitRepo)
	petugasUsecase := usecase.NewPetugasUsecase(petugasRepo)
	bukuUsecase := usecase.NewBukuUsecase(bukuRepo)
	dendaUsecase := usecase.NewDendaUsecase(dendaRepo)
	detailPeminjamanUsecase := usecase.NewDetailPeminjamanUsecase(detailPeminjamanRepo)
	logAktivitasUsecase := usecase.NewLogAktivitasUsecase(logAktivitasRepo)
	peminjamanUsecase := usecase.NewPeminjamanUsecase(peminjamanRepo)
	pengembalianUsecase := usecase.NewPengembalianUsecase(pengembalianRepo)
	reservasiUsecase := usecase.NewReservasiUsecase(reservasiRepo)

	// Controller
	anggotaController := controllers.NewAnggotaController(anggotaUsecase)
	kategoriController := controllers.NewKategoriBukuController(kategoriUsecase)
	penulisController := controllers.NewPenulisController(penulisUsecase)
	penerbitController := controllers.NewPenerbitController(penerbitUsecase)
	petugasController := controllers.NewPetugasController(petugasUsecase)
	bukuController := controllers.NewBukuController(bukuUsecase)
	dendaController := controllers.NewDendaController(dendaUsecase)
	detailPeminjamanController := controllers.NewDetailPeminjamanController(detailPeminjamanUsecase)
	logAktivitasController := controllers.NewLogAktivitasController(logAktivitasUsecase)
	peminjamanController := controllers.NewPeminjamanController(peminjamanUsecase)
	pengembalianController := controllers.NewPengembalianController(pengembalianUsecase)
	reservasiController := controllers.NewReservasiController(reservasiUsecase)

	// Router
	r := router.SetupRouter(
		anggotaController,
		kategoriController,
		penulisController,
		penerbitController,
		petugasController,
		bukuController,
		dendaController,
		detailPeminjamanController,
		logAktivitasController,
		peminjamanController,
		pengembalianController,
		reservasiController,
	)

	port := ":8080"
	baseURL := fmt.Sprintf("http://localhost%s", port)
	
	fmt.Println("\n=== Perpustakaan API Server ===")
	fmt.Println("Server berhasil dijalankan!")
	fmt.Printf("Base URL: %s\n", baseURL)
	fmt.Println("\nEndpoint yang tersedia:")
	fmt.Printf("1. Anggota:\n   %s/api/anggota\n", baseURL)
	fmt.Printf("2. Buku:\n   %s/api/buku\n", baseURL)
	fmt.Printf("3. Kategori Buku:\n   %s/api/kategori-buku\n", baseURL)
	fmt.Printf("4. Penulis:\n   %s/api/penulis\n", baseURL)
	fmt.Printf("5. Penerbit:\n   %s/api/penerbit\n", baseURL)
	fmt.Printf("6. Petugas:\n   %s/api/petugas\n", baseURL)
	fmt.Printf("7. Denda:\n   %s/api/denda\n", baseURL)
	fmt.Printf("8. Detail Peminjaman:\n   %s/api/detail-peminjaman\n", baseURL)
	fmt.Printf("9. Log Aktivitas:\n   %s/api/log-aktivitas\n", baseURL)
	fmt.Printf("10. Peminjaman:\n   %s/api/peminjaman\n", baseURL)
	fmt.Printf("11. Pengembalian:\n   %s/api/pengembalian\n", baseURL)
	fmt.Printf("12. Reservasi:\n   %s/api/reservasi\n", baseURL)
	fmt.Println("\nUntuk setiap endpoint di atas, tersedia operasi:")
	fmt.Println("   - GET    /api/[endpoint]         (mendapatkan semua data)")
	fmt.Println("   - GET    /api/[endpoint]/{id}   (mendapatkan data berdasarkan ID)")
	fmt.Println("   - POST   /api/[endpoint]        (menambah data baru)")
	fmt.Println("   - PUT    /api/[endpoint]/{id}   (mengupdate data)")
	fmt.Println("   - DELETE /api/[endpoint]/{id}   (menghapus data)")
	fmt.Println("\nKhusus untuk Petugas, tersedia endpoint tambahan:")
	fmt.Printf("   - GET    %s/api/petugas/username?username={username}\n", baseURL)
	fmt.Println("\nTekan Ctrl+C untuk menghentikan server")
	fmt.Println("==============================\n")

	log.Fatal(http.ListenAndServe(port, middleware.CORSMiddleware(r)))
}
