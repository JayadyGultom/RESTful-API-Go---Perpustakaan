package router

import (
	"perpustakaan/internal/interface/controllers"

	"github.com/gorilla/mux"
)

func SetupRouter(
	anggotaController *controllers.AnggotaController,
	kategoriController *controllers.KategoriBukuController,
	penulisController *controllers.PenulisController,
	penerbitController *controllers.PenerbitController,
	petugasController *controllers.PetugasController,
	bukuController *controllers.BukuController,
	dendaController *controllers.DendaController,
	detailPeminjamanController *controllers.DetailPeminjamanController,
	logAktivitasController *controllers.LogAktivitasController,
	peminjamanController *controllers.PeminjamanController,
	pengembalianController *controllers.PengembalianController,
	reservasiController *controllers.ReservasiController,
) *mux.Router {
	r := mux.NewRouter()

	// Anggota
	r.HandleFunc("/api/anggota", anggotaController.Create).Methods("POST")
	r.HandleFunc("/api/anggota/{id}", anggotaController.Update).Methods("PUT")
	r.HandleFunc("/api/anggota/{id}", anggotaController.Delete).Methods("DELETE")
	r.HandleFunc("/api/anggota/{id}", anggotaController.GetByID).Methods("GET")
	r.HandleFunc("/api/anggota", anggotaController.GetAll).Methods("GET")

	// Kategori Buku
	r.HandleFunc("/api/kategori", kategoriController.Create).Methods("POST")
	r.HandleFunc("/api/kategori/{id}", kategoriController.Update).Methods("PUT")
	r.HandleFunc("/api/kategori/{id}", kategoriController.Delete).Methods("DELETE")
	r.HandleFunc("/api/kategori/{id}", kategoriController.GetByID).Methods("GET")
	r.HandleFunc("/api/kategori", kategoriController.GetAll).Methods("GET")
	r.HandleFunc("/api/kategori/search", kategoriController.GetByNama).Methods("GET")

	// Penulis
	r.HandleFunc("/api/penulis", penulisController.Create).Methods("POST")
	r.HandleFunc("/api/penulis/{id}", penulisController.Update).Methods("PUT")
	r.HandleFunc("/api/penulis/{id}", penulisController.Delete).Methods("DELETE")
	r.HandleFunc("/api/penulis/{id}", penulisController.GetByID).Methods("GET")
	r.HandleFunc("/api/penulis", penulisController.GetAll).Methods("GET")
	r.HandleFunc("/api/penulis/search", penulisController.GetByNama).Methods("GET")

	// Penerbit
	r.HandleFunc("/api/penerbit", penerbitController.Create).Methods("POST")
	r.HandleFunc("/api/penerbit/{id}", penerbitController.Update).Methods("PUT")
	r.HandleFunc("/api/penerbit/{id}", penerbitController.Delete).Methods("DELETE")
	r.HandleFunc("/api/penerbit/{id}", penerbitController.GetByID).Methods("GET")
	r.HandleFunc("/api/penerbit", penerbitController.GetAll).Methods("GET")
	r.HandleFunc("/api/penerbit/search", penerbitController.GetByNama).Methods("GET")

	// Petugas
	r.HandleFunc("/api/petugas", petugasController.Create).Methods("POST")
	r.HandleFunc("/api/petugas/{id}", petugasController.Update).Methods("PUT")
	r.HandleFunc("/api/petugas/{id}", petugasController.Delete).Methods("DELETE")
	r.HandleFunc("/api/petugas/{id}", petugasController.GetByID).Methods("GET")
	r.HandleFunc("/api/petugas", petugasController.GetAll).Methods("GET")
	r.HandleFunc("/api/petugas/search", petugasController.GetByUsername).Methods("GET")

	// Buku
	r.HandleFunc("/api/buku", bukuController.Create).Methods("POST")
	r.HandleFunc("/api/buku/{id}", bukuController.Update).Methods("PUT")
	r.HandleFunc("/api/buku/{id}", bukuController.Delete).Methods("DELETE")
	r.HandleFunc("/api/buku/{id}", bukuController.GetByID).Methods("GET")
	r.HandleFunc("/api/buku", bukuController.GetAll).Methods("GET")

	// Denda
	r.HandleFunc("/api/denda", dendaController.Create).Methods("POST")
	r.HandleFunc("/api/denda/{id}", dendaController.Update).Methods("PUT")
	r.HandleFunc("/api/denda/{id}", dendaController.Delete).Methods("DELETE")
	r.HandleFunc("/api/denda/{id}", dendaController.GetByID).Methods("GET")
	r.HandleFunc("/api/denda", dendaController.GetAll).Methods("GET")

	// Detail Peminjaman
	r.HandleFunc("/api/detail-peminjaman", detailPeminjamanController.Create).Methods("POST")
	r.HandleFunc("/api/detail-peminjaman/{id}", detailPeminjamanController.Update).Methods("PUT")
	r.HandleFunc("/api/detail-peminjaman/{id}", detailPeminjamanController.Delete).Methods("DELETE")
	r.HandleFunc("/api/detail-peminjaman/{id}", detailPeminjamanController.GetByID).Methods("GET")
	r.HandleFunc("/api/detail-peminjaman", detailPeminjamanController.GetAll).Methods("GET")

	// Log Aktivitas
	r.HandleFunc("/api/log-aktivitas", logAktivitasController.Create).Methods("POST")
	r.HandleFunc("/api/log-aktivitas/{id}", logAktivitasController.Update).Methods("PUT")
	r.HandleFunc("/api/log-aktivitas/{id}", logAktivitasController.Delete).Methods("DELETE")
	r.HandleFunc("/api/log-aktivitas/{id}", logAktivitasController.GetByID).Methods("GET")
	r.HandleFunc("/api/log-aktivitas", logAktivitasController.GetAll).Methods("GET")

	// Peminjaman
	r.HandleFunc("/api/peminjaman", peminjamanController.Create).Methods("POST")
	r.HandleFunc("/api/peminjaman/{id}", peminjamanController.Update).Methods("PUT")
	r.HandleFunc("/api/peminjaman/{id}", peminjamanController.Delete).Methods("DELETE")
	r.HandleFunc("/api/peminjaman/{id}", peminjamanController.GetByID).Methods("GET")
	r.HandleFunc("/api/peminjaman", peminjamanController.GetAll).Methods("GET")

	// Pengembalian
	r.HandleFunc("/api/pengembalian", pengembalianController.Create).Methods("POST")
	r.HandleFunc("/api/pengembalian/{id}", pengembalianController.Update).Methods("PUT")
	r.HandleFunc("/api/pengembalian/{id}", pengembalianController.Delete).Methods("DELETE")
	r.HandleFunc("/api/pengembalian/{id}", pengembalianController.GetByID).Methods("GET")
	r.HandleFunc("/api/pengembalian", pengembalianController.GetAll).Methods("GET")

	// Reservasi
	r.HandleFunc("/api/reservasi", reservasiController.Create).Methods("POST")
	r.HandleFunc("/api/reservasi/{id}", reservasiController.Update).Methods("PUT")
	r.HandleFunc("/api/reservasi/{id}", reservasiController.Delete).Methods("DELETE")
	r.HandleFunc("/api/reservasi/{id}", reservasiController.GetByID).Methods("GET")
	r.HandleFunc("/api/reservasi", reservasiController.GetAll).Methods("GET")

	return r
} 