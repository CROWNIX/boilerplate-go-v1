# intern-practice

Repo latihan untuk memahami pattern **Controller → UseCase → QueryBuilder → DB** yang dipakai di seluruh service Tagsamurai.

Kamu akan mengimplementasi dua modul dari nol: `borrowing` (GET list + confirm) dan `missing` (report, list, detail, handle). Semua skema database dan seed data sudah disediakan — kamu hanya menulis kode Go.

Baca `intern-practice-prd.md` untuk spesifikasi lengkap.

---

## Prerequisites

- Go 1.26+ — download di [go.dev/dl](https://go.dev/dl/), pilih versi terbaru
- PostgreSQL — download di [postgresql.org/download](https://www.postgresql.org/download/)
- [`mig`](https://github.com/mystaline/mig) — migration tool

### Install mig

**Linux / macOS**
```bash
git clone https://github.com/mystaline/mig
cd mig
go build -o mig ./cmd/main.go
sudo mv mig /usr/local/bin/
```

**Windows (PowerShell)**
```powershell
git clone https://github.com/mystaline/mig
cd mig
go build -o "$HOME\go\bin\mig.exe" ./cmd/main.go
```

> Pastikan `%USERPROFILE%\go\bin` sudah ada di `PATH`. Biasanya sudah otomatis ditambahkan saat instalasi Go — verifikasi dengan `mig --version` setelah build.

---

## Setup

### 1. Akses private module

Repo ini menggunakan `ts-utils`, library internal Tagsamurai yang hosted di Gitea private. Jalankan sekali di mesin kamu:

```bash
go env -w GOPRIVATE=gitea.qwertysystem.net
git config --global url."https://<username>:<token>@gitea.qwertysystem.net/".insteadOf "https://gitea.qwertysystem.net/"
```

Generate token di Gitea: **Settings → Applications → Generate Token**

### 2. Copy dan isi env

```bash
cp .env.example .env
```

| Key | Keterangan |
|---|---|
| `POSTGRES_USERNAME` | Username PostgreSQL |
| `POSTGRES_PASSWORD` | Password PostgreSQL |
| `DB_URL` | Connection string untuk `mig` (sudah ada default di `.env.example`) |

### 3. Buat database

```bash
# Linux / macOS
createdb intern_practice

# Windows / universal
psql -U postgres -c "CREATE DATABASE intern_practice"
```

### 4. Download dependencies

```bash
go mod tidy
```

### 5. Init migration tracker

> Semua perintah `mig` harus dijalankan dari direktori `intern-practice` (tempat file `.env` berada), karena `mig` membaca `DB_URL` dari environment. Pastikan `.env` sudah diisi sebelum lanjut.

```bash
mig init
```

Membuat tabel `schema_migrations` di database. Wajib dijalankan **sekali** sebelum `mig up` pertama kali.

### 6. Test migrations

```bash
mig test
```

Menjalankan seluruh migration di temporary DB lalu rollback — tidak ada efek ke database utama. Output sukses: `All migrations passed!`

### 7. Jalankan migrations

```bash
mig up
```

### 8. Jalankan server

```bash
go run main.go
```

### 9. Verifikasi

```bash
curl localhost:8080/borrowing/health
curl localhost:8080/missing/health
```

Keduanya harus return `{"status":"ok","domain":"..."}`. Kalau sudah jalan, kamu siap mulai implementasi.

---

## Struktur

```
intern-practice/
├── app/                        ← bootstrap Fiber + middleware
├── internal/
│   ├── borrowing/              ← domain borrowing
│   │   ├── dto/                ← request & response structs
│   │   ├── query_builder/      ← SQL query construction
│   │   ├── usecase/            ← business logic
│   │   └── delivery/http/
│   │       ├── controller/     ← HTTP handlers
│   │       └── route/          ← route registration + DI wiring
│   └── missing/                ← domain missing (struktur sama)
├── migrations/                 ← SQL migration files (sudah disediakan)
├── pkg/
│   └── entity/
│       └── usecase.go          ← interface UseCase[P, R]
├── main.go
└── .env.example
```

File skeleton sudah ada di semua folder — isinya hanya `package xxx`. Tugas kamu mengisi implementasinya.

---

## Migrations

```bash
mig init     # buat tabel schema_migrations (wajib sekali, sebelum pertama kali up)
mig up       # jalankan semua migration
mig down     # rollback satu migration terakhir
mig status   # lihat status migration
mig test     # test semua migration di temporary DB
```
