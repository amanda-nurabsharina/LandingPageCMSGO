# BE LP GO (Backend Landing Page Go)

Backend untuk Landing Page CMS yang dibangun menggunakan **Go + Fiber + GORM + PostgreSQL**, menggantikan backend PHP Laravel sebelumnya (`BE LP`).

---

## 🚀 Fitur Utama

1. **High Performance REST API**:
   - `GET /api/landing-page` : Mengembalikan seluruh data konfigurasi website, hero section, benefits (why choose us), statistics, services, order steps, portfolios, dan testimonials dalam satu payload terpadu (100% kompatibel dengan frontend React di `FE LP`).
   - CORS diaktifkan otomatis untuk domain frontend.
2. **Static Asset / Storage Serving**:
   - `GET /storage/{path}` : Melayani asset gambar, logo, portfolio yang diunggah (pengganti `php artisan storage:link` Laravel).
3. **Admin CMS Panel**:
   - URL: `http://127.0.0.1:8000/admin`
   - Manajemen lengkap untuk General Settings, Hero Settings, Benefits, Statistics, Services, Order Steps, Portfolios, dan Testimonials.
   - Dilengkapi proteksi autentikasi, session cookie HttpOnly, CSRF token, brute-force rate limiter, dan live image preview.
4. **Database & ORM**:
   - Database: **PostgreSQL** (`suppliyam_lp`)
   - ORM: **GORM** dengan AutoMigrate otomatis.
   - Compatible password hash: Mendukung hash bcrypt `User` bawaan Laravel.
5. **Built-in CLI Utilities**:
   - Migrasi skema database.
   - Seeder otomatis jika tabel kosong.
   - CLI pembuatan admin (`create-admin`).
   - Import data otomatis dari SQLite lama (`import-sqlite`).

---

## 📂 Struktur Direktori

```
BE LP GO/
├── bin/                       # Output binary hasil build (be-lp-go.exe)
├── internal/
│   ├── admin/                 # Panel Admin CMS (pengganti Filament)
│   │   ├── admin.go           # Setup session, routing & render engine
│   │   ├── auth.go            # Login, logout, proteksi middleware & password hashing
│   │   ├── crud.go            # List, search, pagination, bulk delete, create, edit, save
│   │   ├── resources.go       # Metadata konfigurasi resource & skema form/tabel
│   │   └── templates/         # Template HTML (Tailwind CSS + Lucide Icons)
│   ├── config/                # Environment loader (.env)
│   ├── database/              # PostgreSQL connection pool & migrasi GORM
│   ├── handlers/              # Endpoint REST API (/api/landing-page)
│   ├── models/                # Model GORM & cast JSON
│   ├── seed/                  # Seeder default content
│   └── storage/               # File upload & static disk
├── storage/public/            # Direktori file upload (logo, hero, portfolio)
├── .env                       # Konfigurasi aktif PostgreSQL & server
├── .env.example               # Template environment
├── go.mod                     # Dependensi Go
├── import.go                  # Utilitas migrasi data SQLite -> PostgreSQL
└── main.go                    # Entrypoint & routing HTTP Fiber
```

---

## 🛠️ Konfigurasi `.env`

```ini
APP_NAME="General CMS"
APP_ENV=local
APP_DEBUG=true
APP_URL=http://127.0.0.1:8000
APP_PORT=8000
APP_TIMEZONE=Asia/Jakarta

DB_HOST=127.0.0.1
DB_PORT=5432
DB_DATABASE=suppliyam_lp
DB_USERNAME=postgres
DB_PASSWORD=admin12345
DB_SSLMODE=disable
DB_AUTO_MIGRATE=true

CORS_ALLOW_ORIGINS=*
STORAGE_PATH=./storage/public

ADMIN_NAME="Admin"
ADMIN_EMAIL="admin@example.com"
ADMIN_PASSWORD=""
```

---

## 💻 Cara Menjalankan

### 1. Menjalankan Backend (Development)
```bash
cd "c:\Jody\Landing Page Nanda\BE LP GO"
go run .
# atau jalankan binary yang sudah di-compile:
.\bin\be-lp-go.exe
```
Server akan berjalan di `http://127.0.0.1:8000`.

### 2. Menjalankan Frontend (`FE LP`)
```bash
cd "c:\Jody\Landing Page Nanda\FE LP"
npm run dev
```
Frontend otomatis terhubung ke `http://127.0.0.1:8000` via file `.env` yang sudah disiapkan.

---

## 🔑 Akses Panel Admin

- **URL:** [http://127.0.0.1:8000/admin](http://127.0.0.1:8000/admin)
- **Email:** `admin@example.com`
- **Password:** Password admin yang sudah ada di database, atau reset/buat admin baru dengan perintah:
```bash
go run . create-admin -name "Admin" -email admin@example.com -password "admin12345"
```

