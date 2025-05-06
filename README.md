# 📚 RESTful API Go - Perpustakaan

Repositori ini merupakan **tugas Ujian Tengah Semester (UTS)** untuk mata kuliah **Integrative Programming and Technology**.  
Aplikasi ini mengimplementasikan **RESTful API** sederhana untuk sistem perpustakaan, dibuat dengan bahasa pemrograman **Golang**.

---

## 🚀 Fitur Utama

- CRUD data buku
- CRUD data pengguna (anggota perpustakaan)
- Peminjaman dan pengembalian buku
- Struktur modular dan clean code

---

## 🧰 Teknologi yang Digunakan

- Golang (Go)
- REST API + JSON

---

## 📥 Cara Clone Repository

```bash
git clone https://github.com/JayadyGultom/RESTful-API-Go---Perpustakaan.git
cd RESTful-API-Go---Perpustakaan
````

---

## ⚙️ Cara Menjalankan Aplikasi

1. **Pastikan Golang sudah terinstal**

   ```bash
   go version
   ```

2. **Install dependensi**

   ```bash
   go mod tidy
   ```

3. **Jalankan aplikasi**

   ```bash
   go run cmd/main.go
   ```

4. Aplikasi akan berjalan di `localhost:{PORT}` sesuai pengaturan pada file `cmd/main.go`.

---

## 📬 Contoh Endpoint (Sesuaikan)

| Method | Endpoint    | Deskripsi               |
| ------ | ----------- | ----------------------- |
| GET    | /books      | Menampilkan semua buku  |
| POST   | /books      | Menambah buku baru      |
| GET    | /books/{id} | Menampilkan detail buku |
| PUT    | /books/{id} | Mengedit data buku      |
| DELETE | /books/{id} | Menghapus buku          |

---

## 👤 Kontributor

* Jayady Gultom
* Agung Dwi Pradipta
* Firman Sulaiman
* Adrian Wahyuda Aditya P
* Ahmad Hardiansyah
