package main

import (
    "fmt"
    "log"
    "net/http"
    "perpustakaan/config"
    controller "perpustakaan/internal/interface/controllers"
    "perpustakaan/internal/interface/middleware"
    "perpustakaan/internal/repository"
    "perpustakaan/internal/usecase"
)

func main() {
    // Initialize database connection
    config.ConnectDB()
    db := config.DB

    // Repository
    anggotaRepo := repository.NewAnggotaRepository(db)
    bukuRepo := repository.NewBukuRepository(db)
    kategoriRepo := repository.NewKategoriBukuRepository(db)
    penerbitRepo := repository.NewPenerbitRepository(db)
    penulisRepo := repository.NewPenulisRepository(db)
    petugasRepo := repository.NewPetugasRepository(db)
    peminjamanRepo := repository.NewPeminjamanRepository(db)
    detailPeminjamanRepo := repository.NewDetailPeminjamanRepository(db)
    pengembalianRepo := repository.NewPengembalianRepository(db)
    dendaRepo := repository.NewDendaRepository(db)

    // Usecase
    anggotaUsecase := usecase.NewAnggotaUsecase(anggotaRepo)
    bukuUsecase := usecase.NewBukuUsecase(bukuRepo)
    kategoriUsecase := usecase.NewKategoriBukuUsecase(kategoriRepo)
    penerbitUsecase := usecase.NewPenerbitUsecase(penerbitRepo)
    penulisUsecase := usecase.NewPenulisUsecase(penulisRepo)
    petugasUsecase := usecase.NewPetugasUsecase(petugasRepo)
    peminjamanUsecase := usecase.NewPeminjamanUsecase(peminjamanRepo)
    detailPeminjamanUsecase := usecase.NewDetailPeminjamanUsecase(detailPeminjamanRepo)
    pengembalianUsecase := usecase.NewPengembalianUsecase(pengembalianRepo)
    dendaUsecase := usecase.NewDendaUsecase(dendaRepo)

    // Controller
    anggotaController := controller.NewAnggotaController(anggotaUsecase)
    bukuController := controller.NewBukuController(bukuUsecase)
    kategoriController := controller.NewKategoriBukuController(kategoriUsecase)
    penerbitController := controller.NewPenerbitController(penerbitUsecase)
    penulisController := controller.NewPenulisController(penulisUsecase)
    petugasController := controller.NewPetugasController(petugasUsecase)
    peminjamanController := controller.NewPeminjamanController(peminjamanUsecase)
    detailPeminjamanController := controller.NewDetailPeminjamanController(detailPeminjamanUsecase)
    pengembalianController := controller.NewPengembalianController(pengembalianUsecase)
    dendaController := controller.NewDendaController(dendaUsecase)

    // HTTP Handler
    http.HandleFunc("/anggota/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        if r.URL.Path == "/anggota" || r.URL.Path == "/anggota/" {
            anggotaController.HandleAnggota(w, r)
        } else {
            anggotaController.HandleAnggotaByID(w, r)
        }
    })

    http.HandleFunc("/buku/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        if r.URL.Path == "/buku" || r.URL.Path == "/buku/" {
            bukuController.HandleBuku(w, r)
        } else {
            bukuController.HandleBukuByID(w, r)
        }
    })

    http.HandleFunc("/kategori-buku/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        if r.URL.Path == "/kategori-buku" || r.URL.Path == "/kategori-buku/" {
            kategoriController.HandleKategoriBuku(w, r)
        } else {
            kategoriController.HandleKategoriBukuByID(w, r)
        }
    })

    http.HandleFunc("/penerbit/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        if r.URL.Path == "/penerbit" || r.URL.Path == "/penerbit/" {
            penerbitController.HandlePenerbit(w, r)
        } else {
            penerbitController.HandlePenerbitByID(w, r)
        }
    })

    http.HandleFunc("/penulis/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        if r.URL.Path == "/penulis" || r.URL.Path == "/penulis/" {
            penulisController.HandlePenulis(w, r)
        } else {
            penulisController.HandlePenulisByID(w, r)
        }
    })

    http.HandleFunc("/petugas/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        if r.URL.Path == "/petugas" || r.URL.Path == "/petugas/" {
            petugasController.HandlePetugas(w, r)
        } else {
            petugasController.HandlePetugasByID(w, r)
        }
    })

    http.HandleFunc("/peminjaman/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        if r.URL.Path == "/peminjaman" || r.URL.Path == "/peminjaman/" {
            peminjamanController.HandlePeminjaman(w, r)
        } else {
            peminjamanController.HandlePeminjamanByID(w, r)
        }
    })

    http.HandleFunc("/detail-peminjaman/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        if r.URL.Path == "/detail-peminjaman" || r.URL.Path == "/detail-peminjaman/" {
            detailPeminjamanController.HandleDetailPeminjaman(w, r)
        } else {
            detailPeminjamanController.HandleDetailPeminjamanByID(w, r)
        }
    })

    http.HandleFunc("/pengembalian/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        if r.URL.Path == "/pengembalian" || r.URL.Path == "/pengembalian/" {
            pengembalianController.HandlePengembalian(w, r)
        } else {
            pengembalianController.HandlePengembalianByID(w, r)
        }
    })

    http.HandleFunc("/denda/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        if r.URL.Path == "/denda" || r.URL.Path == "/denda/" {
            dendaController.HandleDenda(w, r)
        } else {
            dendaController.HandleDendaByID(w, r)
        }
    })

    // Add middleware
    handler := middleware.Logger(middleware.CORSMiddleware(http.DefaultServeMux))

    // Server information and endpoints
    port := ":8080"
    fmt.Printf("========================================================================\n")
    fmt.Printf("Server running on port %s...\n", port)
    fmt.Printf("Available endpoints:\n")
    fmt.Printf("Anggota:\n")
    fmt.Printf("  GET, POST: http://localhost%s/anggota\n", port)
    fmt.Printf("  GET, PUT, DELETE: http://localhost%s/anggota/{id}\n", port)
    fmt.Printf("Buku:\n")
    fmt.Printf("  GET, POST: http://localhost%s/buku\n", port)
    fmt.Printf("  GET, PUT, DELETE: http://localhost%s/buku/{id}\n", port)
    fmt.Printf("Kategori Buku:\n")
    fmt.Printf("  GET, POST: http://localhost%s/kategori-buku\n", port)
    fmt.Printf("  GET, PUT, DELETE: http://localhost%s/kategori-buku/{id}\n", port)
    fmt.Printf("Penerbit:\n")
    fmt.Printf("  GET, POST: http://localhost%s/penerbit\n", port)
    fmt.Printf("  GET, PUT, DELETE: http://localhost%s/penerbit/{id}\n", port)
    fmt.Printf("Penulis:\n")
    fmt.Printf("  GET, POST: http://localhost%s/penulis\n", port)
    fmt.Printf("  GET, PUT, DELETE: http://localhost%s/penulis/{id}\n", port)
    fmt.Printf("Petugas:\n")
    fmt.Printf("  GET, POST: http://localhost%s/petugas\n", port)
    fmt.Printf("  GET, PUT, DELETE: http://localhost%s/petugas/{id}\n", port)
    fmt.Printf("Peminjaman:\n")
    fmt.Printf("  GET, POST: http://localhost%s/peminjaman\n", port)
    fmt.Printf("  GET, PUT, DELETE: http://localhost%s/peminjaman/{id}\n", port)
    fmt.Printf("Detail Peminjaman:\n")
    fmt.Printf("  GET, POST: http://localhost%s/detail-peminjaman\n", port)
    fmt.Printf("  GET, PUT, DELETE: http://localhost%s/detail-peminjaman/{id}\n", port)
    fmt.Printf("Pengembalian:\n")
    fmt.Printf("  GET, POST: http://localhost%s/pengembalian\n", port)
    fmt.Printf("  GET, PUT, DELETE: http://localhost%s/pengembalian/{id}\n", port)
    fmt.Printf("Denda:\n")
    fmt.Printf("  GET, POST: http://localhost%s/denda\n", port)
    fmt.Printf("  GET, PUT, DELETE: http://localhost%s/denda/{id}\n", port)
    fmt.Printf("========================================================================\n")

    // Start server
    err := http.ListenAndServe(port, handler)
    if err != nil {
        log.Fatal(err)
    }
}
