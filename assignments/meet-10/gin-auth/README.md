# Gin Auth + GORM

Project autentikasi dengan Gin, GORM, Neon PostgreSQL, bcrypt, dan JWT HS256.
Struktur per fitur mengikuti `../basic-orm`: Handler, Service, Repository.
Tabel `users` dibuat otomatis lewat `AutoMigrate` ketika server dimulai.

## Menjalankan

Siapkan Go 1.25+, Git untuk mengambil repository, dan database PostgreSQL di Neon.
Driver PostgreSQL memakai pgx; compiler C tidak diperlukan. Jalankan perintah
berikut dari folder project `starter/meet-10/gin-auth`:

```bash
go mod download
cp .env.example .env
openssl rand -hex 32
```

Di Neon Dashboard, buka project lalu pilih **Connect**. Salin connection string
ke `DATABASE_URL` di `.env`. Pilih host `-pooler` untuk koneksi pooled atau host
langsung untuk koneksi direct. Contoh format, bukan kredensial aktif:

```dotenv
DATABASE_URL='postgresql://USER:PASSWORD@ep-PROJECT-pooler.REGION.aws.neon.tech/neondb?sslmode=require&channel_binding=require'
JWT_SECRET='hasil-openssl-rand-hex-32'
PORT=8080
```

Isi `JWT_SECRET` dengan hasil `openssl rand -hex 32`, minimal 32 byte.
Jangan commit `.env`. Aplikasi otomatis membaca `.env` memakai
`github.com/joho/godotenv` sebelum membaca konfigurasi. Jalankan dari folder project:

```bash
go run .
```

Environment variable yang sudah tersedia di shell tetap diprioritaskan.
File `.env` boleh tidak ada jika konfigurasi sudah disediakan lewat environment;
kesalahan format atau kegagalan membaca file menghentikan startup.

| Variable | Wajib | Keterangan |
| --- | --- | --- |
| `DATABASE_URL` | Ya | URL PostgreSQL dengan TLS; role database perlu izin membuat/mengubah tabel untuk `AutoMigrate` |
| `JWT_SECRET` | Ya | Secret penandatangan JWT, minimal 32 byte |
| `PORT` | Tidak | Port HTTP, default `8080` |
| `TEST_DATABASE_URL` | Hanya test integrasi | URL database terpisah untuk test; tidak dipakai server |

Default server: `http://localhost:8080`. `DATABASE_URL` wajib diisi dan harus memakai
TLS (`sslmode=require`, `verify-ca`, atau `verify-full`). Gunakan `verify-full`
untuk verifikasi sertifikat dan hostname. Password dengan karakter khusus dalam
URL harus di-encode; connection string dari Neon sudah memakai format URL.

Koneksi diperiksa saat startup, dengan timeout default 10 detik. Pool dibatasi
10 koneksi aktif dan 5 idle. `PreferSimpleProtocol` menghindari cache prepared
statement pada koneksi pooler. Aplikasi berhenti jika konfigurasi, koneksi,
atau migrasi gagal; tidak ada fallback ke database lain.

## Endpoint dan middleware

Semua request melewati logger, recovery, serta pembatas body 1 MiB.
Setiap endpoint memiliki hook `MockEndpoint` sendiri di `internal/middleware/mock.go`.

| Method | Endpoint | Header `X-Mock-Middleware` | Auth |
| --- | --- | --- | --- |
| GET | `/health` | `health` | Publik |
| POST | `/api/auth/register` | `auth.register` | Publik |
| POST | `/api/auth/login` | `auth.login` | Publik |
| GET | `/api/me` | `auth.me` | Bearer JWT |
| GET | `/api/public/info` | `public.info` | Publik |
| GET | `/api/private/profile` | `private.profile` | Bearer JWT |

Modul `internal/public/routes.go` menyediakan grup `/api/public` tanpa auth.
`GET /api/public/info` mengembalikan `200` dengan:

```json
{"data":{"message":"API publik dapat diakses tanpa login"}}
```

Modul `internal/private/routes.go` menyediakan grup `/api/private` dengan
`group.Use(requireAuth)`. Semua endpoint dalam grup mewarisi middleware JWT.
`GET /api/private/profile` mengembalikan profil pengguna seperti `/api/me`:
`200` jika berhasil, `401` jika token hilang/salah/kedaluwarsa atau pengguna tidak
ditemukan, `500` jika database gagal dibaca. Request yang ditolak auth tidak
mendapat header hook karena middleware auth berjalan sebelum hook endpoint.

```bash
curl -i http://localhost:8080/api/public/info

# Tanpa token: 401
curl -i http://localhost:8080/api/private/profile

# Isi dari hasil login: 200
TOKEN='<access_token dari login>'
curl -i http://localhost:8080/api/private/profile \
  -H "Authorization: Bearer $TOKEN"
```

Tambahkan endpoint berikutnya pada grup public atau private sesuai kebutuhan auth.

`GET /health` mengembalikan `200` dengan `{"data":{"status":"ok"}}`.
Endpoint ini menunjukkan proses HTTP aktif; tidak menjalankan pemeriksaan database
baru pada tiap request.

| Endpoint | Body JSON / header wajib | Respons berhasil | Respons gagal dari handler/middleware |
| --- | --- | --- | --- |
| `POST /api/auth/register` | `name`, `email`, `password`; `Content-Type: application/json` | `201`, `data` berisi pengguna | `400` input salah, `409` email terdaftar, `500` gagal menyimpan |
| `POST /api/auth/login` | `email`, `password`; `Content-Type: application/json` | `200`, `data.access_token`, `data.token_type`, `data.expires_in` | `400` input salah, `401` kredensial salah, `500` gagal memproses |
| `GET /api/me` | `Authorization: Bearer <access_token>` | `200`, `data` berisi pengguna | `401` token hilang/salah/kedaluwarsa atau pengguna tidak ditemukan, `500` gagal membaca pengguna |

Nama wajib berisi teks, maksimal 100 karakter; spasi di awal/akhir dibuang.
Email wajib valid, maksimal 254 karakter. Password minimal 8 karakter dan
maksimal 72 byte. Body request dibatasi 1 MiB; body JSON terlalu besar ditolak
handler dengan `400`.

Mock hanya menambahkan header untuk menunjukkan middleware berjalan.
Tambahkan logika sebelum/sesudah `c.Next()` atau kondisi berdasarkan `name`
untuk latihan aturan tiap endpoint. Middleware `RequireAuth` memvalidasi JWT
dan mengecek pengguna di database. Mock tidak melewati validasi auth.
Route auth dan middleware terdaftar di `internal/auth/routes.go`;
route health dan perakitan dependency berada di `setupEngine` pada `main.go`.

## Contoh request

```bash
curl -i http://localhost:8080/health

curl -i -X POST http://localhost:8080/api/auth/register \
  -H 'Content-Type: application/json' \
  -d '{"name":"Budi","email":"budi@example.com","password":"password123"}'

curl -i -X POST http://localhost:8080/api/auth/login \
  -H 'Content-Type: application/json' \
  -d '{"email":"budi@example.com","password":"password123"}'
```

Login mengembalikan bentuk berikut; salin `access_token` ke request profil:

```json
{"data":{"access_token":"<JWT>","token_type":"Bearer","expires_in":3600}}
```

```bash
TOKEN='<access_token dari login>'
curl -i http://localhost:8080/api/me -H "Authorization: Bearer $TOKEN"
```

Register mengembalikan `201` dan profil pengguna tanpa password/hash.
Contoh bentuk respons register dan profil (`id` dan waktu mengikuti data tersimpan):

```json
{
  "data": {
    "id": 1,
    "name": "Budi",
    "email": "budi@example.com",
    "created_at": "2026-10-20T10:00:00Z"
  }
}
```

Password minimal 8 karakter dan maksimal 72 byte; email dinormalisasi ke huruf kecil
dan dibatasi unique index. Login/profil berhasil mengembalikan `200`.
Input salah mendapat `400`, email duplikat `409`, kredensial atau token salah `401`,
dan kegagalan database `500`. Format error: `{"error":"pesan"}`.

JWT berlaku satu jam, memerlukan issuer `gin-auth`, dan hanya menerima HS256.
Untuk latihan ini belum ada refresh token atau logout server-side;
tambahkan revocation store saat membutuhkan pencabutan token sebelum kedaluwarsa.

## Verifikasi

```bash
go test ./...
go vet ./...
```

Test tanpa database mencakup validasi URL/TLS, hook endpoint publik,
input tidak valid, serta penolakan token salah/kedaluwarsa.

Untuk test integrasi, buat branch/database Neon khusus test, lalu isi
`TEST_DATABASE_URL` dalam `.env` dengan connection string database tersebut.
`go test` tidak menjalankan startup server, sehingga `.env` perlu dimuat ke shell
untuk test integrasi. Jalankan pada shell Bash/Zsh, dengan file `.env` milik sendiri:

```bash
set -a
source .env
set +a
go test -v ./...
```

Test integrasi menjalankan migrasi tabel `users`, register/login/profil,
hash password, email duplikat, serta pembacaan lewat koneksi PostgreSQL terpisah.
Pengguna test memakai email unik dan dihapus setelah test. Test integrasi dilewati
jika `TEST_DATABASE_URL` kosong; keberhasilan test tanpa URL bukan bukti koneksi Neon.

## Artifact project

| Artifact | Kegunaan |
| --- | --- |
| `README.md` | Panduan setup, kontrak endpoint, contoh curl, dan verifikasi |
| `.env.example` | Template konfigurasi; salin menjadi `.env` lokal |
| `go.mod`, `go.sum` | Dependency dan checksum versi yang digunakan |
| `main.go` | Load `.env`, validasi konfigurasi, migrasi GORM, dan server HTTP |
| `internal/auth/` | Model pengguna, repository GORM, bcrypt/JWT, handler, route, middleware auth |
| `internal/public/routes.go` | Grup API publik tanpa middleware auth |
| `internal/private/routes.go` | Grup API private dengan middleware auth |
| `internal/database/` | Validasi URL/TLS dan pool koneksi PostgreSQL |
| `internal/middleware/mock.go` | Header hook berbeda untuk setiap endpoint |
| `*_test.go` | Unit test dan test integrasi PostgreSQL |

GORM membuat tabel `users` dengan kolom `id`, `name`, `email`, `password_hash`,
dan `created_at`; email memiliki unique index. Password disimpan sebagai hash bcrypt.
Migrasi dijalankan otomatis saat startup; tidak ada perintah migrasi terpisah.
Contoh request tersedia sebagai perintah curl di atas; project belum menyediakan
file OpenAPI/Swagger atau collection Postman.

## Struktur file

```text
gin-auth/
├── main.go              # Perakitan dependency, migrasi, server
├── main_test.go         # Test route dan integrasi PostgreSQL
├── internal/
│   ├── auth/
│   │   ├── model.go
│   │   ├── repository.go
│   │   ├── service.go
│   │   ├── service_test.go
│   │   ├── handler.go
│   │   ├── routes.go
│   │   └── middleware.go
│   ├── database/
│   │   ├── database.go      # Koneksi Neon PostgreSQL
│   │   └── database_test.go
│   ├── middleware/
│   │   └── mock.go         # Hook mock tiap endpoint
│   ├── public/
│   │   └── routes.go       # /api/public/info
│   └── private/
│       └── routes.go       # /api/private/profile, wajib JWT
├── .env.example         # DATABASE_URL, JWT_SECRET, PORT, TEST_DATABASE_URL
├── .gitignore
├── go.mod
└── go.sum
```
