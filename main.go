package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Customer struct {
	Name string
}

type SparePart struct {
	Name     string
	Category string
	Price    int
}

type Transaction struct {
	Customer  Customer
	SparePart SparePart
}

type HistoryFreq struct {
	SparePart SparePart
	Frequency int
}

type HistoryCust struct {
	Customer  Customer
	SparePart SparePart
	Date      time.Time
}

func clear() {
	fmt.Print("\033[H\033[J")
}

func tampilTransaction(transactions [100]Transaction, count int) {
	if count == 0 {
		fmt.Println(strings.Repeat("=", 80))
		fmt.Println("Tidak ada transaksi.")
		return
	}
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("Daftar Transaksi:")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("%-5s %-23s %-23s %-20s\n", "No", "Nama Pelanggan", "Nama Spare Part", "Harga")
	fmt.Println(strings.Repeat("=", 80))
	for i := 0; i < count; i++ {
		fmt.Printf("%-5d %-23s %-23s Rp%-15d\n", i+1, transactions[i].Customer.Name, transactions[i].SparePart.Name, transactions[i].SparePart.Price)
	}
	fmt.Println(strings.Repeat("=", 80))
}

// -------------------------------------------------------------------------------------------------------------------------------------------------------
// Point A
func showAllSpareParts(spareParts [100]SparePart, count int) {
	fmt.Println(strings.Repeat("=", 80))
	fmt.Print(strings.Repeat("=", 16))
	fmt.Print("================ Data Spare Part ===============")
	fmt.Println(strings.Repeat("=", 16))
	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("%-7s %-27s %-17s %-12s\n", "No", "Nama Spare Part", "Kategori", "Harga")

	for i := 0; i < count; i++ {
		fmt.Printf("%-7d %-27s %-17s Rp%-12d\n", i+1, spareParts[i].Name, spareParts[i].Category, spareParts[i].Price)
	}
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println()

}

func kelolaSparePart(spareParts *[100]SparePart, countSpareParts *int, transactions *[100]Transaction, countTransactions *int) {
	var keluar bool

	for !keluar {
		clear()
		showAllSpareParts(*spareParts, *countSpareParts)
		fmt.Println(strings.Repeat("=", 80))
		fmt.Print(strings.Repeat("=", 12))
		fmt.Print("================ Menu Kelola Spare-part ================")
		fmt.Println(strings.Repeat("=", 12))
		fmt.Println(strings.Repeat("=", 80))
		fmt.Println("1. Tambah Spare-part ke Transaksi")
		fmt.Println("2. Edit Spare-part di Transaksi")
		fmt.Println("3. Hapus Spare-part di Transaksi")
		fmt.Println("4. Kembali Ke Menu Utama")
		fmt.Println(strings.Repeat("=", 80))

		var input int
		fmt.Print("Pilihan: ")
		fmt.Scanln(&input)
		fmt.Println(strings.Repeat("=", 80))
		switch input {
		case 1:
			reader := bufio.NewReader(os.Stdin)

			// Input nama pelanggan dengan spasi
			fmt.Print("Masukkan nama pelanggan: ")
			customerName, _ := reader.ReadString('\n')
			customerName = strings.TrimSpace(customerName) // Menghapus newline
			customer := Customer{Name: customerName}
			fmt.Println(strings.Repeat("=", 80))

			// Input nomor spare part (menggunakan fmt.Scanln seperti sebelumnya)
			fmt.Print("Masukkan nomor spare part yang ingin ditambahkan ke transaksi: ")
			var partNumber int
			fmt.Scanln(&partNumber)
			fmt.Println(strings.Repeat("=", 80))

			if partNumber > 0 && partNumber <= *countSpareParts {
				selectedPart := (*spareParts)[partNumber-1]
				transactions[*countTransactions] = Transaction{Customer: customer, SparePart: selectedPart}
				*countTransactions++
				fmt.Println(strings.Repeat("=", 80))
				fmt.Printf("%s berhasil ditambahkan ke transaksi atas nama %s.\n", selectedPart.Name, customerName)
				fmt.Println(strings.Repeat("=", 80))
				fmt.Print("Tekan Enter Untuk melanjutkan ....")
				fmt.Scanln()
				fmt.Println()
			} else {
				fmt.Println("Nomor tidak valid.")
			}

		case 2:
			clear()
			tampilTransaction(*transactions, *countTransactions)
			fmt.Print("Masukkan nomor transaksi yang ingin diubah: ")
			var transactionNumber int
			fmt.Scanln(&transactionNumber)
			if transactionNumber > 0 && transactionNumber <= *countTransactions {
				showAllSpareParts(*spareParts, *countSpareParts)
				fmt.Println(strings.Repeat("=", 80))
				fmt.Print("Masukkan nomor spare part baru: ")
				var partNumber int
				fmt.Scanln(&partNumber)
				if partNumber > 0 && partNumber <= *countSpareParts {
					transactions[transactionNumber-1].SparePart = (*spareParts)[partNumber-1]
					fmt.Println("Spare part dalam transaksi berhasil diubah.")
				} else {
					fmt.Println("Nomor spare part tidak valid.")
				}
			} else {
				fmt.Println("Nomor transaksi tidak valid.")
			}

		case 3:
			clear()
			tampilTransaction(*transactions, *countTransactions)
			fmt.Print("Masukkan nomor transaksi yang ingin dihapus: ")
			var transactionNumber int
			fmt.Scanln(&transactionNumber)
			if transactionNumber > 0 && transactionNumber <= *countTransactions {
				for i := transactionNumber - 1; i < *countTransactions-1; i++ {
					transactions[i] = transactions[i+1]
				}
				*countTransactions--
				fmt.Println("Transaksi berhasil dihapus.")
			} else {
				fmt.Println("Nomor transaksi tidak valid.")
			}

		case 4:
			keluar = true
			fmt.Println("Kembali ke Menu Utama")

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func updateHistoryFreq(history *[100]HistoryFreq, countHistory *int, sparePart SparePart) {
	for i := 0; i < *countHistory; i++ {
		if history[i].SparePart.Name == sparePart.Name {
			history[i].Frequency++
			return
		}
	}
	if *countHistory < 100 {
		history[*countHistory] = HistoryFreq{SparePart: sparePart, Frequency: 1}
		*countHistory++
	}
}

func updateHistPelanggan(history *[100]HistoryCust, historyCount *int, customer Customer, sparePart SparePart) {
	// Cari apakah data pelanggan dan spare part sudah ada di history
	for i := 0; i < *historyCount; i++ {
		if history[i].Customer.Name == customer.Name && history[i].SparePart.Name == sparePart.Name {
			// Jika sudah ada, perbarui tanggal pembelian
			history[i].Date = time.Now()
			return
		}
	}

	// Jika tidak ditemukan, tambahkan data baru ke history
	if *historyCount < 100 {
		history[*historyCount] = HistoryCust{
			Customer:  customer,
			SparePart: sparePart,
			Date:      time.Now(),
		}
		*historyCount++ // Increment jumlah history
	} else {
		fmt.Println("Riwayat pelanggan penuh. Tidak dapat menambahkan data baru.")
	}
}

func processTransaction(transactions *[100]Transaction, countTransactions *int, historyFreq *[100]HistoryFreq, historyFreqCount *int, historyPlgn *[100]HistoryCust, historyPlgnCount *int) {
	clear()
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("========================= Total Transaksi per Pelanggan ========================")
	fmt.Println(strings.Repeat("=", 80))
	totalPerPelanggan := make(map[string]int)
	servicePerPelanggan := make(map[string]int)

	for i := 0; i < *countTransactions; i++ {
		transaction := transactions[i]
		serviceCharge := int(float64(transaction.SparePart.Price) * 0.2)
		totalPerPelanggan[transaction.Customer.Name] += transaction.SparePart.Price
		servicePerPelanggan[transaction.Customer.Name] += serviceCharge
	}

	fmt.Printf("%-20s %-15s %-15s %-15s\n", "Nama Pelanggan", "Total Barang", "Harga Service", "Total Keseluruhan")
	fmt.Println(strings.Repeat("=", 80))

	for name := range totalPerPelanggan {
		totalBarang := totalPerPelanggan[name]
		hargaService := servicePerPelanggan[name]
		totalKeseluruhan := totalBarang + hargaService
		fmt.Printf("%-20s Rp%-14d Rp%-14d Rp%-14d\n", name, totalBarang, hargaService, totalKeseluruhan)
	}

	fmt.Println(strings.Repeat("=", 80))
	reader := bufio.NewReader(os.Stdin)

	// Input nama pelanggan dengan spasi
	fmt.Println("Pilih pelanggan untuk pembayaran:")
	fmt.Print("Nama Pelanggan: ")
	customerName, _ := reader.ReadString('\n')
	customerName = strings.TrimSpace(customerName) // Menghapus newline
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println()
	totalBarang, barangExists := totalPerPelanggan[customerName]
	hargaService, serviceExists := servicePerPelanggan[customerName]

	if barangExists && serviceExists {
		totalKeseluruhan := totalBarang + hargaService

		fmt.Println(strings.Repeat("=", 80))
		fmt.Printf("Detail Pembayaran untuk %s:\n", customerName)
		fmt.Println(strings.Repeat("=", 80))
		fmt.Printf("Total Barang    : Rp%d\n", totalBarang)
		fmt.Printf("Harga Service   : Rp%d\n", hargaService)
		fmt.Printf("Total Keseluruhan: Rp%d\n", totalKeseluruhan)
		fmt.Println(strings.Repeat("=", 80))
		fmt.Println("1. Lanjutkan Pembayaran")
		fmt.Println("2. Batalkan Pembayaran")
		fmt.Println(strings.Repeat("=", 80))
		fmt.Print("Pilihan: ")
		var input int
		fmt.Scanln(&input)

		if input == 1 {
			fmt.Println(strings.Repeat("=", 80))
			fmt.Printf("Pembayaran atas nama %s berhasil!\n", customerName)

			// Iterasi transaksi dan hapus sesuai pelanggan
			for i := 0; i < *countTransactions; {
				if transactions[i].Customer.Name == customerName {
					// Perbarui history frekuensi spare part
					updateHistoryFreq(historyFreq, historyFreqCount, transactions[i].SparePart)

					// Perbarui history pelanggan
					// Tambahkan ini di dalam loop penghapusan transaksi saat pembayaran berhasil:
					updateHistPelanggan(historyPlgn, historyPlgnCount, transactions[i].Customer, transactions[i].SparePart)

					// Hapus transaksi dengan menggeser elemen
					for j := i; j < *countTransactions-1; j++ {
						transactions[j] = transactions[j+1]
					}
					*countTransactions-- // Kurangi jumlah transaksi
				} else {
					i++ // Increment jika tidak menghapus transaksi
				}
			}
			loading()
		} else {
			fmt.Println(strings.Repeat("=", 80))
			fmt.Println("Pembayaran dibatalkan.")
			loading()
		}
	} else {
		fmt.Printf("Tidak ditemukan transaksi atas nama pelanggan '%s'.\n", customerName)
		fmt.Println("Tekan Enter untuk kembali ke menu...")
		fmt.Scanln()
	}
}

func loading() {
	fmt.Println("Memproses...")
	time.Sleep(1 * time.Second)
}

// -------------------------------------------------------------------------------------------------------------------------------------------------------
// Point C
// Fungsi Selection Sort untuk mengurutkan berdasarkan tanggal
func selectionSortByDate(historyplgn *[100]HistoryCust, dataCount int) {
	for i := 0; i < dataCount-1; i++ {
		minIdx := i
		for j := i + 1; j < dataCount; j++ {
			if historyplgn[j].Date.Month() < historyplgn[minIdx].Date.Month() {
				minIdx = j
			}
		}
		historyplgn[i], historyplgn[minIdx] = historyplgn[minIdx], historyplgn[i]
	}
}

// Binary Search berdasarkan bulan
func binarySearchByMonth(historyplgn [100]HistoryCust, dataCount, month int) ([100]HistoryCust, int) {
	low, high := 0, dataCount-1
	var result [100]HistoryCust
	resultCount := 0

	for low <= high {
		mid := (low + high) / 2
		currentMonth := int(historyplgn[mid].Date.Month())

		if currentMonth == month {
			// Tambahkan elemen tengah ke array statis
			result[resultCount] = historyplgn[mid]
			resultCount++

			// Periksa elemen sebelum dengan bulan yang sama
			for i := mid - 1; i >= 0 && int(historyplgn[i].Date.Month()) == month; i-- {
				if resultCount < 100 {
					result[resultCount] = historyplgn[i]
					resultCount++
				}
			}

			// Periksa elemen setelah dengan bulan yang sama
			for i := mid + 1; i < dataCount && int(historyplgn[i].Date.Month()) == month; i++ {
				if resultCount < 100 {
					result[resultCount] = historyplgn[i]
					resultCount++
				}
			}
			break
		} else if currentMonth < month {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return result, resultCount
}

// Sequential Search berdasarkan spare-part
func sequentialSearchBySparePart(historyplgn [100]HistoryCust, dataCount int, sparePartName string) []HistoryCust {
	result := []HistoryCust{}
	sparePartName = strings.ToLower(sparePartName)

	for i := 0; i < dataCount; i++ {
		if strings.ToLower(historyplgn[i].SparePart.Name) == sparePartName {
			result = append(result, historyplgn[i])
		}
	}

	return result
}

// Fungsi untuk menampilkan daftar pelanggan
func daftarPelanggan(historyplgn [100]HistoryCust, dataCount int) {
	if dataCount == 0 {
		fmt.Println("Belum ada transaksi.")
		fmt.Println("Tekan Enter untuk kembali...")
		fmt.Scanln()
		return
	}
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("============================== Daftar Pelanggan ================================")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("1. Cari pelanggan berdasarkan bulan transaksi")
	fmt.Println("2. Cari pelanggan berdasarkan spare-part tertentu")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Println(strings.Repeat("=", 80))
	var pilihan int
	fmt.Print("Pilihan: ")
	fmt.Scanln(&pilihan)
	fmt.Println(strings.Repeat("=", 80))
	switch pilihan {
	case 1:

		fmt.Print("Masukkan Pencarian Periode bulan tertentu (1-12): ")
		var bulan int
		_, err := fmt.Scanln(&bulan)
		if err != nil || bulan < 1 || bulan > 12 {
			fmt.Println("Bulan tidak valid.")
			return
		}

		// Sort data terlebih dahulu untuk binary search
		selectionSortByDate(&historyplgn, dataCount)

		// Gunakan binary search untuk mencari pelanggan berdasarkan bulan
		resultCount := 0
		result, resultCount := binarySearchByMonth(historyplgn, dataCount, bulan)

		fmt.Println(strings.Repeat("=", 80))
		fmt.Printf("Pelanggan yang melakukan transaksi pada bulan %d:\n", bulan)
		fmt.Println(strings.Repeat("=", 80))
		fmt.Printf("%-5s %-25s %-25s %-20s\n", "No", "Nama Pelanggan", "Nama Spare Part", "Tanggal Pembelian")
		fmt.Println(strings.Repeat("=", 80))

		if resultCount == 0 {
			fmt.Printf("Tidak ada pelanggan yang melakukan transaksi pada bulan %d.\n", bulan)
		} else {
			for i := 0; i < resultCount; i++ {
				record := result[i]
				fmt.Printf("%-5d %-25s %-25s %-20s\n", i+1, record.Customer.Name, record.SparePart.Name, record.Date.Format("2006-01-02 15:04:05"))
			}
		}
		fmt.Println(strings.Repeat("=", 80))

		fmt.Println()

	case 2:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Masukkan nama spare-part: ")
		sparePartName, _ := reader.ReadString('\n')
		sparePartName = strings.TrimSpace(sparePartName) // Menghilangkan newline di akhir input
		fmt.Println(strings.Repeat("=", 80))

		// Gunakan sequential search untuk mencari pelanggan berdasarkan spare-part
		result := [100]HistoryCust{}
		resultCount := 0
		sequentialResult := sequentialSearchBySparePart(historyplgn, dataCount, sparePartName)
		for _, record := range sequentialResult {
			if resultCount < 100 {
				result[resultCount] = record
				resultCount++
			}
		}
		fmt.Printf("Pelanggan yang membeli spare-part '%s':\n", sparePartName)
		fmt.Println(strings.Repeat("=", 80))
		fmt.Printf("%-5s %-25s %-25s %-20s\n", "No", "Nama", "Spare-part", "Tanggal Pembelian")
		fmt.Println(strings.Repeat("=", 80))
		if resultCount == 0 {
			fmt.Printf("Tidak ada pelanggan yang membeli spare-part '%s'.\n", sparePartName)
		} else {
			for i := 0; i < resultCount; i++ {
				record := result[i]
				fmt.Printf("%-5d %-25s %-25s %-20s\n", i+1, record.Customer.Name, record.SparePart.Name, record.Date.Format("2006-01-02 15:04:05"))
			}
		}
		fmt.Println(strings.Repeat("=", 80))

	case 3:
		fmt.Println("Kembali ke Menu Utama.")
		fmt.Println(strings.Repeat("=", 80))
		return

	default:
		fmt.Println("Pilihan tidak valid.")
	}
}
func tampilkanHistoryPelanggan(history []HistoryCust, historyCount int) {
	clear() // Bersihkan layar
	if historyCount == 0 {
		fmt.Println("Belum ada data riwayat pembelian pelanggan.")
		fmt.Println("Tekan Enter untuk kembali ke menu...")
		fmt.Scanln()
		return
	}
	fmt.Println()
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("========================= Riwayat Pembelian Pelanggan ==========================")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("%-5s %-25s %-25s %-20s\n", "No", "Nama Pelanggan", "Nama Spare Part", "Tanggal Pembelian")
	fmt.Println(strings.Repeat("=", 80))
	for i := 0; i < historyCount; i++ {
		fmt.Printf("%-5d %-25s %-25s %-20s\n", i+1, history[i].Customer.Name, history[i].SparePart.Name, history[i].Date.Format("2006-01-02 15:04:05"))
	}
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println()

}

//-------------------------------------------------------------------------------------------------------------------------------------------------------

// Point D
// tampil history frekuensi
func tampilkanHistoryFreq(history [100]HistoryFreq, dataCount int) {
	var pilihan int
	clear()
	if dataCount == 0 {
		fmt.Println("Belum ada data riwayat pembelian.")
		return
	}

	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("=================== Pilihan Urutan Riwayat Spare Part =====================")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("1. Urutkan Frekuensi Ascending (Kecil ke Besar)")
	fmt.Println("2. Urutkan Frekuensi Descending (Besar ke Kecil)")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Print("Pilihan: ")
	fmt.Scanln(&pilihan)
	fmt.Println(strings.Repeat("=", 80))
	var sortedHistory [100]HistoryFreq
	switch pilihan {
	case 1:
		// Urutkan secara ascending
		sortedHistory = insertionSortArray(history, dataCount, true)
	case 2:
		// Urutkan secara descending
		sortedHistory = insertionSortArray(history, dataCount, false)
	default:
		fmt.Println("Pilihan tidak valid, kembali ke menu utama.")
		fmt.Println(strings.Repeat("=", 80))
		return
	}

	// Tampilkan hasil setelah diurutkan
	clear()
	fmt.Println()
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("============================= Riwayat Spare Part ===============================")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("%-5s %-25s %-15s %-10s\n", "No", "Nama Spare Part", "Frekuensi", "Harga")
	fmt.Println(strings.Repeat("=", 80))
	for i := 0; i < dataCount; i++ {
		h := sortedHistory[i]
		fmt.Printf("%-5d %-25s %-15d Rp%-10d\n", i+1, h.SparePart.Name, h.Frequency, h.SparePart.Price)
	}
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println()
}

func insertionSortArray(history [100]HistoryFreq, count int, ascending bool) [100]HistoryFreq {
	// Salin data untuk pengurutan
	var sortedHistory [100]HistoryFreq
	copy(sortedHistory[:], history[:])

	// Implementasi Insertion Sort
	for i := 1; i < count; i++ {
		key := sortedHistory[i]
		j := i - 1

		for j >= 0 && ((ascending && sortedHistory[j].Frequency > key.Frequency) || (!ascending && sortedHistory[j].Frequency < key.Frequency)) {
			sortedHistory[j+1] = sortedHistory[j]
			j--
		}
		sortedHistory[j+1] = key
	}
	return sortedHistory
}

func menuUtama() {
	spareParts := [100]SparePart{
		{"Piston", "Mesin", 500000},
		{"Ring Piston", "Mesin", 300000},
		{"Kampas Kopling", "Mesin", 250000},
		{"Gasket Mesin", "Mesin", 100000},
		{"Bearing Kruk As", "Mesin", 400000},
		{"Aki", "Kelistrikan", 700000},
		{"Busi", "Kelistrikan", 50000},
		{"Lampu Utama", "Kelistrikan", 150000},
		{"Regulator Rectifier", "Kelistrikan", 200000},
		{"Kabel Pengapian", "Kelistrikan", 80000},
		{"Ban Depan", "Kaki-Kaki", 600000},
		{"Ban Belakang", "Kaki-Kaki", 650000},
		{"Rantai Roda", "Kaki-Kaki", 300000},
		{"Gir Depan", "Kaki-Kaki", 150000},
		{"Shockbreaker", "Kaki-Kaki", 500000},
		{"Fairing Depan", "Bodi", 1000000},
		{"Cover Samping", "Bodi", 400000},
		{"Spion", "Bodi", 100000},
		{"Jok", "Bodi", 250000},
		{"Stang Kemudi", "Bodi", 350000},
		{"Oli Mesin", "Pendukung", 75000},
		{"Filter Oli", "Pendukung", 50000},
		{"Filter Udara", "Pendukung", 60000},
		{"Kampas Rem Depan", "Pendukung", 120000},
		{"Kampas Rem Belakang", "Pendukung", 110000},
	}
	var transactions [100]Transaction
	var historyfrq [100]HistoryFreq
	var historyplgn [100]HistoryCust

	historyfrq[0] = HistoryFreq{
		SparePart: SparePart{Name: "Shockbreaker", Price: 5000000},
		Frequency: 3,
	}

	historyfrq[1] = HistoryFreq{
		SparePart: SparePart{Name: "Filter Oil", Price: 50000},
		Frequency: 2,
	}

	historyfrq[2] = HistoryFreq{
		SparePart: SparePart{Name: "Kampas Rem Depan", Price: 120000},
		Frequency: 1,
	}
	historyfrq[3] = HistoryFreq{
		SparePart: SparePart{Name: "Oli Mesin", Price: 75000},
		Frequency: 7,
	}
	historyfrq[4] = HistoryFreq{
		SparePart: SparePart{Name: "Piston", Price: 500000},
		Frequency: 10,
	}
	historyfrq[5] = HistoryFreq{
		SparePart: SparePart{Name: "Filter Udara", Price: 60000},
		Frequency: 8,
	}
	historyfrq[6] = HistoryFreq{
		SparePart: SparePart{Name: "Spion", Price: 100000},
		Frequency: 15,
	}
	historyfrq[7] = HistoryFreq{
		SparePart: SparePart{Name: "Jok", Price: 250000},
		Frequency: 19,
	}
	historyfrq[8] = HistoryFreq{
		SparePart: SparePart{Name: "Gir Depan", Price: 150000},
		Frequency: 6,
	}
	historyfrq[9] = HistoryFreq{
		SparePart: SparePart{Name: "Ring Piston", Price: 300000},
		Frequency: 3,
	}
	historyfrq[10] = HistoryFreq{
		SparePart: SparePart{Name: "Busi", Price: 50000},
		Frequency: 8,
	}
	historyfrq[11] = HistoryFreq{
		SparePart: SparePart{Name: "Aki", Price: 700000},
		Frequency: 7,
	}

	historyplgn[0] = HistoryCust{
		Customer:  Customer{Name: "John Doe"},
		SparePart: SparePart{Name: "Aki"},
		Date:      time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC),
	}

	historyplgn[1] = HistoryCust{
		Customer:  Customer{Name: "Aji Noto"},
		SparePart: SparePart{Name: "Jok"},
		Date:      time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC),
	}

	historyplgn[2] = HistoryCust{
		Customer:  Customer{Name: "Alice Brown"},
		SparePart: SparePart{Name: "Busi"},
		Date:      time.Date(2024, 4, 10, 0, 0, 0, 0, time.UTC),
	}
	historyplgn[3] = HistoryCust{
		Customer:  Customer{Name: "Ghilbran raka"},
		SparePart: SparePart{Name: "Spion"},
		Date:      time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC),
	}

	historyplgn[4] = HistoryCust{
		Customer:  Customer{Name: "Jane Smith"},
		SparePart: SparePart{Name: "Ban Depan"},
		Date:      time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC),
	}

	historyplgn[5] = HistoryCust{
		Customer:  Customer{Name: "Afra Lintang"},
		SparePart: SparePart{Name: "Ban belakang"},
		Date:      time.Date(2024, 4, 10, 0, 0, 0, 0, time.UTC),
	}
	historyplgn[6] = HistoryCust{
		Customer:  Customer{Name: "Wafiq Nur"},
		SparePart: SparePart{Name: "Piston"},
		Date:      time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC),
	}

	historyplgn[7] = HistoryCust{
		Customer:  Customer{Name: "Jiaa"},
		SparePart: SparePart{Name: "Filter Oli"},
		Date:      time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC),
	}

	historyplgn[8] = HistoryCust{
		Customer:  Customer{Name: "Udin Brown"},
		SparePart: SparePart{Name: "Filter Udara"},
		Date:      time.Date(2024, 4, 10, 0, 0, 0, 0, time.UTC),
	}
	historyplgn[9] = HistoryCust{
		Customer:  Customer{Name: "Prima Doe"},
		SparePart: SparePart{Name: "Kampas Kopling"},
		Date:      time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC),
	}

	historyplgn[10] = HistoryCust{
		Customer:  Customer{Name: "Jane Smith"},
		SparePart: SparePart{Name: "Oil Filter"},
		Date:      time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC),
	}

	historyplgn[11] = HistoryCust{
		Customer:  Customer{Name: "Lintang Brown"},
		SparePart: SparePart{Name: "Lampu utama"},
		Date:      time.Date(2024, 4, 10, 0, 0, 0, 0, time.UTC),
	}

	//Inibuat perhitungan panjang data saat ditampilkan
	historyFreqCount := 12
	countSpareParts := 25
	countTransactions := 0
	historyPlgnCount := 12

	var keluar bool

	for !keluar {
		clear()
		fmt.Println(strings.Repeat("=", 70))
		fmt.Print(strings.Repeat("=", 29))
		fmt.Print(" Menu Utama ")
		fmt.Println(strings.Repeat("=", 29))
		fmt.Println(strings.Repeat("=", 70))
		fmt.Println("1. Pesan Spare-part")
		fmt.Println("2. Tampil Transaksi")
		fmt.Println("3. Transaksi Pembayaran")
		fmt.Println("4. Riwayat Pembelian Pelanggan")
		fmt.Println("5. Spare-part yang sering dibeli")
		fmt.Println("6. Keluar Program")
		fmt.Println(strings.Repeat("=", 70))

		var input int
		fmt.Print("Pilihan: ")
		fmt.Scanln(&input)

		if input == 6 {
			clear()
			fmt.Println("Terima kasih!")
			loading()
			keluar = true
		} else if input == 1 {
			clear()
			kelolaSparePart(&spareParts, &countSpareParts, &transactions, &countTransactions)
		} else if input == 2 {
			clear()
			tampilTransaction(transactions, countTransactions) // Gunakan countTransactions langsung
			fmt.Println("Tekan Enter untuk kembali...")
			fmt.Scanln()

		} else if input == 3 {
			clear()
			processTransaction(&transactions, &countTransactions, &historyfrq, &historyFreqCount, &historyplgn, &historyPlgnCount)

		} else if input == 4 {
			tampilkanHistoryPelanggan(historyplgn[:], historyPlgnCount)
			daftarPelanggan(historyplgn, historyPlgnCount)
			fmt.Println("Tekan Enter untuk kembali...")
			fmt.Scanln()
		} else if input == 5 {
			clear()
			tampilkanHistoryFreq(historyfrq, historyFreqCount)
			fmt.Println("Tekan Enter untuk kembali...")
			fmt.Scanln()
		} else {
			fmt.Println(strings.Repeat("=", 80))
			fmt.Println("Pilihan tidak valid")

			loading()
		}
	}
}

// main
func main() {
	menuUtama()
}
