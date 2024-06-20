package seeder

import (
	"e-complaint-api/entities"
	"errors"

	"gorm.io/gorm"
)

func SeedNews(db *gorm.DB) {
	if err := db.First(&entities.News{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		news := []entities.News{
			{
				AdminID:    2,
				CategoryID: 1,
				Title:      "Perbaikan Fasilitas Kesehatan di RSUD Tangerang Tanggapi Kritik Masyarakat",
				Content:    "TANGERANG - Rumah Sakit Umum Daerah (RSUD) Tangerang melakukan langkah signifikan dalam meningkatkan kualitas pelayanan kesehatan dengan memperbaiki fasilitas yang ada. Langkah ini diambil sebagai tanggapan terhadap kritik dari masyarakat terkait kondisi fasilitas yang kurang memadai sebelumnya. \n\nManajer RSUD Tangerang, dr. Kartika Pratiwi, menjelaskan bahwa perbaikan tersebut meliputi renovasi ruang perawatan, peningkatan ketersediaan peralatan medis, dan penambahan tenaga medis untuk memastikan pelayanan yang lebih baik bagi pasien. \"Kami mendengar keluhan masyarakat dan kami berkomitmen untuk memberikan pelayanan yang lebih baik lagi. Renovasi ini adalah langkah awal kami untuk memenuhi standar kesehatan yang lebih tinggi,\" ujar dr. Kartika. \n\nSalah seorang pasien, Bapak Hadi (45), mengapresiasi upaya RSUD Tangerang dalam memperbaiki fasilitas. \"Saya melihat perubahan yang positif di sini. Ruang perawatan lebih bersih dan nyaman, serta pelayanan dari tenaga medis lebih baik,\" katanya. \n\nRSUD Tangerang juga mengundang masukan dari masyarakat untuk terus meningkatkan kualitas pelayanan. \"Kami terbuka untuk kritik dan saran dari masyarakat. Bersama-sama, kami berupaya memberikan pelayanan kesehatan terbaik untuk semua pasien,\" tambah dr. Kartika. \n\nPerbaikan fasilitas di RSUD Tangerang ini diharapkan dapat memberikan dampak positif bagi pelayanan kesehatan di wilayah tersebut dan menjadi contoh untuk institusi kesehatan lainnya dalam memenuhi kebutuhan masyarakat akan layanan kesehatan yang berkualitas.",
				TotalLikes: 3,
			},
			{
				AdminID:    3,
				CategoryID: 2,
				Title:      "Perbaikan Fasilitas Pendidikan di SMA Negeri 1 Cilegon Menjawab Aspirasi Masyarakat",
				Content:    "CILEGON - SMA Negeri 1 Cilegon mengumumkan langkah signifikan dalam memperbaiki fasilitas pendidikan demi meningkatkan kualitas belajar mengajar sesuai dengan harapan masyarakat. Langkah ini diambil sebagai respons terhadap kritik dari orang tua siswa dan komunitas pendidikan terkait kondisi fasilitas yang kurang memadai sebelumnya.\n\nKepala Sekolah SMA Negeri 1 Cilegon, Bapak Budi Santoso, menjelaskan bahwa perbaikan tersebut meliputi renovasi ruang kelas, perpustakaan, dan laboratorium, serta peningkatan sarana olahraga untuk mendukung aktivitas siswa di bidang pendidikan dan non-akademik. \"Kami mendengar masukan dari masyarakat dan kami berkomitmen untuk memberikan lingkungan belajar yang lebih baik. Renovasi ini adalah langkah awal kami untuk menciptakan lingkungan pembelajaran yang kondusif dan modern,\" ujar Bapak Budi.\n\nSeorang siswa, Ani (17), mengapresiasi upaya SMA Negeri 1 Cilegon dalam memperbaiki fasilitas. \"Saya merasa senang dengan perubahan ini. Ruang kelas sekarang lebih nyaman dan perpustakaan lebih lengkap,\" katanya.\n\nSMA Negeri 1 Cilegon juga mengundang masukan dari siswa dan orang tua untuk terus meningkatkan kualitas pendidikan. \"Kami terbuka untuk kritik dan saran dari masyarakat pendidikan. Bersama-sama, kami berupaya memberikan pendidikan terbaik untuk semua siswa,\" tambah Bapak Budi.\n\nPerbaikan fasilitas di SMA Negeri 1 Cilegon ini diharapkan dapat memberikan dampak positif bagi proses pendidikan di sekolah tersebut dan menjadi contoh bagi lembaga pendidikan lainnya dalam memenuhi kebutuhan masyarakat akan pendidikan yang berkualitas.",
				TotalLikes: 2,
			},
			{
				AdminID:    2,
				CategoryID: 4,
				Title:      "Upaya Peningkatan Keamanan di Kota Tangerang Mendapat Dukungan Luas",
				Content:    "TANGERANG - Pemerintah Kota Tangerang mengumumkan langkah signifikan dalam meningkatkan keamanan di berbagai wilayah, sebagai respons terhadap meningkatnya kekhawatiran masyarakat akan tingkat kejahatan. Langkah ini diambil setelah serangkaian konsultasi dengan warga dan peningkatan patroli polisi di titik-titik rawan.\n\nWali Kota Tangerang, Ibu Siti Nurhayati, menjelaskan bahwa upaya peningkatan keamanan meliputi penambahan jumlah polisi patroli, pemasangan CCTV di titik strategis, dan peningkatan koordinasi antara kepolisian dengan masyarakat. \"Kami mendengar aspirasi warga dan kami bertekad untuk menciptakan lingkungan yang lebih aman. Langkah-langkah ini adalah langkah awal kami untuk meningkatkan rasa aman dan ketertiban,\" ujar Ibu Siti.\n\nSeorang warga, Bapak Adi (40), menyambut baik upaya pemerintah dalam meningkatkan keamanan. \"Saya merasa lebih aman dengan adanya peningkatan patroli dan pemasangan CCTV. Semoga ini dapat membuat lingkungan kami lebih nyaman,\" katanya.\n\nPemerintah Kota Tangerang juga mengundang masukan dari warga untuk terus meningkatkan keamanan. \"Kami terbuka untuk kritik dan saran dari masyarakat. Bersama-sama, kami berupaya menciptakan kota yang lebih aman untuk semua warga,\" tambah Ibu Siti.\n\nPeningkatan keamanan di Kota Tangerang ini diharapkan dapat memberikan dampak positif bagi tingkat kepercayaan masyarakat dan menjadi contoh bagi kota-kota lain dalam memenuhi kebutuhan akan keamanan yang terjamin.",
				TotalLikes: 3,
			},
		}

		if err := db.CreateInBatches(news, len(news)).Error; err != nil {
			panic(err)
		}
	}
}
