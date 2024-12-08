# system-booking-hotels

## System Booking Hotel

System Booking Hotel adalah sistem backend berbasis arsitektur microservices yang dirancang untuk mengelola proses pemesanan hotel secara efisien. Sistem ini menggunakan orchestrator sebagai pengelola utama data dan gateway, sehingga memudahkan frontend dalam mengakses setiap endpoint yang tersedia. Proyek ini dibangun menggunakan teknologi modern untuk memastikan performa yang tinggi, keamanan, dan skalabilitas.

---

## Tech Stack yang Digunakan

### Golang

Bahasa pemrograman yang cepat, efisien, dan cocok untuk membangun aplikasi dengan arsitektur microservices.

### Fiber Framework

Framework web yang ringan dan cepat untuk Golang, menawarkan performa tinggi dan pengalaman pengembangan yang mudah.

### Docker dan Docker Compose

- **Docker:** Platform untuk mengelola aplikasi dalam container, memastikan konsistensi lingkungan pengembangan dan produksi.
- **Docker Compose:** Alat untuk mendefinisikan dan menjalankan aplikasi multi-container secara sederhana.

---

## Package dan Library yang Digunakan

- **`github.com/joho/godotenv`:** Membantu dalam membaca file `.env` untuk pengelolaan konfigurasi aplikasi berbasis lingkungan.
- **`github.com/gofiber/fiber/v2`:** Framework utama untuk membangun aplikasi web dengan Fiber.
- **`github.com/gofiber/fiber/v2/middleware/proxy`:** Middleware untuk mengelola proxy request, biasanya digunakan dalam orchestrator.
- **`github.com/stripe/stripe-go` dan `github.com/stripe/stripe-go/paymentintent`:** Library untuk integrasi pembayaran menggunakan Stripe.
- **`gorm.io/driver/postgres` dan `gorm.io/gorm`:** ORM (Object-Relational Mapping) untuk mengelola database PostgreSQL.
- **`gorm.io/driver/sqlite`:** Driver untuk menggunakan SQLite sebagai database alternatif.
- **`github.com/lib/pq`:** Library driver PostgreSQL untuk integrasi database.
- **`github.com/golang-jwt/jwt/v5`:** Library untuk menangani autentikasi berbasis JSON Web Token (JWT).
- **`go.uber.org/zap`:** Library logging dengan performa tinggi untuk debugging dan monitoring aplikasi.
- **`golang.org/x/crypto/bcrypt`:** Digunakan untuk enkripsi password demi menjaga keamanan data pengguna.

---
