# 📚 RESTful API Perpustakaan – Golang

Repository ini berisi **RESTful API** untuk sistem informasi **Perpustakaan**, yang dibangun menggunakan bahasa pemrograman **Go (Golang)**.

Proyek ini merupakan **tugas kelompok** untuk mata kuliah **Integrative Programming and Technology**. Dalam proyek ini, sistem dikembangkan secara **lintas platform**, yaitu:

- **API Backend**: Golang
- **Web Frontend**: Laravel
- **Mobile App**: Flutter
- **Desktop App**: VB.NET

---

## 🏗️ Struktur Tabel Database

Terdapat 12 tabel utama yang digunakan dalam database perpustakaan ini:

| Tabel                | Fungsi                                                                 |
|----------------------|------------------------------------------------------------------------|
| `anggota`            | Menyimpan data anggota perpustakaan (nama, alamat, dsb).               |
| `buku`               | Menyimpan data buku (judul, penulis, penerbit, stok, dsb).             |
| `denda`              | Mencatat denda karena keterlambatan atau pelanggaran peminjaman.       |
| `detail_peminjaman`  | Rincian buku yang dipinjam dalam satu transaksi.                       |
| `kategori_buku`      | Kategori atau klasifikasi buku (misalnya: fiksi, sejarah, dll).        |
| `log_aktivitas`      | Mencatat aktivitas pengguna dalam sistem (untuk audit trail).          |
| `peminjaman`         | Menyimpan transaksi peminjaman buku oleh anggota.                      |
| `penerbit`           | Informasi tentang penerbit buku.                                       |
| `pengembalian`       | Data pengembalian buku dan penghitungan denda (jika ada).              |
| `penulis`            | Data penulis dari masing-masing buku.                                 |
| `petugas`            | Akun petugas/admin yang mengelola sistem perpustakaan.                 |
| `reservasi`          | Data reservasi buku oleh anggota untuk peminjaman mendatang.           |

---

## 🚀 Cara Menjalankan API

### 1. Clone Repository

```bash
git clone https://github.com/JayadyGultom/RESTful-API-Go---Perpustakaan.git
cd RESTful-API-Go---Perpustakaan
````

### 2. Siapkan Database

* Buat database di MySQL/PostgreSQL (sesuai konfigurasi).
* Jalankan skrip SQL (jika tersedia) untuk membuat semua tabel.

### 3. Konfigurasi

Buat file `.env` atau sesuaikan variabel koneksi langsung dalam kode Anda:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=perpustakaan
PORT=8080
```

### 4. Install Dependencies

```bash
go mod tidy
```

### 5. Jalankan API

```bash
go run cmd/main.go
```

API akan berjalan di: `http://localhost:8080`

---

## 🔗 Contoh Endpoint API

| Method | Endpoint          | Keterangan             |
| ------ | ----------------- | ---------------------- |
| GET    | `/buku`           | Ambil semua data buku  |
| POST   | `/anggota`        | Tambah anggota baru    |
| PUT    | `/peminjaman/:id` | Update data peminjaman |
| DELETE | `/denda/:id`      | Hapus denda tertentu   |

Contoh penggunaan `curl`:

```bash
curl -X POST http://localhost:8080/anggota \
  -H "Content-Type: application/json" \
  -d '{"nama":"Budi","alamat":"Jakarta"}'
```

---

## 📌 Tujuan Proyek

* Mengintegrasikan berbagai platform dan bahasa pemrograman.
* Mengembangkan REST API yang dapat diakses oleh web, mobile, dan desktop.
* Menerapkan arsitektur modular berbasis service dan handler.
* Melatih kerja tim dalam pengembangan perangkat lunak kolaboratif.

---

## 👨‍💻 Anggota Kelompok

Proyek ini dikerjakan oleh mahasiswa dari mata kuliah **Integrative Programming and Technology**:

* Jayady Managam Gultom
* Agung Dwi Pradipta
* Firman Sulaiman
* Adrian Wahyuda Aditya P
* Ahmad Hardiansyah

---

## 📄 Lisensi

MIT © 2025 – Tim Proyek Integrative Programming
