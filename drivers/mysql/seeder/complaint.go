package seeder

import (
	"e-complaint-api/entities"
	"errors"
	"time"

	"gorm.io/gorm"
)

func SeedComplaint(db *gorm.DB) {
	if err := db.First(&entities.Complaint{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		complaints := []entities.Complaint{
			{
				ID:          "C-81j9aK9280",
				UserID:      1,
				CategoryID:  1,
				Description: "Saya ingin melaporkan kondisi toilet yang sangat kotor di Puskesmas Sentra Kesehatan Masyarakat di wilayah saya. Toilet umum di gedung ini sering kali tidak terawat dengan baik, dinding dan lantai kamar mandi penuh dengan noda dan kotoran, serta bau yang tidak sedap sangat mengganggu penggunaannya. Kondisi ini sangat mengkhawatirkan karena kebersihan toilet yang buruk dapat menyebabkan penyebaran penyakit di antara pengunjung dan staf Puskesmas. Saya berharap pihak terkait segera melakukan perbaikan dan meningkatkan standar kebersihan di fasilitas ini agar dapat memberikan pelayanan kesehatan yang lebih baik dan aman bagi semua orang yang datang ke Puskesmas.",
				RegencyID:   "3601",
				Address:     "Jl. Raya Pandeglang No. 123, Desa Cikupa, Kecamatan Labuan, Kabupaten Pandeglang, Provinsi Banten, 42264",
				Status:      "Pending",
				Type:        "public",
				Date:        time.Now(),
				TotalLikes:  3,
			},
			{
				ID:          "C-8kshis9280",
				UserID:      1,
				CategoryID:  2,
				Description: "Saya ingin melaporkan kurangnya fasilitas di SDN 05 Tangerang, yang sangat memprihatinkan. Ruang kelas yang ada saat ini sangat sempit dan tidak cukup untuk menampung jumlah siswa yang terus bertambah. Selain itu, banyak bangku dan meja yang sudah rusak dan belum diganti, sehingga membuat proses belajar mengajar menjadi tidak nyaman. Fasilitas pendukung seperti perpustakaan dan laboratorium juga sangat minim dan kurang terawat. Hal ini sangat mempengaruhi kualitas pendidikan dan kenyamanan belajar siswa. Saya berharap pihak terkait dapat segera memperbaiki dan menambah fasilitas yang dibutuhkan agar proses belajar mengajar dapat berjalan dengan lebih baik.",
				RegencyID:   "3603",
				Address:     "Jl. Raya Serpong No. 456, Desa Lengkong Karya, Kecamatan Serpong Utara, Kabupaten Tangerang, Provinsi Banten, 15310",
				Status:      "Selesai",
				Type:        "private",
				Date:        time.Now(),
				TotalLikes:  2,
			},
			{
				ID:          "C-81jas92581",
				UserID:      2,
				CategoryID:  3,
				Description: "Saya ingin melaporkan pelayanan yang kurang memadai di Kantor Dinas Kependudukan dan Pencatatan Sipil Kota Serang. Antrian untuk pengurusan KTP, KK, dan akta kelahiran sering kali sangat panjang, dengan waktu tunggu yang tidak menentu dan pelayanan yang lambat. Selain itu, fasilitas ruang tunggu sangat minim dan tidak nyaman, tidak tersedia kursi yang memadai untuk jumlah warga yang datang. Banyak warga yang harus menunggu di luar gedung tanpa perlindungan dari panas dan hujan. Saya berharap pihak terkait dapat meningkatkan efisiensi pelayanan dan memperbaiki fasilitas di kantor tersebut agar warga dapat dilayani dengan lebih cepat dan nyaman.",
				RegencyID:   "3673",
				Address:     "Jl. KH. Abdul Hadi No. 89, Kelurahan Lopang, Kecamatan Serang, Kota Serang, Provinsi Banten, 42111.",
				Status:      "Verifikasi",
				Type:        "private",
				Date:        time.Now(),
				TotalLikes:  2,
			},
			{
				ID:          "C-271j9ak280",
				UserID:      3,
				CategoryID:  4,
				Description: "Saya ingin melaporkan kondisi keamanan yang memprihatinkan di beberapa wilayah di Kota Tangerang. Belakangan ini, sering terjadi tindak kejahatan seperti pencurian dan perampokan di sekitar Jalan Pahlawan dan sekitarnya. Lampu penerangan jalan yang kurang memadai menjadi salah satu faktor yang mempermudah pelaku kejahatan beraksi pada malam hari. Selain itu, kepolisian setempat terkesan lamban dalam merespons laporan kejahatan dari masyarakat, sehingga membuat warga merasa tidak aman. Saya berharap pihak berwenang dapat meningkatkan patroli keamanan, memperbaiki penerangan jalan, dan meningkatkan respon terhadap laporan kejahatan untuk meningkatkan keamanan dan ketertiban di Kota Tangerang.",
				RegencyID:   "3671",
				Address:     "Jl. Gatot Subroto No. 234, Kelurahan Karawaci, Kecamatan Karawaci, Kota Tangerang, Provinsi Banten, 15810",
				Status:      "On Progress",
				Type:        "public",
				Date:        time.Now(),
				TotalLikes:  1,
			},
			{
				ID:          "C-123j9ak280",
				UserID:      3,
				CategoryID:  6,
				Description: "Saya ingin melaporkan masalah pencemaran lingkungan yang serius di sekitar kawasan industri di Kota Cilegon. Pabrik-pabrik di sekitar Jalan Raya Merak terus melakukan pembuangan limbah secara ilegal ke sungai yang mengakibatkan sungai tersebut tercemar berat. Air sungai yang sudah tercemar ini sangat berbahaya bagi lingkungan sekitar dan juga kesehatan masyarakat yang menggunakan air dari sungai tersebut. Selain itu, kebisingan dan polusi udara dari aktivitas industri juga telah menciptakan lingkungan yang tidak sehat bagi penduduk sekitar. Saya berharap pemerintah setempat segera mengambil tindakan tegas untuk mengendalikan pembuangan limbah industri, memulihkan kondisi sungai, serta meningkatkan pengawasan terhadap kegiatan industri agar lingkungan Kota Cilegon bisa menjadi lebih bersih dan sehat.",
				RegencyID:   "3672",
				Address:     "Jl. Letnan Jenderal Suprapto No. 78, Kelurahan Cibeber, Kecamatan Cilegon, Kota Cilegon, Provinsi Banten, 42441",
				Status:      "Ditolak",
				Type:        "public",
				Date:        time.Now(),
				TotalLikes:  0,
			},
		}

		if err := db.CreateInBatches(&complaints, len(complaints)).Error; err != nil {
			panic(err)
		}
	}
}
