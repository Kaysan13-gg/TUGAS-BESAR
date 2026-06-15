package main

import "fmt"

type Warga struct {
	id         int
	nama       string
	alamat     string
	totalBerat float64
}

type Setoran struct { 
	idWarga int
	tanggal string
	jenis   string
	berat   float64
}

type ArrWarga [100]Warga
type ArrSetoran [500]Setoran

func tambahWarga(tabWarga *ArrWarga, n *int, idCounter *int) {
	if *n >= 100 {
		fmt.Println("Data sudah penuh!")
		return
	}

	var nama, alamat string
	fmt.Print("Nama   : ")
	fmt.Scan(&nama)
	fmt.Print("Alamat : ")
	fmt.Scan(&alamat)

	*idCounter++
	tabWarga[*n].id = *idCounter
	tabWarga[*n].nama = nama
	tabWarga[*n].alamat = alamat
	tabWarga[*n].totalBerat = 0
	*n++

	fmt.Println("Warga berhasil ditambahkan! ID:", *idCounter)
}

func ubahWarga(tabWarga *ArrWarga, n int) {
	var idCari int
	fmt.Print("Masukkan ID warga yang ingin diubah: ")
	fmt.Scan(&idCari)

	idx := cariWargaByID(*tabWarga, n, idCari)
	if idx == -1 {
		fmt.Println("Warga tidak ditemukan.")
		return
	}

	var namaBaru, alamatBaru string
	fmt.Println("Data sekarang - Nama:", tabWarga[idx].nama, "| Alamat:", tabWarga[idx].alamat)
	fmt.Print("Nama baru   : ")
	fmt.Scan(&namaBaru)
	fmt.Print("Alamat baru : ")
	fmt.Scan(&alamatBaru)

	tabWarga[idx].nama = namaBaru
	tabWarga[idx].alamat = alamatBaru
	fmt.Println("Data berhasil diubah!")
}

func hapusWarga(tabWarga *ArrWarga, n *int) {
	var idCari int
	fmt.Print("Masukkan ID warga yang ingin dihapus: ")
	fmt.Scan(&idCari)

	idx := cariWargaByID(*tabWarga, *n, idCari)
	if idx == -1 {
		fmt.Println("Warga tidak ditemukan.")
		return
	}

	namaHapus := tabWarga[idx].nama

	i := idx
	for i < *n-1 {
		tabWarga[i] = tabWarga[i+1]
		i++
	}
	*n--

	fmt.Println("Warga", namaHapus, "berhasil dihapus!")
}

func tampilSemuaWarga(tabWarga ArrWarga, n int) {
	if n == 0 {
		fmt.Println("Belum ada data warga.")
		return
	}

	fmt.Println("----------------------------------------------------------")
	fmt.Println("ID   | Nama                | Alamat         | Total (kg)")
	fmt.Println("----------------------------------------------------------")
	i := 0
	for i < n {
		fmt.Printf("%-4d | %-20s | %-14s | %.2f\n",
			tabWarga[i].id, tabWarga[i].nama,
			tabWarga[i].alamat, tabWarga[i].totalBerat)
		i++
	}
	fmt.Println("----------------------------------------------------------")
	fmt.Println("Total warga:", n)
}

func catatSetoran(tabSetoran *ArrSetoran, m *int, tabWarga *ArrWarga, n int) {
	if *m >= 500 {
		fmt.Println("Log setoran sudah penuh!")
		return
	}

	var idWarga int
	fmt.Print("ID Warga             : ")
	fmt.Scan(&idWarga)

	idx := cariWargaByID(*tabWarga, n, idWarga)
	if idx == -1 {
		fmt.Println("Warga tidak ditemukan.")
		return
	}

	var tanggal, jenis string
	var berat float64
	fmt.Print("Tanggal (DD/MM/YYYY) : ")
	fmt.Scan(&tanggal)
	fmt.Print("Jenis sampah         : ")
	fmt.Scan(&jenis)
	fmt.Print("Berat (kg)           : ")
	fmt.Scan(&berat)

	tabSetoran[*m].idWarga = idWarga
	tabSetoran[*m].tanggal = tanggal
	tabSetoran[*m].jenis = jenis
	tabSetoran[*m].berat = berat
	*m++

	tabWarga[idx].totalBerat = tabWarga[idx].totalBerat + berat

	fmt.Printf("Setoran %.2f kg dari %s berhasil dicatat!\n", berat, tabWarga[idx].nama)
}

func tampilSetoranWarga(tabSetoran ArrSetoran, m int, tabWarga ArrWarga, n int) {
	var idCari int
	fmt.Print("Masukkan ID warga: ")
	fmt.Scan(&idCari)

	idx := cariWargaByID(tabWarga, n, idCari)
	if idx == -1 {
		fmt.Println("Warga tidak ditemukan.")
		return
	}

	fmt.Println("Riwayat setoran untuk:", tabWarga[idx].nama)
	fmt.Println("No | Tanggal    | Jenis       | Berat (kg)")
	fmt.Println("-------------------------------------------")

	no := 1
	i := 0
	for i < m {
		if tabSetoran[i].idWarga == idCari {
			fmt.Printf("%-2d | %-10s | %-11s | %.2f\n",
				no, tabSetoran[i].tanggal,
				tabSetoran[i].jenis, tabSetoran[i].berat)
			no++
		}
		i++
	}

	if no == 1 {
		fmt.Println("Warga ini belum punya setoran.")
	}
	fmt.Printf("Total sampah terkumpul: %.2f kg\n", tabWarga[idx].totalBerat)
}

func cariWargaByID(tabWarga ArrWarga, n int, idCari int) int {
	found := -1
	i := 0
	for i < n && found == -1 {
		if tabWarga[i].id == idCari {
			found = i
		}
		i++
	}
	return found
}

func cariWargaByNama(tabWarga ArrWarga, n int, namaCari string) int {
	found := -1
	i := 0
	for i < n && found == -1 {
		if tabWarga[i].nama == namaCari {
			found = i
		}
		i++
	}
	return found
}

func binarySearchByID(tabWarga ArrWarga, n int, idCari int) int {
	kiri := 0
	kanan := n - 1
	found := -1

	for kiri <= kanan && found == -1 {
		tengah := (kiri + kanan) / 2
		if idCari < tabWarga[tengah].id {
			kanan = tengah - 1
		} else if idCari > tabWarga[tengah].id {
			kiri = tengah + 1
		} else {
			found = tengah
		}
	}
	return found
}

func menuCariWarga(tabWarga ArrWarga, n int) {
	var pilih int
	fmt.Println("Cari berdasarkan:")
	fmt.Println("1. Nama (Sequential Search)")
	fmt.Println("2. ID   (Sequential Search)")
	fmt.Println("3. ID   (Binary Search - data harus sudah diurutkan by ID)")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		var namaCari string
		fmt.Print("Nama yang dicari: ")
		fmt.Scan(&namaCari)
		idx := cariWargaByNama(tabWarga, n, namaCari)
		if idx == -1 {
			fmt.Println("Warga tidak ditemukan.")
		} else {
			fmt.Println("Ditemukan!")
			fmt.Printf("ID: %d | Nama: %s | Alamat: %s | Total: %.2f kg\n",
				tabWarga[idx].id, tabWarga[idx].nama,
				tabWarga[idx].alamat, tabWarga[idx].totalBerat)
		}
	} else if pilih == 2 {
		var idCari int
		fmt.Print("ID yang dicari: ")
		fmt.Scan(&idCari)
		idx := cariWargaByID(tabWarga, n, idCari)
		if idx == -1 {
			fmt.Println("Warga tidak ditemukan.")
		} else {
			fmt.Println("Ditemukan!")
			fmt.Printf("ID: %d | Nama: %s | Alamat: %s | Total: %.2f kg\n",
				tabWarga[idx].id, tabWarga[idx].nama,
				tabWarga[idx].alamat, tabWarga[idx].totalBerat)
		}
	} else if pilih == 3 {
		var idCari int
		fmt.Print("ID yang dicari: ")
		fmt.Scan(&idCari)
		idx := binarySearchByID(tabWarga, n, idCari)
		if idx == -1 {
			fmt.Println("Warga tidak ditemukan.")
		} else {
			fmt.Println("Ditemukan!")
			fmt.Printf("ID: %d | Nama: %s | Alamat: %s | Total: %.2f kg\n",
				tabWarga[idx].id, tabWarga[idx].nama,
				tabWarga[idx].alamat, tabWarga[idx].totalBerat)
		}
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func selectionSort(tabWarga *ArrWarga, n int) {
	i := 0
	for i < n-1 {
		idxTerbesar := i
		j := i + 1
		for j < n {
			if tabWarga[j].totalBerat > tabWarga[idxTerbesar].totalBerat {
				idxTerbesar = j
			}
			j++
		}
		temp := tabWarga[idxTerbesar]
		tabWarga[idxTerbesar] = tabWarga[i]
		tabWarga[i] = temp
		i++
	}
}

func insertionSort(tabWarga *ArrWarga, n int) {
	i := 1
	for i < n {
		temp := tabWarga[i]
		j := i
		for j > 0 && temp.totalBerat > tabWarga[j-1].totalBerat {
			tabWarga[j] = tabWarga[j-1]
			j--
		}
		tabWarga[j] = temp
		i++
	}
}

func menuUrutkan(tabWarga *ArrWarga, n int) {
	var pilih int
	fmt.Println("Urutkan dengan:")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		selectionSort(tabWarga, n)
		fmt.Println("Data berhasil diurutkan dengan Selection Sort!")
	} else if pilih == 2 {
		insertionSort(tabWarga, n)
		fmt.Println("Data berhasil diurutkan dengan Insertion Sort!")
	} else {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	tampilSemuaWarga(*tabWarga, n)
}

func statistikMingguan(tabSetoran ArrSetoran, m int) {
	if m == 0 {
		fmt.Println("Belum ada data setoran.")
		return
	}

	var total float64 = 0
	i := 0
	for i < m {
		total = total + tabSetoran[i].berat
		i++
	}

	rataRata := total / float64(m)

	fmt.Println("===== STATISTIK MINGGUAN =====")
	fmt.Println("Jumlah transaksi :", m)
	fmt.Printf("Total berat      : %.2f kg\n", total)
	fmt.Printf("Rata-rata/setor  : %.2f kg\n", rataRata)
	fmt.Println("==============================")
}

func tampilMenu() {
	fmt.Println()
	fmt.Println("======== WASTE-TRACK ========")
	fmt.Println("1. Tambah warga")
	fmt.Println("2. Ubah data warga")
	fmt.Println("3. Hapus warga")
	fmt.Println("4. Lihat semua warga")
	fmt.Println("5. Catat setoran sampah")
	fmt.Println("6. Lihat riwayat setoran warga")
	fmt.Println("7. Cari warga")
	fmt.Println("8. Urutkan warga by berat sampah")
	fmt.Println("9. Statistik mingguan")
	fmt.Println("0. Keluar")
	fmt.Println("=============================")
	fmt.Print("Pilihan: ")
}

func main() {
	var tabWarga ArrWarga
	var tabSetoran ArrSetoran
	var nWarga int = 0
	var mSetoran int = 0
	var idCounter int = 0
	var pilihan int

	fmt.Println("Selamat datang di WASTE-TRACK!")
	fmt.Println("Aplikasi Manajemen Sampah Lingkungan")

	for {
		tampilMenu()
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahWarga(&tabWarga, &nWarga, &idCounter)
		} else if pilihan == 2 {
			ubahWarga(&tabWarga, nWarga)
		} else if pilihan == 3 {
			hapusWarga(&tabWarga, &nWarga)
		} else if pilihan == 4 {
			tampilSemuaWarga(tabWarga, nWarga)
		} else if pilihan == 5 {
			catatSetoran(&tabSetoran, &mSetoran, &tabWarga, nWarga)
		} else if pilihan == 6 {
			tampilSetoranWarga(tabSetoran, mSetoran, tabWarga, nWarga)
		} else if pilihan == 7 {
			menuCariWarga(tabWarga, nWarga)
		} else if pilihan == 8 {
			menuUrutkan(&tabWarga, nWarga)
		} else if pilihan == 9 {
			statistikMingguan(tabSetoran, mSetoran)
		} else if pilihan == 0 {
			fmt.Println("Terima kasih! Program selesai.")
			break
		} else {
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}