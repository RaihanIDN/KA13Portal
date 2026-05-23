# 🏛️ Portal Kelas KA13 - Universitas Gunadarma
Create by RaihanIDN

Selamat datang di repositori resmi **Portal Kelas KA13** untuk Semester 4 / Kelas ATA 2025/2026. Aplikasi web ini dirancang sebagai pusat informasi akademik terpadu bagi seluruh mahasiswa kelas KA13 Sistem Informasi, Universitas Gunadarma. 

Aplikasi ini dibangun menggunakan arsitektur modern yang menggabungkan performa tangguh backend **Go (Golang)**, fleksibilitas database cloud **Supabase**, serta antarmuka UI yang responsif dan estetik menggunakan **Tailwind CSS** dan **SweetAlert2**.

---
LINK DEMO : 
## 🚀 Fitur Utama Portal

* **📅 Kalender Akademik & Jadwal Kuliah:** Menampilkan agenda kegiatan perkuliahan, rentang waktu krusial, hingga jadwal mingguan yang diperbarui secara dinamis langsung dari database.
* **💾 Pusat Unduhan Materi (Download Area):** Wadah bagi mahasiswa untuk mengunduh berkas materi kuliah, tugas, PPT, maupun berkas pendukung praktikum secara terorganisir.
* **🛠️ Mode Penanggung Jawab (PJ Kelas):** Panel dasbor khusus yang memberikan hak akses penuh kepada PJ Kelas untuk memperbarui konten portal, melakukan manajemen data, serta mengubah latar belakang (*background*) website secara langsung.
* **🎨 Custom Background Canvas:** Fitur interaktif berbasis LocalStorage yang memungkinkan pengguna mengubah estetika visual latar belakang portal dengan file gambar lokal (maksimal 3MB).
* **🔒 Demo Showcase Mode:** Fitur akun khusus (`tamu`) bagi pengunjung luar untuk mencoba seluruh fungsionalitas admin (sandbox) tanpa risiko merusak data asli pada server utama.
* **✨ Modern Dialog UI:** Integrasi penuh komponen SweetAlert2 yang memberikan animasi pop-up super halus, menggantikan `alert()` bawaan browser yang kaku saat melakukan aksi krusial seperti *Logout* atau *Reset Data*.

Masuk ke Demo akun dengan User : tamu dengan Password : tamu 
untuk mencoba fitur full website
---

## 🛠️ Stack Teknologi Yang Digunakan

| Komponen | Teknologi | Deskripsi |
| :--- | :--- | :--- |
| **Backend** | Go (Golang) v1.20+ | Berperan sebagai core web server statis dan router penanganan logika request data. |
| **Database & Auth** | Supabase | Solusi Backend-as-a-Service (BaaS) berbasis PostgreSQL untuk menyimpan relasi data kelas dan autentikasi session user. |
| **Frontend Styling** | Tailwind CSS (CDN) | Framework utility-first untuk menyusun UI modern, bersih, dan sepenuhnya responsif di perangkat mobile maupun desktop. |
| **Interaktivitas UI** | Vanilla JavaScript & SweetAlert2 | Menangani manipulasi DOM dinamis, manajemen state lokal browser, serta dialog pop-up yang interaktif. |
| **Hosting Platform** | Hugging Face Spaces | Infrastruktur cloud gratis dengan SDK Static HTML untuk deployment instan yang andal. |
---

## 📂 Struktur Direktori Proyek

```text
├── .github/workflows/    # [Opsional] Pipeline CI/CD untuk otomatisasi sync ke Hugging Face
├── static/               # Aset gambar statis, favicon, dan logo kelas
│   └── logo.png          # Logo utama identitas kelas 2KA13
├── templates/            # Berkas views/frontend HTML berbasis Go template
│   ├── index.html        # Halaman utama (Kalender Akademik)
│   ├── jadwal.html       # Halaman informasi jadwal perkuliahan
│   ├── download.html     # Halaman repositori berkas unduhan materi
│   └── semester.html     # Halaman rangkuman mata kuliah per-semester
├── .env                  # File konfigurasi rahasia (URL Supabase & API Key) - [LOKAL ONLY]
├── .gitignore            # Daftar file/folder terproteksi yang tidak boleh di-upload ke GitHub
├── go.mod                # Berkas manajemen modul dependensi proyek Go
├── main.go               # Entry point utama web server aplikasi Go
└── README.md             # Dokumentasi panduan teknis proyek (Berkas ini)
```
🔐 Tingkat Hak Akses Pengguna (Role System)
Aplikasi ini membaca data sesi yang tersimpan di localStorage browser untuk menentukan fungsionalitas UI:

PJ KELAS (user_role: "pj"): Memiliki kontrol penuh atas visualisasi situs, modifikasi data, serta panel admin khusus ganti latar belakang.

DEMO MODE (user_role: "tamu"): Ditujukan untuk tujuan pameran/showcase. Dapat memanipulasi UI secara lokal dan mendapatkan notifikasi simulasi sandbox terproteksi.

MAHASISWA (user_role: "mhs"): Pengguna reguler yang dapat melihat seluruh kalender akademik, jadwal, dan mengunduh berkas materi perkuliahan tanpa hak akses modifikasi.


📝 Kontribusi & Lisensi
Proyek ini dikembangkan secara mandiri dan terbuka untuk seluruh mahasiswa 2KA13 yang ingin berkontribusi memperbaiki bug atau menambahkan fitur baru melalui sistem Pull Request (PR) di GitHub.

Hak Cipta © 2026 Portal KA13. Didukung penuh oleh Go & Supabase.
Creator & Maintainer: @Raihan.idn
