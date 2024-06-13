package seeder

import (
	"e-complaint-api/entities"
	"errors"

	"gorm.io/gorm"
)

func SeedFaq(db *gorm.DB) {
	if err := db.First(&entities.Faq{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		faqs := []entities.Faq{
			{
				Question: "Apa itu KeluhProv?",
				Answer:   "KeluhProv adalah platform digital yang memungkinkan dan mempermudah warga banten untuk menyampaikan keluhan-keluhan mereka dengan tanggapan yang cepat oleh pemerintah setempat terkait Kesehatan, Pendidikan, Kependudukan, Keamanan, Infrastruktur, Lingkungan maupun Transportasi Daerah Banten.",
			},
			{
				Question: "Apa yang harus saya sertakan dalam aduan saya?",
				Answer:   "Saat menyampaikan aduan, harap sertakan informasi berikut: Foto aduan, lokasi kabupaten/kota, detail alamat, kategori aduan, tanggal kejadian, dan deskripsi aduan.",
			},
			{
				Question: "Berapa lama waktu yang dibutuhkan untuk mendapatkan tanggapan?",
				Answer:   "Kami berkomitmen untuk merespons keluhan Anda dalam waktu 1-2 hari pada jam kerja.",
			},
			{
				Question: "Berapa lama waktu yang dibutuhkan untuk aduan mendapatkan verifikasi?",
				Answer:   "Kami berkomitmen untuk merespons keluhan Anda dalam waktu 1-2 hari pada jam kerja.",
			},
			{
				Question: "Berapa lama waktu yang dibutuhkan untuk aduan mulai di proses ? ",
				Answer:   "Kami berkomitmen untuk melakukan proses terhadap keluhan Anda dalam waktu maksimal 1 hari pada jam kerja setelah dilakukannya verifikasi.",
			},
			{
				Question: "Berapa lama waktu yang dibutuhkan untuk proses aduan selesai? ",
				Answer:   "Lama proses aduan selesai tergantung pada volume aduan yang diberikan. Namun, kami berkomitmen untuk menyelesaikan proses aduan dalam waktu 3-5 hari pada jam kerja setelah progress aduan dimulai.",
			},
			{
				Question: "Apa yang akan terjadi setelah saya mengajukan aduan?",
				Answer:   "Setelah kami menerima aduan Anda, langkah-langkah berikut akan diambil: 1. Aduan Anda diverifikasi oleh admin dalam waktu 1-2 hari pada jam kerja setelah aduan selesai dibuat; 2. Setelah terverifikasi, admin akan meneruskan aduan ke pemerintah setempat; 3. Pemerintah setempat akan memulai proses terhadap aduan maksimal 1 hari setelah dilakukannya verifikasi, dan admin akan mengubah status aduan menjadi on progress; 4. Proses penyelesaian aduan akan diselesaikan secepat mungkin; 5. Setelah proses selesai, admin akan mengubah status aduan menjadi selesai.",
			},
			{
				Question: "Bagaimana jika saya tidak puas dengan tanggapan atau resolusi yang diberikan?",
				Answer:   "Jika Anda merasa tidak puas dengan tanggapan atau resolusi yang diberikan, Anda dapat menghubungi tim kami melalui komentar/diskusi di Aduan anda atau anda bisa ke menu chat dengan admin untuk berkomunikasi langsung dengan admin.",
			},
			{
				Question: "Bagaimana cara melacak status keluhan saya? ",
				Answer:   "Anda dapat melacak status keluhan Anda dengan cara menuju ke menu “Aduanku” pada halaman utama aplikasi. Pada halaman aduanku tersebut, semua aduan yang anda buat beserta status aduan anda akan tampil. Jika anda ingin melihat progress aduan anda yang lebih rinci, anda dapat menekan tombol “Lihat Detail” lalu menekan tombol “Progress Aduan”.",
			},
			{
				Question: "Bagaimana cara mengajukan pengaduan? ",
				Answer:   "Untuk melakukan pengajuan, silahkan mengikuti langkah-langkah sebagai berikut: 1. Masuk ke aplikasi dengan akun anda; 2. Pindah ke halaman buat aduan dengan cara menekan icon tambah(+); 3. Isi data aduan pada form yang disediakan; 4. Submit aduan yang sudah diisikan",
			},
			{
				Question: "Apakah ada biaya yang dikenakan jika melakukan pengaduan di KeluhProv?",
				Answer:   "Pengguna aplikasi/website KeluhProv tidak dikenakan biaya. Aplikasi/ website KeluhProv dapat diunduh dan digunakan secara gratis.",
			},
			{
				Question: "Apa yang harus saya lakukan jika lupa kata sandi akun saya? ",
				Answer:   "Jika anda lupa kata sandi, lakukan langkah-langkah berikut ini: 1. Pada halaman login, klik “Forgot Password?”; 2. Masukkan email agar mendapatkan instruksi untuk reset password; 3. Masukkan password yang baru; 4. Masukkan password yang baru lagi untuk konfirmasi",
			},
			{
				Question: "Apakah saya bisa mengajukan aduan dengan tanpa identitas (anonim)?",
				Answer:   "Ya, anda bisa mengajukan pengajuan dengan anonim melalui Aplikasi KeluhProv dengan cara memilih opsi “private” di Jenis Aduan saat membuat aduan.",
			},
			{
				Question: "Bagaimana cara menambahkan komentar atau informasi tambahan pada pengaduan yang sudah dilakukan sebelumnya?",
				Answer:   "Untuk menambahkan komentar atau informasi tambahan pada aduan yang sudah ada, anda bisa melakukan hal-hal berikut ini : 1. Masuk ke halaman riwayat aduan; 2. Pilih/ Cari aduan yang ingin ditambahkan komentar atau informasi tambahan; 3. Klik ikon komentar untuk menambahkan komentar atau informasi tambahan",
			},
			{
				Question: "Mengapa complaint saya belum diproses?",
				Answer:   "Ada beberapa alasan mengapa aduan Anda belum diproses: 1. Aduan Anda belum diverifikasi oleh admin; 2. Aduan Anda belum diteruskan ke pemerintah setempat; 3. Aduan Anda belum direspon oleh pemerintah setempat; Kami akan berusaha secepat mungkin untuk memproses aduan Anda.",
			},
		}

		if err := db.CreateInBatches(&faqs, len(faqs)).Error; err != nil {
			panic(err)
		}
	}
}
