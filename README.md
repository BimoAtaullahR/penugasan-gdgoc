# Menu Catalog API (GDGOC Hackathon Backend)

RESTful API backend untuk manajemen katalog menu restoran yang modern, skalabel, dan terintegrasi dengan Google Gemini AI. Proyek ini dibangun untuk memenuhi penugasan seleksi GDGOC UGM Divisi Hacker.

**Live Demo:** https://gdgoc-backend-425291130732.us-central1.run.app

## Fitur Utama

1. CRUD Menu Lengkap
   Mendukung operasi Create, Read, Update, dan Delete untuk data menu makanan.

2. Advanced Search & Filter
   Fitur pencarian teks penuh (berdasarkan nama atau deskripsi) serta filter berdasarkan kategori, rentang harga, dan batas kalori.

3. Pagination & Sorting
   Mendukung penanganan data dalam jumlah besar dengan sistem halaman (pagination) dan pengurutan data yang fleksibel.

4. Aggregation
   Fitur pengelompokan menu berdasarkan kategori, baik dalam mode penghitungan jumlah (count) maupun mode daftar (list).

5. AI-Powered Description
   Integrasi dengan Google Gemini AI untuk secara otomatis menghasilkan deskripsi menu yang menarik apabila pengguna mengosongkan kolom deskripsi saat pembuatan menu.

6. Cloud Native
   Aplikasi dikemas menggunakan Docker dan di-deploy menggunakan Google Cloud Run dengan database PostgreSQL (Neon).

## Tech Stack

Proyek ini dibangun menggunakan teknologi berikut:

- Language: Golang 1.25
- Framework: Gin Gonic
- Database: PostgreSQL (via Neon.tech Serverless)
- ORM: GORM
- AI Integration: Google Gemini 2.5 Flash (google.golang.org/genai)
- Deployment: Google Cloud Run

## Cara Menjalankan (Local)

Ikuti langkah-langkah berikut untuk menjalankan proyek di komputer lokal:

1. Clone Repository
   Jalankan perintah git clone diikuti dengan URL repository ini, lalu masuk ke direktori proyek.

2. Setup Environment Variables
   Buat file .env di direktori root dan isi konfigurasi berikut:
   DATABASE_URL="postgres://user:pass@ep-cool.neon.tech/neondb?sslmode=require"
   GEMINI_API_KEY="AIzaSy..."

3. Install Dependencies & Run
   Jalankan perintah berikut di terminal:
   go mod download
   go run main.go

   Server akan berjalan di http://localhost:8080.

## Dokumentasi API Endpoint

Berikut adalah daftar endpoint utama yang tersedia dalam API ini:

POST /menu
Menambahkan menu baru. Jika deskripsi dikosongkan, AI akan otomatis membuatnya.

GET /menu
Mengambil daftar semua menu. Mendukung query parameter ?q=, ?category=, ?min_price=, ?max_price=, ?page=, dan ?per_page=.

GET /menu/:id
Mengambil detail satu menu berdasarkan ID.

PUT /menu/:id
Memperbarui data menu berdasarkan ID (Partial Update).

DELETE /menu/:id
Menghapus menu berdasarkan ID.

GET /menu/group-by-category
Mengambil statistik atau daftar menu yang dikelompokkan per kategori. Gunakan ?mode=count atau ?mode=list.

GET /menu/search
Pencarian menu menggunakan kata kunci pada nama dan deskripsi (?q=keyword).

## Integrasi AI (Gemini)

Fitur ini berfungsi sebagai asisten otomatis untuk pengisian konten. Logika implementasinya adalah sebagai berikut:
1. Sistem menerima request pembuatan menu baru.
2. Sistem memeriksa apakah kolom deskripsi kosong.
3. Jika kosong, sistem memanggil Gemini API dengan prompt spesifik menggunakan nama menu dan bahan-bahan yang diinputkan.
4. Hasil deskripsi dari AI disimpan ke database bersama data menu lainnya.

## Deployment

Aplikasi ini di-deploy menggunakan Google Cloud Run untuk memastikan skalabilitas dan keandalan tinggi.

Command deployment yang digunakan:
gcloud run deploy gdgoc-backend --source . --platform managed --region us-central1 --allow-unauthenticated --max-instances 1 --memory 256Mi --set-env-vars "DATABASE_URL=...,GEMINI_API_KEY=..."