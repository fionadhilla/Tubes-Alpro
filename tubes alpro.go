package main

import "fmt"

const NMAX int = 6

//100
type perpus struct {
	info        [NMAX]buku
	nBuku       int
	pinjam      [20]bukupinjaman
	nPinjamBuku int
}

type bukupinjaman struct {
	namaOrang     string
	namaBuku      string
	tanggalpinjam tgl
	duedate       tgl
	dikembalikan  tgl
	denda         float64
	status        string
}

type tgl struct {
	tanggal int
	bulan   int
	tahun   int
}

type buku struct {
	nama    string
	tarif   float64
	jpinjam int
}

func main() {
	menuAwal()
}

func menuAwal() {
	var T perpus

	var pilihanuser string
	for pilihanuser != "11" {
		fmt.Println(" ")
		fmt.Println("---------------------------------------------------")
		fmt.Println("|-------------  APLIKASI PERPUSTAKAAN  -----------|")
		fmt.Println("|-------------------------------------------------|")
		fmt.Println("|  1. Tambah Buku                                 |")
		fmt.Println("|  2. Edit Buku                                   |")
		fmt.Println("|  3. Hapus Buku                                  |")
		fmt.Println("|  4. tambah Peminjam Buku                        |")
		fmt.Println("|  5. edit pinjaman Buku                          |")
		fmt.Println("|  6. hapus Peminjam Buku                         |")
		fmt.Println("|  7. Cari Buku                                   |")
		fmt.Println("|  8. tampilkan Array Buku                        |")
		fmt.Println("|  9. tampilkan Array Buku yang sedang dipinjam   |")
		fmt.Println("|  10. tampilkan 5 buku terfavorit                |")
		fmt.Println("|  11. EXIT                                       |")
		fmt.Println("---------------------------------------------------")
		fmt.Println("Masukkan Pilihan Anda : ")
		fmt.Scanln(&pilihanuser)

		if pilihanuser == "1" {
			inputbuku(&T, &T.nBuku)
			// fmt.Println(T.info)

		} else if pilihanuser == "2" {
			editbuku(&T)

		} else if pilihanuser == "3" {
			hapusbuku(&T, &T.nBuku)

		} else if pilihanuser == "4" {
			tambahPeminjamBuku(&T, &T.nPinjamBuku)

		} else if pilihanuser == "5" {
			editpinjambuku(&T, &T.nPinjamBuku)

		} else if pilihanuser == "6" {
			hapuspeminjambuku(&T, &T.nPinjamBuku)

		} else if pilihanuser == "7" {
			caribuku(T)

		} else if pilihanuser == "8" {
			printarray(T, T.nBuku)
			fmt.Println(T.info)
			fmt.Println("T.nbuku:", T.nBuku)

		} else if pilihanuser == "9" {
			tampilkanarrayygsedangdipinjam(T, T.nPinjamBuku)
			fmt.Println(T.pinjam)
			fmt.Println("T.npinjaman:", T.nPinjamBuku)

		} else if pilihanuser == "10" {
			printterfavorit(T, T.nBuku)

		} else if pilihanuser == "11" {
			fmt.Println("keluar...")
		} else {
			fmt.Println("pilihan tidak ada dalam menu harap masukkan antar 1 sampai 11 ")
			// fmt.Scanln(&pilihanuser)
		}

	}

}

func inputbuku(T *perpus, n *int) {
	var nama string
	var tar float64
	if *n < NMAX {
		fmt.Print("masukkan data buku (nama dan tarif buku): ", "\n")
		fmt.Scanln(&nama, &tar)

		for !sama(*T, nama) && nama != "-" {
			// fmt.Scanln(&tar)
			T.info[*n].nama = nama
			T.info[*n].tarif = tar
			T.nBuku = *n
			*n++
			if *n < NMAX {
				fmt.Scanln(&nama, &tar)
			} else {
				fmt.Print("\n", "(!!!)PERINGATAN : array tidak cukup")
				fmt.Println(" ")
			}
		}
	} else {
		fmt.Print("\n", "(!!!)PERINGATAN : array tidak cukup")
		fmt.Println(" ")
	}
	// fmt.Println("n buku :")
	// fmt.Print(*n)

	// fmt.Println(T.info)
}

func sama(T perpus, s string) bool {
	for i := 0; i <= T.nBuku; i++ {
		if T.info[i].nama == s {
			return true
		}
	}
	return false
}

func printarray(T perpus, n int) {
	fmt.Println("buku yang ada di library")
	for i := 0; i < n; i++ {
		fmt.Println("Nama Buku: ", T.info[i].nama, "     Tarif: ", T.info[i].tarif, "     Jumlah Dipinjam: ", T.info[i].jpinjam)
	}
}

func caribuku(T perpus) {
	var nama, pilihan string

	for pilihan != "4" {
		fmt.Println(" ")
		fmt.Println("-------------------------------------- ")
		fmt.Println("|------- Cari Berdasarkan Kategori -------|")
		fmt.Println("|	1. Berdasarkan Nama               |")
		fmt.Println("|	2. Berdasarkan Tarif              |")
		fmt.Println("|	3. Berdasarkan Jumlah Peminjaman  |")
		fmt.Println("|	4. Kembali                        |")
		fmt.Println("-------------------------------------- ")
		fmt.Println("Masukkan Pilihan Anda: ")
		fmt.Scanln(&pilihan)

		if pilihan == "1" {
			fmt.Println("Nama yg ingin dicari :")
			fmt.Scanln(&nama)
			var hasil1 int
			hasil1 = cariBerdasarkanaNama(T, nama)
			if hasil1 >= 0 {
				fmt.Print("nama : ", T.info[hasil1].nama, "\n", "tarif : ", T.info[hasil1].tarif, "\n", "berapa kali dipinjam : ", T.info[hasil1].jpinjam, "\n")
			} else {
				fmt.Println("(!!!)PERINGATAN : Nama Tidak Ditemukan")
			}
		} else if pilihan == "2" {
			var tarif float64
			fmt.Print("Tarif Minimal yang Ingin Dicari :")
			fmt.Scanln(&tarif)
			cariberdasarkantarif(T, tarif)

		} else if pilihan == "3" {
			var hasil3 int
			var pinjam int
			fmt.Print("Jumlah Pinjaman yang Ingin Dicari : ")
			fmt.Scanln(&pinjam)
			hasil3 = CariBerdasarkanJPinjam(T, pinjam)
			fmt.Print("\n", "nama : ", T.info[hasil3].nama, "\n", "tarif : ", T.info[hasil3].tarif, "\n", "berapa kali dipinjam : ", T.info[hasil3].jpinjam, "\n")
		} else if pilihan == "4" {
			fmt.Println("kembali ke menu utama...")
		} else {
			fmt.Println("(!!!)PERINGATAN : pilihan tidak ada pada menu")
		}
	}
}

func cariBerdasarkanaNama(T perpus, s string) int {
	var found int = -1
	var i int
	if T.nBuku >= 0 {
		for i < T.nBuku && found == -1 {
			if T.info[i].nama == s {
				found = i
			}
			i++
		}
	} else {
		fmt.Println("(!!!)PERINGATAN : Array Kosong")
	}

	return found

}

func cariberdasarkantarif(T perpus, s float64) {
	for i := 0; i < len(T.info); i++ {
		if T.info[i].tarif >= s {
			fmt.Print("\n", "Nama : ", T.info[i].nama, "\n", "Tarif : ", T.info[i].tarif, "\n", "Berapa Kali Dipinjam : ", T.info[i].jpinjam, "\n")
		}
	}
}

func CariBerdasarkanJPinjam(T perpus, s int) int {
	var found int = -1
	var i int
	for i < T.nBuku && found == -1 {
		if T.info[i].jpinjam == s {
			found = i
		}
		i++
	}
	return found
}

func editbuku(T *perpus) {
	var nama, namabaru string
	var tarifbaru float64
	var idxbuku, jpinjambaru int
	var pilihan string
	fmt.Println("Masukkan nama buku yang mau diubah : ")
	fmt.Scanln(&nama)
	idxbuku = cariBerdasarkanaNama(*T, nama)

	if idxbuku >= 0 {
		for pilihan != "4" {
			fmt.Println(" ")
			fmt.Println("----------------------------------")
			fmt.Println("|---kategori yang ingin diubah---|")
			fmt.Println("|  1. Nama                       |")
			fmt.Println("|  2. Tarif                      |")
			fmt.Println("|  3. Jumlah Pinjaman            |")
			fmt.Println("|  4. Membali                    |")
			fmt.Println("----------------------------------")
			fmt.Println("masukkan pilihan anda : ")
			fmt.Scanln(&pilihan)

			if pilihan == "1" {
				fmt.Println("Nama Buku Sebelumnya : ", T.info[idxbuku].nama)
				fmt.Println("Masukkan Nama Baru : ")
				fmt.Scanln(&namabaru)
				var pilihanyakin string
				fmt.Println("------------------------------")
				fmt.Println("Apakah yakin ingin mengubah? y/n")
				fmt.Scanln(&pilihanyakin)

				if pilihanyakin == "y" {
					var apakahada int
					apakahada = cariBerdasarkanaNama(*T, namabaru)
					if apakahada >= 0 {

						T.info[idxbuku].nama = namabaru
						fmt.Println("data berhasil diubah")
						fmt.Println(" ")

					} else {
						fmt.Println("(!!!)PERINGATAN : Nama buku sudah ada, silahkan masukkan nama lain")
					}
				} else if pilihanyakin == "n" {
					fmt.Println("(!!!)PERINGATAN : data tidak berhasil diubah")
				} else {
					fmt.Println("(!!!)PERINGATAN : pilihan tidak ada dalam menu, tolong masukkan antara y atau n")
					fmt.Scanln(&pilihanyakin)
				}

			} else if pilihan == "2" {
				fmt.Println("jumlah tarif sebelumnya :", T.info[idxbuku].tarif)
				fmt.Println("masukkan jumlah tarif baru : ")
				fmt.Scanln(&tarifbaru)
				var pilihanyakin string

				fmt.Println("------------------------------")
				fmt.Println("apakah yakin ingin mengubah? y/n")
				fmt.Scanln(&pilihanyakin)

				if pilihanyakin == "y" {
					T.info[idxbuku].tarif = tarifbaru
					fmt.Println("data berhasil diubah")
				} else if pilihanyakin == "n" {
					fmt.Println("data tidak berhasil diubah")
				} else {
					fmt.Println("pilihan tidak ada dalam menu, tolong masukkan antara y atau n")
					fmt.Scanln(&pilihanyakin)
				}

			} else if pilihan == "3" {
				fmt.Println("jumlah dipinjam sebelumnya :", T.info[idxbuku].jpinjam)
				fmt.Println("masukkan jumlah dipinjam baru : ")
				fmt.Scanln(&jpinjambaru)
				var pilihanyakin string

				fmt.Println("------------------------------")
				fmt.Println("apakah yakin ingin mengubah? y/n")
				fmt.Scanln(&pilihanyakin)

				if pilihanyakin == "y" {
					T.info[idxbuku].jpinjam = jpinjambaru
					fmt.Println("data berhasil diubah")
				} else if pilihanyakin == "n" {
					fmt.Println("data tidak berhasil diubah")
				} else {
					fmt.Println("pilihan tidak ada dalam menu, tolong masukkan antara y atau n")
					fmt.Scanln(&pilihanyakin)
				}

			} else if pilihan == "4" {
				fmt.Println("kembali ke menu utama...")
			} else {
				fmt.Println("(!!!)PERINGATAN : pilihan tidak ada dalam pilihan, silahkan masukkan nomor sesuai pada menu")

			}
		}
	} else {
		fmt.Println(" ")
		fmt.Println("(!!!)PERINGATAN : buku tidak ditemukan")
		fmt.Println("kembali ke menu utama...")

	}

}

func hapusbuku(T *perpus, n *int) {
	var namabuku string
	var idx int
	fmt.Println("masukkan nama buku : ")
	fmt.Scanln(&namabuku)
	idx = cariBerdasarkanaNama(*T, namabuku)
	if idx >= 0 {
		fmt.Print("\n", "nama buku : ", T.info[idx].nama, "\n", "dengan tarif ", T.info[idx].tarif, "\n", "jumlah pinjaman ", T.info[idx].jpinjam, "\n")
		fmt.Println("apakah yakin menghapus buku? y/n")
		var pilihanuser string
		fmt.Scanln(&pilihanuser)

		if pilihanuser == "y" {
			for i := idx; i < *n-1; i++ {
				T.info[i].nama = T.info[i+1].nama
				T.info[i].tarif = T.info[i+1].tarif
				T.info[i].jpinjam = T.info[i+1].jpinjam
			}
			T.info[*n-1].nama = ""
			T.info[*n-1].tarif = 0
			T.info[*n-1].jpinjam = 0
			*n--

		} else if pilihanuser == "n" {
			fmt.Println("data tidak jadi dihapus")

		} else {
			fmt.Println("pilihan tidak ada pada menu, tolong masukkan antara y atau n")
			fmt.Scanln(&pilihanuser)
		}

	} else {
		fmt.Println("buku tidak ditemukan")
	}

}

func tambahPeminjamBuku(T *perpus, n *int) {
	var namaorang string
	var namabuku string
	// var ilibuku int
	var tglhari, tglbulan, tgltahun, tglkembali, blnkembali, thnkembali int
	fmt.Print("jumlah buku yg bisa dipinjam :", T.nBuku, "\n")

	if *n < NMAX {
		fmt.Print("masukkan data peminjam (nama orang, nama buku, tanggal pinjam, bulan, dan tahun, serta due date, month, year): ", "\n")
		fmt.Scanln(&namaorang, &namabuku, &tglhari, &tglbulan, &tgltahun, &tglkembali, &blnkembali, &thnkembali)

		if namaorang != "-" {
			if apakahbukuada(*T, T.nBuku, namabuku) {
				// fmt.Scanln(&tglhari, &tglbulan, &tgltahun)
				if valid(tglhari, tglbulan, tgltahun) && valid(tglkembali, blnkembali, thnkembali) {
					if !mengeceknamaBuku(*T, *n, namabuku) {
						// apakahbukuada(*T, T.nBuku, namabuku)
						// fmt.Scanln(&tglkembali, &blnkembali, &thnkembali)
						T.pinjam[*n].namaOrang = namaorang
						T.pinjam[*n].namaBuku = namabuku
						T.pinjam[*n].tanggalpinjam.tanggal = tglhari
						T.pinjam[*n].tanggalpinjam.bulan = tglbulan
						T.pinjam[*n].tanggalpinjam.tahun = tgltahun
						T.pinjam[*n].duedate.tanggal = tglkembali
						T.pinjam[*n].duedate.bulan = blnkembali
						T.pinjam[*n].duedate.tahun = thnkembali
						T.pinjam[*n].status = "dipinjam"
						// ilibuku = cariBerdasarkanaNama(*T, namabuku)
						// T.info[ilibuku].jpinjam++
						*n++
						// fmt.Println("data sudah dimasukkan ke array")
						if *n < NMAX {
							fmt.Scanln(&namaorang, &namabuku, &tglhari, &tglbulan, &tgltahun, &tglkembali, &blnkembali, &thnkembali)
						} else {
							fmt.Println("buku sudah terpinjam semua")
						}

					} else {
						fmt.Sprintln("buku sudah dipinjam")
					}

				} else {
					fmt.Println("tanggal tidak valid")
				}
			} else {
				fmt.Println("buku tidak ada pada library")
			}

		}

		// fmt.Println("buku tidak ditemukan")

		// fmt.Scanln(&namaorang, &namabuku, &tglhari, &tglbulan, &tgltahun)
	} else {
		fmt.Println("buku telah dipinjam semua")
	}

	// fmt.Println(T.pinjam)
}

func mengeceknamaBuku(T perpus, n int, nama string) bool {
	for i := 0; i <= n; i++ {
		if nama == T.pinjam[i].namaBuku && T.pinjam[i].status != "complete" {
			return true
		}
	}
	return false
}

func apakahbukuada(T perpus, n int, s string) bool {
	for i := 0; i <= n; i++ {
		if s == T.info[i].nama {
			return true
		}
	}
	return false
}

func tampilkanarrayygsedangdipinjam(T perpus, n int) {
	for i := 0; i < n; i++ {
		if T.pinjam[i].status == "dipinjam" {
			fmt.Println("nama peminjam: ", T.pinjam[i].namaOrang, "     nama buku: ", T.pinjam[i].namaBuku)
			fmt.Println("tanggal peminjaman: ", T.pinjam[i].tanggalpinjam.tanggal, "/", T.pinjam[i].tanggalpinjam.bulan, "/", T.pinjam[i].tanggalpinjam.tahun)
			fmt.Println("due date: ", T.pinjam[i].duedate.tanggal, "/", T.pinjam[i].duedate.bulan, "/", T.pinjam[i].duedate.tahun)
			fmt.Println("dikembalikan: ", T.pinjam[i].dikembalikan.tanggal, "/", T.pinjam[i].dikembalikan.bulan, "/", T.pinjam[i].dikembalikan.tahun)
			fmt.Println("denda: ", T.pinjam[i].denda)
			fmt.Println("status: ", T.pinjam[i].status)
			// fmt.Println("status : ", T.pinjam[i].status)
			fmt.Println(" ")
		}

	}
}

func valid(tanggal, bulan, tahun int) bool {
	if bulan == 2 && kabisat(tahun) {
		return (tahun > 0 && (bulan > 0 && bulan < 13) && (tanggal > 0 && tanggal <= 29))
	} else if bulan == 2 {
		return (tahun > 0 && (bulan > 0 && bulan < 13) && (tanggal > 0 && tanggal <= 28))
	} else if bulan < 8 && (bulan%2 == 1) || bulan > 6 && (bulan%2 == 0) {
		return (tahun > 0 && (bulan > 0 && bulan < 13) && (tanggal > 0 && tanggal <= 31))
	} else {
		return (tahun > 0 && (bulan > 0 && bulan < 13) && (tanggal > 0 && tanggal <= 30))
	}
}

func kabisat(tahun int) bool {
	if (tahun%400 == 0) && (tahun%100 == 0) {
		return true
	} else if (tahun%4 == 0) && (tahun%100 != 0) {
		return true
	}
	return false

}

func editpinjambuku(T *perpus, n *int) {
	var nama, namabukubaru, namaorangbaru string
	var idx int
	var pilihan string
	fmt.Println("masukkan nama buku yang mau diubah : ")
	fmt.Scanln(&nama)
	idx = caripeminjamdarinamabuku(*T, nama)

	if idx >= 0 {

		fmt.Println("data sebelumnya: ")
		fmt.Println("nama peminjam: ", T.pinjam[idx].namaOrang, "     nama buku: ", T.pinjam[idx].namaBuku)
		fmt.Println("tanggal peminjaman: ", T.pinjam[idx].tanggalpinjam.tanggal, "/", T.pinjam[idx].tanggalpinjam.bulan, "/", T.pinjam[idx].tanggalpinjam.tahun)
		fmt.Println("due date: ", T.pinjam[idx].duedate.tanggal, "/", T.pinjam[idx].duedate.bulan, "/", T.pinjam[idx].duedate.tahun)
		fmt.Println("dikembalikan: ", T.pinjam[idx].dikembalikan.tanggal, "/", T.pinjam[idx].dikembalikan.bulan, "/", T.pinjam[idx].dikembalikan.tahun)
		fmt.Println("denda: ", T.pinjam[idx].denda)
		fmt.Println("status: ", T.pinjam[idx].denda)
		fmt.Println(" ")

		for pilihan != "7" {
			fmt.Println(" ")
			fmt.Println("--------------------------------")
			fmt.Println("|--Kategori yang Ingin Diubah--|")
			fmt.Println("|  1. nama  peminjam           |")
			fmt.Println("|  2. nama buku                |")
			fmt.Println("|  3. tanggal pinjam           |")
			fmt.Println("|  4. duedate                  |")
			fmt.Println("|  5. tanggal dikembalikan     |")
			fmt.Println("|  6. denda                    |")
			fmt.Println("|  7. kembali                  |")
			fmt.Println("--------------------------------")
			fmt.Println("masukkan pilihan : ")
			fmt.Scanln(&pilihan)

			if pilihan == "1" {
				fmt.Println("nama peminjam sebelumnya : ", T.pinjam[idx].namaOrang)
				fmt.Println("masukkan nama baru : ")
				fmt.Scanln(&namaorangbaru)
				var pilihanyakin string

				fmt.Println("------------------------------")
				fmt.Println("apakah yakin ingin mengubah? y/n")
				fmt.Scanln(&pilihanyakin)

				if pilihanyakin == "y" {
					T.pinjam[idx].namaOrang = namaorangbaru
					fmt.Println("data berhasil diubah")

				} else if pilihanyakin == "n" {
					fmt.Println("data tidak berhasil diubah")
				} else {
					fmt.Println("pilihan tidak ada dalam menu, tolong masukkan antara y atau n")
					fmt.Scanln(&pilihanyakin)
				}

			} else if pilihan == "2" {
				fmt.Println("nama buku sebelumnya : ", T.pinjam[idx].namaBuku)
				fmt.Println("masukkan nama baru : ")
				fmt.Scanln(&namabukubaru)
				var pilihanyakin string

				fmt.Println("------------------------------")
				fmt.Println("apakah yakin ingin mengubah? y/n")

				fmt.Scanln(&pilihanyakin)

				if pilihanyakin == "y" {
					var idxbaru int
					idxbaru = cariBerdasarkanaNama(*T, namabukubaru)
					if idxbaru >= 0 {
						if T.pinjam[caripeminjamdarinamabuku(*T, namabukubaru)].status != "dipinjam" {
							T.pinjam[idx].namaBuku = namabukubaru
						} else {
							fmt.Println("buku sedang dipinjam")
						}
					} else {
						fmt.Println("buku tidak ada di library")
					}
					fmt.Println("data berhasil diubah")
				} else if pilihanyakin == "n" {
					fmt.Println("data tidak berhasil diubah")
				} else {
					fmt.Println("pilihan tidak ada dalam menu, tolong masukkan antara y atau n")
					fmt.Scanln(&pilihanyakin)
				}

			} else if pilihan == "3" {
				fmt.Println("tanggal peminjaman sebelumnya : ", T.pinjam[idx].tanggalpinjam.tanggal)
				fmt.Println("masukkan tanggal baru (tgl bln thn) : ")
				var tBaru tgl
				fmt.Scanln(&tBaru.tanggal, &tBaru.bulan, &tBaru.tahun)
				var pilihanyakin string

				fmt.Println("------------------------------")
				fmt.Println("apakah yakin ingin mengubah? y/n")

				fmt.Scanln(&pilihanyakin)

				if pilihanyakin == "y" {
					if valid(tBaru.tanggal, tBaru.bulan, tBaru.tahun) {

						T.pinjam[idx].tanggalpinjam.tanggal = tBaru.tanggal
						T.pinjam[idx].tanggalpinjam.bulan = tBaru.bulan
						T.pinjam[idx].tanggalpinjam.tahun = tBaru.tahun
						fmt.Println("data berhasil diubah")

					} else {
						fmt.Println("(!!!)PERINGATAN : Tanggal Tidak Valid")
					}

				} else if pilihanyakin == "n" {
					fmt.Println("data tidak berhasil diubah")
				} else {
					fmt.Println("pilihan tidak ada dalam menu, tolong masukkan antara y atau n")
					fmt.Scanln(&pilihanyakin)
				}

			} else if pilihan == "4" {
				fmt.Println("tanggal due date sebelumnya : ", T.pinjam[idx].duedate.tanggal, " ", T.pinjam[idx].duedate.bulan, " ", T.pinjam[idx].duedate.tahun)
				// var tglrek, blnrek, thnrek int
				var tBaru tgl
				var pilihanyakin string
				// for valid(T.pinjam[idx].tanggalpinjam.tanggal, T.pinjam[idx].tanggalpinjam.bulan, T.pinjam[idx].tanggalpinjam.tahun) {
				// 	hitungTanggalKembali(T.pinjam[idx].tanggalpinjam.tanggal, T.pinjam[idx].tanggalpinjam.bulan, T.pinjam[idx].tanggalpinjam.tahun, &tglrek, &blnrek, &thnrek)
				// }
				// fmt.Println("rekomendasi duedate : ", tglrek, blnrek, thnrek)
				fmt.Println("rekomendasi due date adalah 3 hari setelah dipinjam")
				fmt.Println("masukkan tanggal : ")
				fmt.Scanln(&tBaru.tanggal, &tBaru.bulan, &tBaru.tahun)
				fmt.Println("------------------------------")
				fmt.Println("apakah yakin ingin mengubah? y/n")

				fmt.Scanln(&pilihanyakin)

				if pilihanyakin == "y" {
					if valid(tBaru.tanggal, tBaru.bulan, tBaru.tahun) {
						T.pinjam[idx].duedate.tanggal = tBaru.tanggal
						T.pinjam[idx].duedate.bulan = tBaru.bulan
						T.pinjam[idx].duedate.tahun = tBaru.tahun
						fmt.Println("data berhasil diubah")
					} else {
						fmt.Println("(!!!)PERINGATAN : tanggal tidak valid")
					}

				} else if pilihanyakin == "n" {
					fmt.Println("data tidak berhasil diubah")
				} else {
					fmt.Println("pilihan tidak ada dalam menu, tolong masukkan antara y atau n")
					fmt.Scanln(&pilihanyakin)
				}

			} else if pilihan == "5" {
				fmt.Println("tanggal dikembaikan sebelumnya:", T.pinjam[idx].dikembalikan.tanggal, " ", T.pinjam[idx].dikembalikan.bulan, " ", T.pinjam[idx].dikembalikan.tahun)
				fmt.Println("masukkan tanggal baru (tgl bln thn) : ")
				var tBaru tgl
				fmt.Scanln(&tBaru.tanggal, &tBaru.bulan, &tBaru.tahun)
				var pilihanyakin string

				fmt.Println("------------------------------")
				fmt.Println("apakah yakin ingin mengubah? y/n")
				fmt.Scanln(&pilihanyakin)

				if pilihanyakin == "y" {
					if valid(tBaru.tanggal, tBaru.bulan, tBaru.tanggal) {
						T.pinjam[idx].dikembalikan.tanggal = tBaru.tanggal
						T.pinjam[idx].dikembalikan.bulan = tBaru.bulan
						T.pinjam[idx].dikembalikan.tahun = tBaru.tahun
						fmt.Println("data berhasil diubah")
						if T.pinjam[idx].tanggalpinjam.tanggal > 0 && T.pinjam[idx].tanggalpinjam.bulan > 0 && T.pinjam[idx].tanggalpinjam.tahun > 0 {

							if T.pinjam[idx].duedate.tanggal > 0 && T.pinjam[idx].duedate.bulan > 0 && T.pinjam[idx].duedate.tahun > 0 {

								if T.pinjam[idx].dikembalikan.tanggal > 0 && T.pinjam[idx].dikembalikan.bulan > 0 && T.pinjam[idx].dikembalikan.tahun > 0 {

									fmt.Println("menghitung denda...")
									var indexslibrary int
									indexslibrary = cariBerdasarkanaNama(*T, T.pinjam[idx].namaBuku)

									// fmt.Println("Rp.", hitungdenda(*T, T.pinjam[idx].duedate.tanggal, T.pinjam[idx].duedate.bulan, T.pinjam[idx].duedate.tahun, T.pinjam[idx].dikembalikan.tanggal, T.pinjam[idx].dikembalikan.bulan, T.pinjam[idx].dikembalikan.tahun))
									fmt.Println("Rp", hitungdenda(*T, T.pinjam[idx].duedate.tanggal, T.pinjam[idx].duedate.bulan, T.pinjam[idx].duedate.tahun, T.pinjam[idx].dikembalikan.tanggal, T.pinjam[idx].dikembalikan.bulan, T.pinjam[idx].dikembalikan.tahun)*T.info[indexslibrary].tarif)
									T.pinjam[idx].denda = hitungdenda(*T, T.pinjam[idx].duedate.tanggal, T.pinjam[idx].duedate.bulan, T.pinjam[idx].duedate.tahun, T.pinjam[idx].dikembalikan.tanggal, T.pinjam[idx].dikembalikan.bulan, T.pinjam[idx].dikembalikan.tahun) * T.info[indexslibrary].tarif
									T.pinjam[idx].status = "Complete"
									T.info[indexslibrary].jpinjam++

								} else {
									fmt.Println("menghitung denda...")
									fmt.Println("(!!!)PERINGATAM : Penghitungan gagal")
									fmt.Println("(!!!)PERINGATAN : tanggal dikembalikan belum dimasukkan, tidak bisa menghitung denda")
								}

							} else {
								fmt.Println("menghitung denda...")
								fmt.Println("(!!!)PERINGATAM : Penghitungan gagal")
								fmt.Println("(!!!)PERINGATAN : tanggal due pengembalian belum dimasukkan, tidak bisa menghitung denda")
							}

						} else {
							fmt.Println("menghitung denda...")
							fmt.Println("(!!!)PERINGATAM : Penghitungan gagal")
							fmt.Println("(!!!)PERINGATAN : tanggal peminjaman belum dimasukkan, tidak bisa menghitung denda")
						}

					} else {
						fmt.Println("(!!!)PERINGATAN : Tanggal Tidak Valid")
					}
					//hitung dendanya dan otomatis berubah

				} else if pilihanyakin == "n" {
					fmt.Println("data tidak berhasil diubah")
				} else {
					fmt.Println("pilihan tidak ada dalam menu, tolong masukkan antara y atau n")
					fmt.Scanln(&pilihanyakin)
				}

			} else if pilihan == "6" {
				fmt.Println("nominal denda sebelumnya : ", "Rp.", T.pinjam[idx].denda)
				fmt.Println("masukkan nominal baru :")
				var dendabaru float64
				fmt.Scanln(&dendabaru)
				var pilihanyakin string

				fmt.Println("------------------------------")
				fmt.Println("apakah yakin ingin mengubah? y/n")

				fmt.Scanln(&pilihanyakin)

				if pilihanyakin == "y" {
					T.pinjam[idx].denda = dendabaru
					fmt.Println("data berhasil diubah")

				} else if pilihanyakin == "n" {
					fmt.Println("data tidak berhasil diubah")
				} else {
					fmt.Println("pilihan tidak ada dalam menu, tolong masukkan antara y atau n")
					fmt.Scanln(&pilihanyakin)
				}
			} else if pilihan == "7" {
				fmt.Println("kembali...")
			} else {
				fmt.Println("pilihan tidak ada pada menu mohon masukkan antara 1 sampai 7 ")
				fmt.Scanln(&pilihan)
			}
		}
	} else {
		fmt.Println("buku tidak ada pada data, masukkan ulang")

	}

}

func caripeminjamdarinamabuku(T perpus, s string) int {
	var found int = -1
	var i int = 0
	for i < T.nPinjamBuku && found == -1 {
		if T.pinjam[i].namaBuku == s {
			found = i
		}
		i++
	}
	return found
}

func printterfavorit(T perpus, n int) {
	urutkanbuku(&T, n)
	for i := 0; i < 6; i++ {
		fmt.Println(T.info[i].nama, T.info[i].tarif, T.info[i].jpinjam)
	}
}

//selection sort
func urutkanbuku(T *perpus, n int) {
	var i, j, idx_max int
	var t buku

	i = 1
	for i <= n-1 {
		idx_max = i - 1
		j = i
		for j < n {
			if T.info[idx_max].jpinjam < T.info[j].jpinjam {
				idx_max = j
			}
			j++
		}
		t = T.info[idx_max]
		T.info[idx_max] = T.info[i-1]
		T.info[i-1] = t
		i++
	}
}

// func getJumlahHari(bulan, tahun int, jmlHari *int) {
// 	var totalHari int
// 	for tahun > 1 {
// 		if kabisat(tahun) && bulan > 2 {
// 			totalHari = totalHari + 366
// 		} else {
// 			totalHari = totalHari + 365
// 		}
// 		tahun--
// 	}
// 	for i := 1; i != bulan; i++ {
// 		if i == 2 {
// 			totalHari = totalHari + 28
// 		} else if i < 8 && (i%2 == 1) || i > 6 && (i%2 == 0) {
// 			totalHari = totalHari + 31
// 		} else {
// 			totalHari = totalHari + 30
// 		}

// 	}
// 	*jmlHari = totalHari
// }

// func hitungTanggalKembali(tanggal1, bulan1, tahun1 int, tanggal2, bulan2, tahun2 *int) {
// 	var totalHari int
// 	getJumlahHari(bulan1, tahun1, &totalHari)
// 	totalHari = totalHari + tanggal1 + 3
// 	*tahun2 = 1
// 	for totalHari > 366 {
// 		if kabisat(*tahun2) && bulan1 > 2 {
// 			totalHari = totalHari - 366
// 			*tahun2++
// 		} else {
// 			totalHari = totalHari - 365
// 			*tahun2++
// 		}
// 	}

// 	*bulan2 = 1
// 	*tanggal2 = 0
// 	for totalHari > 29 {
// 		if *bulan2 == 2 {
// 			if kabisat(*tahun2) {
// 				totalHari = totalHari - 29
// 			} else {
// 				totalHari = totalHari - 28
// 			}
// 		} else if *bulan2 < 8 && (*bulan2%2 == 1) || *bulan2 > 6 && (*bulan2%2 == 0) {
// 			totalHari = totalHari - 31
// 		} else {
// 			totalHari = totalHari - 30
// 		}
// 		*bulan2++
// 		if *bulan2 == 2 && (totalHari > 28) {
// 			if kabisat(*tahun2) {
// 				totalHari = 29
// 			} else {
// 				totalHari = totalHari - 28
// 				*bulan2++
// 			}
// 		}
// 		if totalHari < 0 {
// 			*tanggal2 = (-totalHari)
// 			totalHari = 0
// 		}

// 	}
// 	*tanggal2 = *tanggal2 + totalHari
// }

func hapuspeminjambuku(T *perpus, n *int) {
	var namabuku string
	var outputhasil, outputhasil2, outputhasil3 int
	fmt.Println("masukkan nama buku yang ingin dihapus: ")
	fmt.Scanln(&namabuku)

	for i := 0; i < T.nPinjamBuku; i++ {
		if T.pinjam[i].namaBuku == namabuku {
			fmt.Println("nama peminjam: ", T.pinjam[i].namaOrang, "     nama buku: ", T.pinjam[i].namaBuku)
			fmt.Println("tanggal peminjaman: ", T.pinjam[i].tanggalpinjam.tanggal, "/", T.pinjam[i].tanggalpinjam.bulan, "/", T.pinjam[i].tanggalpinjam.tahun)
			fmt.Println("due date: ", T.pinjam[i].duedate.tanggal, "/", T.pinjam[i].duedate.bulan, "/", T.pinjam[i].duedate.tahun)
			fmt.Println("dikembalikan: ", T.pinjam[i].dikembalikan.tanggal, "/", T.pinjam[i].dikembalikan.bulan, "/", T.pinjam[i].dikembalikan.tahun)
			fmt.Println("denda: ", T.pinjam[i].denda)
			fmt.Println("status: ", T.pinjam[i].denda)
			fmt.Println(" ")
			outputhasil++

			if outputhasil == 1 {
				fmt.Println("apakah yakin menghapus buku? y/n")
				var pilihanuser string
				fmt.Scanln(&pilihanuser)
				if pilihanuser == "y" {
					var j int
					for j = i; j < *n-1; j++ {
						T.pinjam[j].namaBuku = T.pinjam[j+1].namaBuku
						T.pinjam[j].namaOrang = T.pinjam[j+1].namaOrang
						T.pinjam[j].tanggalpinjam.tanggal = T.pinjam[j+1].tanggalpinjam.tanggal
						T.pinjam[j].tanggalpinjam.bulan = T.pinjam[j+1].tanggalpinjam.bulan
						T.pinjam[j].tanggalpinjam.bulan = T.pinjam[j+1].tanggalpinjam.bulan
						T.pinjam[j].duedate.tanggal = T.pinjam[j+1].duedate.tanggal
						T.pinjam[j].duedate.bulan = T.pinjam[j+1].duedate.bulan
						T.pinjam[j].duedate.tahun = T.pinjam[j+1].duedate.tahun
						T.pinjam[j].dikembalikan.tanggal = T.pinjam[j+1].dikembalikan.tanggal
						T.pinjam[j].dikembalikan.bulan = T.pinjam[j+1].dikembalikan.bulan
						T.pinjam[j].dikembalikan.tahun = T.pinjam[j+1].dikembalikan.tahun
						T.pinjam[j].denda = T.pinjam[j+1].denda
						T.pinjam[j].status = T.pinjam[j+1].status

					}
					T.pinjam[*n-1].namaBuku = ""
					T.pinjam[*n-1].namaOrang = ""
					T.pinjam[*n-1].tanggalpinjam.tanggal = 0
					T.pinjam[*n-1].tanggalpinjam.bulan = 0
					T.pinjam[*n-1].tanggalpinjam.bulan = 0
					T.pinjam[*n-1].duedate.tanggal = 0
					T.pinjam[*n-1].duedate.bulan = 0
					T.pinjam[*n-1].duedate.tahun = 0
					T.pinjam[*n-1].dikembalikan.tanggal = 0
					T.pinjam[*n-1].dikembalikan.bulan = 0
					T.pinjam[*n-1].dikembalikan.tahun = 0
					T.pinjam[*n-1].denda = 0
					T.pinjam[*n-1].status = ""
					*n--

				} else if pilihanuser == "n" {
					fmt.Println("data tidak jadi dihapus")

				} else {
					fmt.Println("pilihan tidak ada pada menu, tolong masukkan antara y atau n")
					fmt.Scanln(&pilihanuser)
				}

			}
			//cari nama buku + cari nama orang
		} else {
			fmt.Println("nama buku tidak ada pada array")
		}
	}

	if outputhasil >= 2 {
		fmt.Println("ada banyak data yang sama masukkan nama peminjam : ")
		var namaorang string
		fmt.Scanln(&namaorang)
		for i := 0; i < T.nPinjamBuku; i++ {
			if T.pinjam[i].namaOrang == namaorang {
				fmt.Println("nama peminjam: ", T.pinjam[i].namaOrang, "     nama buku: ", T.pinjam[i].namaBuku)
				fmt.Println("tanggal peminjaman: ", T.pinjam[i].tanggalpinjam.tanggal, "/", T.pinjam[i].tanggalpinjam.bulan, "/", T.pinjam[i].tanggalpinjam.tahun)
				fmt.Println("due date: ", T.pinjam[i].duedate.tanggal, "/", T.pinjam[i].duedate.bulan, "/", T.pinjam[i].duedate.tahun)
				fmt.Println("dikembalikan: ", T.pinjam[i].dikembalikan.tanggal, "/", T.pinjam[i].dikembalikan.bulan, "/", T.pinjam[i].dikembalikan.tahun)
				fmt.Println("denda: ", T.pinjam[i].denda)
				fmt.Println("status: ", T.pinjam[i].denda)
				fmt.Println(" ")
				outputhasil2++

				if outputhasil2 == 1 {
					fmt.Println("apakah yakin menghapus buku? y/n")
					var pilihanuser string
					fmt.Scanln(&pilihanuser)
					if pilihanuser == "y" {
						for j := i; j < *n-1; j++ {
							T.pinjam[j].namaBuku = T.pinjam[j+1].namaBuku
							T.pinjam[j].namaOrang = T.pinjam[j+1].namaOrang
							T.pinjam[j].tanggalpinjam.tanggal = T.pinjam[j+1].tanggalpinjam.tanggal
							T.pinjam[j].tanggalpinjam.bulan = T.pinjam[j+1].tanggalpinjam.bulan
							T.pinjam[j].tanggalpinjam.bulan = T.pinjam[j+1].tanggalpinjam.bulan
							T.pinjam[j].duedate.tanggal = T.pinjam[j+1].duedate.tanggal
							T.pinjam[j].duedate.bulan = T.pinjam[j+1].duedate.bulan
							T.pinjam[j].duedate.tahun = T.pinjam[j+1].duedate.tahun
							T.pinjam[j].dikembalikan.tanggal = T.pinjam[j+1].dikembalikan.tanggal
							T.pinjam[j].dikembalikan.bulan = T.pinjam[j+1].dikembalikan.bulan
							T.pinjam[j].dikembalikan.tahun = T.pinjam[j+1].dikembalikan.tahun
							T.pinjam[j].denda = T.pinjam[j+1].denda
							T.pinjam[j].status = T.pinjam[j+1].status

						}
						T.pinjam[*n-1].namaBuku = ""
						T.pinjam[*n-1].namaOrang = ""
						T.pinjam[*n-1].tanggalpinjam.tanggal = 0
						T.pinjam[*n-1].tanggalpinjam.bulan = 0
						T.pinjam[*n-1].tanggalpinjam.bulan = 0
						T.pinjam[*n-1].duedate.tanggal = 0
						T.pinjam[*n-1].duedate.bulan = 0
						T.pinjam[*n-1].duedate.tahun = 0
						T.pinjam[*n-1].dikembalikan.tanggal = 0
						T.pinjam[*n-1].dikembalikan.bulan = 0
						T.pinjam[*n-1].dikembalikan.tahun = 0
						T.pinjam[*n-1].denda = 0
						T.pinjam[*n-1].status = ""
						*n--

					} else if pilihanuser == "n" {
						fmt.Println("data tidak jadi dihapus")

					} else {
						fmt.Println("pilihan tidak ada pada menu, tolong masukkan antara y atau n")
						fmt.Scanln(&pilihanuser)
					}

				}

			} else {
				fmt.Println("data tidak ditemukan")
			}
		}
	}

	if outputhasil2 >= 2 {
		fmt.Println("ada banyak data yang sama masukkan tanggal peminjaman : ")
		var hpstgl, hpsbln, hpsthn int
		fmt.Scanln(&hpstgl, &hpsbln, &hpsthn)
		for i := 0; i < T.nPinjamBuku; i++ {
			if T.pinjam[i].tanggalpinjam.tanggal == hpstgl && T.pinjam[i].tanggalpinjam.bulan == hpsbln && T.pinjam[i].tanggalpinjam.tahun == hpsthn {
				fmt.Println("nama peminjam: ", T.pinjam[i].namaOrang, "     nama buku: ", T.pinjam[i].namaBuku)
				fmt.Println("tanggal peminjaman: ", T.pinjam[i].tanggalpinjam.tanggal, "/", T.pinjam[i].tanggalpinjam.bulan, "/", T.pinjam[i].tanggalpinjam.tahun)
				fmt.Println("due date: ", T.pinjam[i].duedate.tanggal, "/", T.pinjam[i].duedate.bulan, "/", T.pinjam[i].duedate.tahun)
				fmt.Println("dikembalikan: ", T.pinjam[i].dikembalikan.tanggal, "/", T.pinjam[i].dikembalikan.bulan, "/", T.pinjam[i].dikembalikan.tahun)
				fmt.Println("denda: ", T.pinjam[i].denda)
				fmt.Println("status: ", T.pinjam[i].denda)
				fmt.Println(" ")
				outputhasil3++

				if outputhasil3 == 1 {
					fmt.Println("apakah yakin menghapus buku? y/n")
					var pilihanuser string
					fmt.Scanln(&pilihanuser)
					if pilihanuser == "y" {
						for j := i; j < *n-1; j++ {
							T.pinjam[j].namaBuku = T.pinjam[j+1].namaBuku
							T.pinjam[j].namaOrang = T.pinjam[j+1].namaOrang
							T.pinjam[j].tanggalpinjam.tanggal = T.pinjam[j+1].tanggalpinjam.tanggal
							T.pinjam[j].tanggalpinjam.bulan = T.pinjam[j+1].tanggalpinjam.bulan
							T.pinjam[j].tanggalpinjam.bulan = T.pinjam[j+1].tanggalpinjam.bulan
							T.pinjam[j].duedate.tanggal = T.pinjam[j+1].duedate.tanggal
							T.pinjam[j].duedate.bulan = T.pinjam[j+1].duedate.bulan
							T.pinjam[j].duedate.tahun = T.pinjam[j+1].duedate.tahun
							T.pinjam[j].dikembalikan.tanggal = T.pinjam[j+1].dikembalikan.tanggal
							T.pinjam[j].dikembalikan.bulan = T.pinjam[j+1].dikembalikan.bulan
							T.pinjam[j].dikembalikan.tahun = T.pinjam[j+1].dikembalikan.tahun
							T.pinjam[j].denda = T.pinjam[j+1].denda
							T.pinjam[j].status = T.pinjam[j+1].status

						}
						T.pinjam[*n-1].namaBuku = ""
						T.pinjam[*n-1].namaOrang = ""
						T.pinjam[*n-1].tanggalpinjam.tanggal = 0
						T.pinjam[*n-1].tanggalpinjam.bulan = 0
						T.pinjam[*n-1].tanggalpinjam.bulan = 0
						T.pinjam[*n-1].duedate.tanggal = 0
						T.pinjam[*n-1].duedate.bulan = 0
						T.pinjam[*n-1].duedate.tahun = 0
						T.pinjam[*n-1].dikembalikan.tanggal = 0
						T.pinjam[*n-1].dikembalikan.bulan = 0
						T.pinjam[*n-1].dikembalikan.tahun = 0
						T.pinjam[*n-1].denda = 0
						T.pinjam[*n-1].status = ""
						*n--

					} else if pilihanuser == "n" {
						fmt.Println("data tidak jadi dihapus")

					} else {
						fmt.Println("pilihan tidak ada pada menu, tolong masukkan antara y atau n")
						fmt.Scanln(&pilihanuser)
					}

				}

			} else {
				fmt.Println("data tidak ditemukan")
			}
		}
	}

}

// func hitungdenda(T perpus, tgl, bln, thn int, tglkembali, blnkembali, thnkembali int) float64 {
// 	denda := 0
// 	telat := -1
// 	i := 0
// 	found := false

// 	for i < T.nPinjamBuku && !found {
// 		if T.pinjam[i].duedate.tanggal == tgl && T.pinjam[i].duedate.bulan == bln && T.pinjam[i].duedate.tahun == thn {

// 			if tglkembali < T.pinjam[i].dikembalikan.tanggal || blnkembali < T.pinjam[i].dikembalikan.bulan || thnkembali < T.pinjam[i].dikembalikan.tahun {

// 				telat = 0
// 				if blnkembali == T.pinjam[i].dikembalikan.bulan && thnkembali == T.pinjam[i].dikembalikan.tahun {
// 					telat = tglkembali - T.pinjam[i].dikembalikan.tanggal
// 				} else if blnkembali == T.pinjam[i].dikembalikan.bulan {
// 					telat = (thnkembali - T.pinjam[i].dikembalikan.tahun) * 30
// 				} else {
// 					telat = (blnkembali - T.pinjam[i].dikembalikan.bulan) * 360
// 				}
// 				denda = telat * 500
// 			}
// 			found = true
// 		}
// 		i++
// 	}

// 	return float64(denda)
// }

func hitungdenda(T perpus, tgl, bln, thn, tglkembali, blnkembali, thnkembali int) float64 {
	//menghitung jumlah hari telat mengembalikan buku
	telat := 0.0

	for i := 0; i < T.nPinjamBuku; i++ {
		if T.pinjam[i].duedate.tanggal == tgl && T.pinjam[i].duedate.bulan == bln && T.pinjam[i].duedate.tahun == thn {
			if tglkembali > T.pinjam[i].duedate.tanggal || blnkembali > T.pinjam[i].duedate.bulan || thnkembali > T.pinjam[i].duedate.tahun {
				if blnkembali == T.pinjam[i].duedate.bulan && thnkembali == T.pinjam[i].duedate.tahun {
					telat = float64(tglkembali - T.pinjam[i].duedate.tanggal)
				} else if blnkembali == T.pinjam[i].duedate.bulan {
					telat = float64((thnkembali - T.pinjam[i].duedate.tahun) * 30)
				} else {
					telat = float64((blnkembali - T.pinjam[i].duedate.bulan) * 360)
				}

			}
		}
	}

	return telat
}
