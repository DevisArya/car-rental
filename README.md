# Project Name

## Description

Project ini adalah aplikasi backend rent car yang dibuat dengan golang, Echo Framework, GORM, Postgree.

---

## Api Doc

API documentation untuk v1.0.0 [here](https://documenter.getpostman.com/view/20325566/2sAYQanC5D).
API documentation untuk v2.0.0 [here](https://documenter.getpostman.com/view/20325566/2sAYQanC9W).

## Prasyarat

Pastikan telah menginstal berikut:

- [Go](https://golang.org/dl/) versi minimal 1.19
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/)

---

## Persiapan Database

1. Buat database PostgreSQL baru.

   Contoh SQL:

   ```sql
   CREATE DATABASE nama_database;
   ```

2. Pastikan kredensial database sudah disiapkan, seperti username, password, host, port, dll.

---

## Konfigurasi

1. Duplikat file `.env.example` dan ubah namanya menjadi `.env`.
2. Isi file `.env` dengan informasi berikut:

   ```env
   DB_USERNAME=   # Masukkan username PostgreSQL
   DB_PASSWORD=   # Masukkan password PostgreSQL
   DB_PORT=       # Masukkan port PostgreSQL (default: 5432)
   DB_HOST=       # Masukkan host PostgreSQL (contoh: localhost)
   DB_NAME=       # Masukkan nama database
   DB_SSL_MODE=    (contoh: disable)
   DB_TIMEZONE=    (contoh: Asia/Jakarta)
   PORT=8080      # Masukkan port untuk aplikasi
   ```

---

## Langkah Menjalankan Aplikasi

1. Clone repository ini:

   ```bash
   git clone https://github.com/username/repository-name.git
   ```

2. Masuk ke direktori project:

   ```bash
   cd repository-name
   ```

3. Install dependensi:

   ```bash
   go mod tidy
   ```

4. Jalankan migrasi database (jika ada):

   ```bash
   go run main.go migrate
   ```

5. Jalankan aplikasi:

   ```bash
   go run main.go
   ```

6. Akses aplikasi melalui browser atau tool API seperti Postman di URL berikut:
   ```
   http://localhost:8080
   ```
