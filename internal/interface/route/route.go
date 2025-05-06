package routes

import (
	controller "perpustakaan/internal/interface/controllers"
	"perpustakaan/internal/repository"
	"perpustakaan/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Inisialisasi repository
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

	// Inisialisasi usecase
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

	// Inisialisasi controller
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

	// Routing Anggota
	anggota := r.Group("/anggota")
	{
		anggota.GET("", func(c *gin.Context) {
			anggotaController.HandleAnggota(c.Writer, c.Request)
		})
		anggota.POST("", func(c *gin.Context) {
			anggotaController.HandleAnggota(c.Writer, c.Request)
		})
		anggota.GET("/:id", func(c *gin.Context) {
			anggotaController.HandleAnggotaByID(c.Writer, c.Request)
		})
		anggota.PUT("/:id", func(c *gin.Context) {
			anggotaController.HandleAnggotaByID(c.Writer, c.Request)
		})
		anggota.DELETE("/:id", func(c *gin.Context) {
			anggotaController.HandleAnggotaByID(c.Writer, c.Request)
		})
	}

	// Routing Buku
	buku := r.Group("/buku")
	{
		buku.GET("", func(c *gin.Context) {
			bukuController.HandleBuku(c.Writer, c.Request)
		})
		buku.POST("", func(c *gin.Context) {
			bukuController.HandleBuku(c.Writer, c.Request)
		})
		buku.GET("/:id", func(c *gin.Context) {
			bukuController.HandleBukuByID(c.Writer, c.Request)
		})
		buku.PUT("/:id", func(c *gin.Context) {
			bukuController.HandleBukuByID(c.Writer, c.Request)
		})
		buku.DELETE("/:id", func(c *gin.Context) {
			bukuController.HandleBukuByID(c.Writer, c.Request)
		})
	}

	// Routing Kategori Buku
	kategori := r.Group("/kategori-buku")
	{
		kategori.GET("", func(c *gin.Context) {
			kategoriController.HandleKategoriBuku(c.Writer, c.Request)
		})
		kategori.POST("", func(c *gin.Context) {
			kategoriController.HandleKategoriBuku(c.Writer, c.Request)
		})
		kategori.GET("/:id", func(c *gin.Context) {
			kategoriController.HandleKategoriBukuByID(c.Writer, c.Request)
		})
		kategori.PUT("/:id", func(c *gin.Context) {
			kategoriController.HandleKategoriBukuByID(c.Writer, c.Request)
		})
		kategori.DELETE("/:id", func(c *gin.Context) {
			kategoriController.HandleKategoriBukuByID(c.Writer, c.Request)
		})
	}

	// Routing Penerbit
	penerbit := r.Group("/penerbit")
	{
		penerbit.GET("", func(c *gin.Context) {
			penerbitController.HandlePenerbit(c.Writer, c.Request)
		})
		penerbit.POST("", func(c *gin.Context) {
			penerbitController.HandlePenerbit(c.Writer, c.Request)
		})
		penerbit.GET("/:id", func(c *gin.Context) {
			penerbitController.HandlePenerbitByID(c.Writer, c.Request)
		})
		penerbit.PUT("/:id", func(c *gin.Context) {
			penerbitController.HandlePenerbitByID(c.Writer, c.Request)
		})
		penerbit.DELETE("/:id", func(c *gin.Context) {
			penerbitController.HandlePenerbitByID(c.Writer, c.Request)
		})
	}

	// Routing Penulis
	penulis := r.Group("/penulis")
	{
		penulis.GET("", func(c *gin.Context) {
			penulisController.HandlePenulis(c.Writer, c.Request)
		})
		penulis.POST("", func(c *gin.Context) {
			penulisController.HandlePenulis(c.Writer, c.Request)
		})
		penulis.GET("/:id", func(c *gin.Context) {
			penulisController.HandlePenulisByID(c.Writer, c.Request)
		})
		penulis.PUT("/:id", func(c *gin.Context) {
			penulisController.HandlePenulisByID(c.Writer, c.Request)
		})
		penulis.DELETE("/:id", func(c *gin.Context) {
			penulisController.HandlePenulisByID(c.Writer, c.Request)
		})
	}

	// Routing Petugas
	petugas := r.Group("/petugas")
	{
		petugas.GET("", func(c *gin.Context) {
			petugasController.HandlePetugas(c.Writer, c.Request)
		})
		petugas.POST("", func(c *gin.Context) {
			petugasController.HandlePetugas(c.Writer, c.Request)
		})
		petugas.GET("/:id", func(c *gin.Context) {
			petugasController.HandlePetugasByID(c.Writer, c.Request)
		})
		petugas.PUT("/:id", func(c *gin.Context) {
			petugasController.HandlePetugasByID(c.Writer, c.Request)
		})
		petugas.DELETE("/:id", func(c *gin.Context) {
			petugasController.HandlePetugasByID(c.Writer, c.Request)
		})
	}

	// Routing Peminjaman
	peminjaman := r.Group("/peminjaman")
	{
		peminjaman.GET("", func(c *gin.Context) {
			peminjamanController.HandlePeminjaman(c.Writer, c.Request)
		})
		peminjaman.POST("", func(c *gin.Context) {
			peminjamanController.HandlePeminjaman(c.Writer, c.Request)
		})
		peminjaman.GET("/:id", func(c *gin.Context) {
			peminjamanController.HandlePeminjamanByID(c.Writer, c.Request)
		})
		peminjaman.PUT("/:id", func(c *gin.Context) {
			peminjamanController.HandlePeminjamanByID(c.Writer, c.Request)
		})
		peminjaman.DELETE("/:id", func(c *gin.Context) {
			peminjamanController.HandlePeminjamanByID(c.Writer, c.Request)
		})
	}

	// Routing Detail Peminjaman
	detailPeminjaman := r.Group("/detail-peminjaman")
	{
		detailPeminjaman.GET("", func(c *gin.Context) {
			detailPeminjamanController.HandleDetailPeminjaman(c.Writer, c.Request)
		})
		detailPeminjaman.POST("", func(c *gin.Context) {
			detailPeminjamanController.HandleDetailPeminjaman(c.Writer, c.Request)
		})
		detailPeminjaman.GET("/:id", func(c *gin.Context) {
			detailPeminjamanController.HandleDetailPeminjamanByID(c.Writer, c.Request)
		})
		detailPeminjaman.PUT("/:id", func(c *gin.Context) {
			detailPeminjamanController.HandleDetailPeminjamanByID(c.Writer, c.Request)
		})
		detailPeminjaman.DELETE("/:id", func(c *gin.Context) {
			detailPeminjamanController.HandleDetailPeminjamanByID(c.Writer, c.Request)
		})
	}

	// Routing Pengembalian
	pengembalian := r.Group("/pengembalian")
	{
		pengembalian.GET("", func(c *gin.Context) {
			pengembalianController.HandlePengembalian(c.Writer, c.Request)
		})
		pengembalian.POST("", func(c *gin.Context) {
			pengembalianController.HandlePengembalian(c.Writer, c.Request)
		})
		pengembalian.GET("/:id", func(c *gin.Context) {
			pengembalianController.HandlePengembalianByID(c.Writer, c.Request)
		})
		pengembalian.PUT("/:id", func(c *gin.Context) {
			pengembalianController.HandlePengembalianByID(c.Writer, c.Request)
		})
		pengembalian.DELETE("/:id", func(c *gin.Context) {
			pengembalianController.HandlePengembalianByID(c.Writer, c.Request)
		})
	}

	// Routing Denda
	denda := r.Group("/denda")
	{
		denda.GET("", func(c *gin.Context) {
			dendaController.HandleDenda(c.Writer, c.Request)
		})
		denda.POST("", func(c *gin.Context) {
			dendaController.HandleDenda(c.Writer, c.Request)
		})
		denda.GET("/:id", func(c *gin.Context) {
			dendaController.HandleDendaByID(c.Writer, c.Request)
		})
		denda.PUT("/:id", func(c *gin.Context) {
			dendaController.HandleDendaByID(c.Writer, c.Request)
		})
		denda.DELETE("/:id", func(c *gin.Context) {
			dendaController.HandleDendaByID(c.Writer, c.Request)
		})
	}

	return r
}
