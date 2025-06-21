# Perpustakaan REST API

Aplikasi REST API untuk sistem perpustakaan yang dibangun menggunakan Go.

## Fitur

- Manajemen Kategori Buku
- Manajemen Penulis
- Manajemen Penerbit
- Manajemen Petugas

## Persyaratan

- Go 1.21 atau lebih baru
- PostgreSQL 12 atau lebih baru

## Instalasi

1. Clone repository ini
2. Install dependensi:
   ```bash
   go mod download
   ```
3. Buat database PostgreSQL dengan nama `perpustakaan`
4. Sesuaikan konfigurasi database di `cmd/main.go` jika diperlukan
5. Jalankan aplikasi:
   ```bash
   go run cmd/main.go
   ```

## API Endpoints

### Kategori Buku

- `POST /api/kategori` - Membuat kategori baru
- `PUT /api/kategori/{id}` - Mengupdate kategori
- `DELETE /api/kategori/{id}` - Menghapus kategori
- `GET /api/kategori/{id}` - Mendapatkan kategori berdasarkan ID
- `GET /api/kategori` - Mendapatkan semua kategori
- `GET /api/kategori/search?nama={nama}` - Mencari kategori berdasarkan nama

### Penulis

- `POST /api/penulis` - Membuat penulis baru
- `PUT /api/penulis/{id}` - Mengupdate penulis
- `DELETE /api/penulis/{id}` - Menghapus penulis
- `GET /api/penulis/{id}` - Mendapatkan penulis berdasarkan ID
- `GET /api/penulis` - Mendapatkan semua penulis
- `GET /api/penulis/search?nama={nama}` - Mencari penulis berdasarkan nama

### Penerbit

- `POST /api/penerbit` - Membuat penerbit baru
- `PUT /api/penerbit/{id}` - Mengupdate penerbit
- `DELETE /api/penerbit/{id}` - Menghapus penerbit
- `GET /api/penerbit/{id}` - Mendapatkan penerbit berdasarkan ID
- `GET /api/penerbit` - Mendapatkan semua penerbit
- `GET /api/penerbit/search?nama={nama}` - Mencari penerbit berdasarkan nama

### Petugas

- `POST /api/petugas` - Membuat petugas baru
- `PUT /api/petugas/{id}` - Mengupdate petugas
- `DELETE /api/petugas/{id}` - Menghapus petugas
- `GET /api/petugas/{id}` - Mendapatkan petugas berdasarkan ID
- `GET /api/petugas` - Mendapatkan semua petugas
- `GET /api/petugas/search?username={username}` - Mencari petugas berdasarkan username

## Struktur Proyek

```
.
├── cmd/
│   └── main.go
├── internal/
│   ├── entity/
│   │   ├── kategori_buku.go
│   │   ├── penulis.go
│   │   ├── penerbit.go
│   │   └── petugas.go
│   ├── repository/
│   │   ├── kategori_buku_repository.go
│   │   ├── penulis_repository.go
│   │   ├── penerbit_repository.go
│   │   └── petugas_repository.go
│   ├── usecase/
│   │   ├── kategori_buku_usecase.go
│   │   ├── penulis_usecase.go
│   │   ├── penerbit_usecase.go
│   │   └── petugas_usecase.go
│   └── interface/
│       ├── controllers/
│       │   ├── kategori_buku_controller.go
│       │   ├── penulis_controller.go
│       │   ├── penerbit_controller.go
│       │   └── petugas_controller.go
│       └── router/
│           └── router.go
├── go.mod
└── README.md
```

## Lisensi

MIT 