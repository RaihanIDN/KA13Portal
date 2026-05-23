package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type JadwalAkademik struct {
	Kegiatan string
	Tanggal  string
}

type DataPage struct {
	Title    string
	Kalender []JadwalAkademik
}

type Mahasiswa struct {
	NPM  string
	Nama string
}

func main() {
	// =========================================================================
	// SETUP SEEDER DATA SUPABASE AUTH (MASS REGISTER)
	// =========================================================================
	// Database Supabase Users sudah terisi penuh ! 
	// Baris di bawah ini sengaja dikomentari kembali agar tidak mengeksekusi double register 
	// saat dideploy ke Hugging Face Spaces. Hapus tanda '//' jika ingin digunakan kembali di lokal.
	
	// const supabaseURL = 
	// const serviceRoleKey =
	
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	seedMahasiswaToSupabase(supabaseURL, serviceRoleKey)
	// }()
	// =========================================================================

	// Menyediakan file statis (logo, gambar, css, dll) 
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Routing Page Handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/download", downloadHandler)
	http.HandleFunc("/jadwal", jadwalHandler)
	http.HandleFunc("/semester", semesterHandler)

	port := "7860" // Port standar wajib Hugging Face Spaces & Lokal 
	log.Printf("Server berjalan dengan aman di http://localhost:%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Gagal menjalankan server backend Go: ", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	dataKalender := []JadwalAkademik{
		{Kegiatan: "Perkuliahan sebelum UTS.", Tanggal: ""},
		{Kegiatan: "a. Sebelum Hari Raya Idul Fitri", Tanggal: "2 Maret – 14 Maret 2026"},
		{Kegiatan: "b. Setelah Hari Raya Idul Fitri", Tanggal: "30 Maret – 23 Mei 2026"},
		{Kegiatan: "Libur Hari Raya Idul Fitri", Tanggal: "16 Maret – 28 Maret 2026"},
		{Kegiatan: "Pendistribusian FRS ke mahasiswa melalui situs www.baak.gunadarma.ac.id", Tanggal: "2 Maret – 11 April 2026"},
		{Kegiatan: "Kegiatan Pengisian dan Cetak KRS Online (termasuk Batal/Tambah/Ubah)", Tanggal: "4 Maret – 18 April 2026"},
		{Kegiatan: "Batas Akhir Cetak KRS Online", Tanggal: "16 Mei 2026"},
		{Kegiatan: "Ujian Tengah Semester (UTS)", Tanggal: "25 Mei – 13 Juni 2026"},
		{Kegiatan: "Batas akhir pengurusan cuti akademik", Tanggal: "5 Juni 2026"},
		{Kegiatan: "Daftar Ulang Online", Tanggal: "25 Mei – 13 Juni 2026"},
		{Kegiatan: "Kursus/Pelatihan Berbasis Kompetensi untuk Kelas 4 Yang Tidak Ada UTS", Tanggal: "25 Mei – 6 Juni 2026"},
		{Kegiatan: "Kursus/Pelatihan dan Uji Berbasis Kompetensi untuk Kelas 4 Jenjang S1 dan Kelas 3 Jenjang D3", Tanggal: "2 Juni – 13 Juni 2026"},
		{Kegiatan: "Kursus/Pelatihan Berbasis Kompetensi untuk Kelas 1, 2 dan 3 serta Uji Kompetensi untuk Kelas 4 Jenjang S1 dan Kelas 3 Jenjang D3", Tanggal: "15 Juni – 20 Juni 2026"},
		{Kegiatan: "Perkuliahan setelah UTS", Tanggal: "22 Juni – 18 Juli 2026"},
		{Kegiatan: "Ujian Utama", Tanggal: "20 Juli – 25 Juli 2026"},
		{Kegiatan: "Ujian Akhir Semester (UAS)", Tanggal: "27 Juli – 8 Agustus 2026"},
		{Kegiatan: "Uji Kompetensi untuk kelas 4 Jenjang S1 dan Kelas 3 Jenjang D3", Tanggal: "20 Juli – 8 Agustus 2026"},
		{Kegiatan: "Libur Antar Semester", Tanggal: "10 Agustus – 12 September 2026"},
	}

	pageData := DataPage{
		Title:    "Dashboard Kelas 2KA13",
		Kalender: dataKalender,
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error: File index.html tidak ditemukan !", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, pageData)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// ── 🔒 INJEKSI KEY LOGIN DARI BACKEND GO ──
	pageData := map[string]interface{}{
		"SupabaseKey": "sb_publishable_e34PjLfSoqg--E1Y2kdOTw_RXuiIJPS",
	}
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, "Error: File login.html tidak ditemukan !", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, pageData)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	// ── 🔒 INJEKSI KEY DOWNLOAD DARI BACKEND GO ──
	pageData := map[string]interface{}{
		"Title":       "Download Center - 2KA13",
		"SupabaseKey": "sb_publishable_e34PjLfSoqg--E1Y2kdOTw_RXuiIJPS",
	}
	tmpl, err := template.ParseFiles("templates/download.html")
	if err != nil {
		http.Error(w, "Error: File download.html tidak ditemukan !", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, pageData)
}

func jadwalHandler(w http.ResponseWriter, r *http.Request) {
	// ── 🔒 INJEKSI KEY JADWAL DARI BACKEND GO  ──
	pageData := map[string]interface{}{
		"Title":       "Jadwal Kuliah - 2KA13",
		"SupabaseKey": "sb_publishable_e34PjLfSoqg--E1Y2kdOTw_RXuiIJPS",
	}
	tmpl, err := template.ParseFiles("templates/jadwal.html")
	if err != nil {
		http.Error(w, "Error: File jadwal.html tidak ditemukan !", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, pageData)
}

func semesterHandler(w http.ResponseWriter, r *http.Request) {
	pageData := DataPage{
		Title: "Arsip Semester - 2KA13",
	}
	tmpl, err := template.ParseFiles("templates/semester.html")
	if err != nil {
		http.Error(w, "Error: File semester.html tidak ditemukan !", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, pageData)
}

func seedMahasiswaToSupabase(supabaseURL string, serviceRoleKey string) {
	mapMahasiswa := []Mahasiswa{
		{"10124043", "ADRIANO BAGUS WICAKSONO"},
		{"10124045", "ADZRIL TRIBRATA PUTRA"},
		{"10124052", "AGUNG WISHNU WARDANA"},
		{"10124160", "ANGELO GRACIA SANGGU"},
		{"10124217", "AUDRIC NAUFAL WINENDA"},
		{"10124221", "AULIA RAHMA YUNITA PUTRI"},
		{"10124294", "DAFI FATKHUR RAHMAN"},
		{"10124384", "EIDIYYA FENY PERMADI"},
		{"10124396", "ETHANAEL ALEXANDER SEAN KAETER"},
		{"10124424", "FAISAL RAMDHONI"},
		{"10124448", "FARREL ADRIAN SIAHAAN"},
		{"10124451", "FARREL AVATAR"},
		{"10124452", "FARREL DANENDRA RIFQI"},
		{"10124460", "FATH ABDURACHMAN ROBBANI"},
		{"10124476", "FAYZA KULLA AZMINA"},
		{"10124558", "HASBI ALFARIZI"},
		{"10124560", "HELMY FATHURACHMAN"},
		{"11124430", "INOLIKOVIA ROOSHADI"},
		{"10124606", "IRHAMNA BARIQ FIKRILHAQ"},
		{"10124687", "M K FAIZ"},
		{"10124702", "MARCELLINOUS LOUIS BYL"},
		{"10124751", "MORENO NABIEL PUTRATAMA"},
		{"10124774", "MUHAMAD NABIL FAKHRIY"},
		{"10124813", "MUHAMMAD AZFA AZIZ"},
		{"10124886", "MUHAMMAD IMRON ALFATAH"},
		{"10124890", "MUHAMMAD IRFANSYAH"},
		{"10124912", "MUHAMMAD NAJMI ARDIA RYVANSAH"},
		{"10124956", "MUHAMMAD RIFKI ADITYA"},
		{"10124962", "MUHAMMAD RIZAL ZAM ZAM"},
		{"11124442", "MUHAMMAD TEGAR RAMADHAN"},
		{"11124008", "NABIL SYAHFA HIDAYAT"},
		{"11124009", "NABIL SYUKRI"},
		{"11124050", "NAZWA FITRIYANA"},
		{"11124138", "RAIHAN FAISAL YUSUF"},
		{"11124224", "RIFAY PARULIAN SIMBOLON"},
		{"11124366", "WAHYU DWI SUTANTO"},
		{"11124431", "YOHANA ELISABETH"},
	}

	url := fmt.Sprintf("%s/auth/v1/admin/users", supabaseURL)
	log.Println("Memulai proses otomatis seeding akun mahasiswa ke Supabase Cloud...")

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	for _, mhs := range mapMahasiswa {
		email := fmt.Sprintf("%s@ka13.web.id", mhs.NPM)
		password := "ka13hebat"

		userData := map[string]interface{}{
			"email":         email,
			"password":      password,
			"email_confirm": true,
			"user_metadata": map[string]interface{}{
				"nama_lengkap": mhs.Nama,
				"npm":          mhs.NPM,
			},
		}

		jsonData, _ := json.Marshal(userData)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			log.Printf("Gagal membuat data request untuk: %s\n", mhs.Nama)
			continue
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+serviceRoleKey)
		req.Header.Set("apikey", serviceRoleKey)

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Koneksi gagal saat mendaftarkan %s: %v\n", mhs.Nama, err)
			continue
		}

		if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK {
			log.Printf("Akun Mahasiswa Terbuat -> %s (%s)\n", mhs.Nama, mhs.NPM)
		} else {
			log.Printf("Skip/Gagal mendaftarkan %s, Status Code: %d\n", mhs.Nama, resp.StatusCode)
		}
		resp.Body.Close()
		
		time.Sleep(100 * time.Millisecond)
	}
	log.Println("Proses seeding selesai.")
}