# EcoVantory: Enterprise-Grade Inventory & CRM System

**EcoVantory** adalah sistem manajemen inventaris dan CRM (Customer Relationship Management) mandiri yang dirancang untuk mengelola siklus hidup barang unik secara presisi[cite: 4]. Proyek ini mengadopsi standar sistem informasi profesional dengan fokus pada efisiensi pelayanan, pengurangan risiko kesalahan pencatatan, dan integritas data melalui rekam medik elektronik (digital record)[cite: 13, 14, 15, 16].

## 🎯 Project Goals
* **Efisiensi Pelayanan**: Mengoptimalkan waktu layanan melalui digitalisasi proses inventaris[cite: 14, 17].
* **Akurasi Data**: Memastikan ketepatan jumlah stok fisik dan sistem melalui mekanisme verifikasi yang ketat[cite: 11, 15].
* **Auditabilitas**: Mencatat setiap perubahan status barang dan riwayat resep/order secara elektronik[cite: 16, 40].
* **Ketepatan Waktu**: Menunjang proses antrian dan penyerahan barang agar tepat waktu[cite: 17, 18].

## 🚀 Key Features

### 1. Robust Stock Ledger & Snapshot System
Sistem ini mengimplementasikan pengelolaan stok farmasi/barang yang mencakup mutasi barang, stok opname, dan penyesuaian stok (*stock adjustment*)[cite: 539, 565, 590].
* **Stock Ledger**: Mencatat setiap pergerakan barang (masuk, keluar, retur) secara mendetail[cite: 539, 561, 562].
* **Balance Snapshot**: Menggunakan kolom `Balance_After` pada setiap mutasi untuk performa pelaporan stok yang instan dan akurat.

### 2. Header-Detail Transaction Architecture
Mengadopsi pola *Normalized Transaction* untuk mencatat setiap order item yang dilakukan dari pelayanan[cite: 37].
* **Historical Accuracy**: Menyimpan harga jual pada saat transaksi terjadi (`Price_At_Sale`) untuk menjaga konsistensi laporan keuangan[cite: 162, 301].
* **Multi-item Support**: Mendukung satu nomor invoice untuk banyak item barang dalam satu sesi pelayanan[cite: 332, 412].

### 3. State Management with Task ID Workflow
Mengelola alur kerja petugas farmasi/inventaris melalui tahapan Task ID yang sistematis[cite: 48, 52]:
* **Task ID 5 (Available)**: Barang telah diverifikasi dan siap untuk didistribusikan[cite: 48, 52].
* **Task ID 6 (Validating)**: Proses pengecekan item resep/order dan penentuan jenis stok[cite: 48, 261].
* **Task ID 7 (Sold/Released)**: Penyerahan barang selesai dan stok diperbarui secara permanen[cite: 48, 52].

### 4. Dynamic Pricing & Margin Configuration
Implementasi konfigurasi harga otomatis berdasarkan parameter bisnis yang fleksibel[cite: 210, 211]:
* **Formula Perhitungan Harga Jual (HJA)**[cite: 238, 239]:
  $$HJA = \frac{HNA}{(100\% - Margin\%)}$$
* **Margin Analysis**: Menghitung margin secara otomatis berdasarkan perbandingan Harga Netto (HNA) dan Harga Jual (HJA)[cite: 243].

### 5. Comprehensive Audit Trail & Soft Delete
* **Audit Logs**: Mencatat riwayat setiap proses dari verifikasi sampai dengan resep/order selesai[cite: 38].
* **Soft Delete**: Menggunakan kolom `Deleted_At` pada seluruh entitas utama (Item, User, Transaction) untuk memastikan data historis tidak pernah hilang sepenuhnya demi keperluan audit[cite: 16, 402].

## 🛠 Tech Stack
* **Language**: Go (Golang) - Untuk performa tinggi dan keamanan tipe data.
* **Database**: SQLite - Manajemen data relasional yang efisien dan portable.
* **UI Engine**: Templ & HTMX - Antarmuka reaktif dan modern dengan pendekatan server-side rendering.
* **Arch**: Layered Architecture (Clean/Hexagonal inspired).

## 📂 Project Structure
```text
.
├── cmd/server/       # Entry point aplikasi
├── internal/
│   ├── models/       # Definisi database schema & structs
│   ├── database/     # Database connection & migration logic
│   ├── repository/   # Data Access Object (DAO) & SQL implementation
│   └── handlers/     # Request handling & business logic orchestration
├── templates/        # UI Components (Templ + CSS)
└── go.mod            # Dependency management