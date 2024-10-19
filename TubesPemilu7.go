package main

import (
	"fmt"
)

const NMAX int = 10

type calon struct {
	namaCalon, partai   string
	noUrut, jumlahSuara int
	thresehold          float64
}

type pemilih struct {
	namaPemilih string
	pilihan     int
}

type tabCaleg [NMAX]calon
type tabPemilih [NMAX]pemilih

func main() {
	var X tabCaleg
	var Z tabPemilih
	var pilihan, petugas, pemilih int
	var indeks int
	var rename string
	var jam float64
	var Y int = 4 //Untuk tabCaleg
	var W int = 0 // untuk tabPemilih
	var flag string
	var nama string
	var delete string

	var ulangi1 bool = true
	var ulangi2 bool = true
	var ulangi3 bool = true
	dataAwalCalon(&X)
	for ulangi1 {
		menuTampilanAwal()
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			fmt.Print("Masukkan Waktu Kondisi Anda Sekarang: ")
			fmt.Scan(&jam)
			ulangi2 = true
			for ulangi2 {
				if jam >= 07.00 && jam <= 17.00 {
					menuTampilanPemilihOpsi()
					fmt.Scan(&pemilih)
					if pemilih == 1 {
						pemilihanSuara(&X, &Z, Y, &W)
					} else if pemilih == 2 {
						sortingDataCalon(&X, Y)
						printDataCalonn(&X, Y)
					} else {
						fmt.Println("<================  Terima Kasih ===================>")
						ulangi2 = false
					}
				} else {
					menuTampilanPemilihOpsi2()
					fmt.Scan(&pemilih)
					if pemilih == 1 {
						sortingDataCalon(&X, Y)
						printDataCalonn(&X, Y)
					} else {
						fmt.Println("<================  Terima Kasih ===================>")
						ulangi2 = false
					}
				}
			}
		} else if pilihan == 2 {
			ulangi3 = true
			for ulangi3 {
				menuTampilanPetugas()
				fmt.Scan(&petugas)
				if petugas == 1 {
					tambahDataCalon(&X, &Y)
				} else if petugas == 2 {
					fmt.Print("Masukkan Nama Calon yang Ingin Dicari: ")
					fmt.Scan(&rename)
					renameDataCalon(&X, Y, rename)
				} else if petugas == 3 {
					fmt.Print("Masukkan Nama Calon yang Ingin Dihapus: ")
					fmt.Scan(&delete)
					deleteDataCalon(&X, &Y, delete)
				} else if petugas == 4 {
					sortingDataCalon(&X, Y)
					printDataCalonn(&X, Y)
				} else if petugas == 5 {
					menuCariDataCalon()
					fmt.Scan(&flag)
					if flag == "1" {
						indeks = searchByNama(X, Y, flag)
						if indeks == -1 {
							fmt.Print("===== Maaf Nama Calon Tidak Ditemukan ====")
						} else {
							fmt.Println("<===== Data Calon Berhasil Ditemukan ! =====>")
							fmt.Println("Nomor Urut.", X[indeks].noUrut, "Nama Calon:", X[indeks].namaCalon, "Thresehold:", X[indeks].thresehold, "Partai:", X[indeks].partai)
						}
					} else if flag == "2" {
						indeks = searchByPartai(X, Y, flag)
						if indeks == -1 {
							fmt.Print("===== Maaf Nama Partai Tidak Ditemukan ====")
						} else {
							fmt.Println("<==== Data Partai Berhasil Ditemukan ! =====>")
							fmt.Println("Nomor Urut.", X[indeks].noUrut, "Nama Calon:", X[indeks].namaCalon, "Thresehold:", X[indeks].thresehold, "Partai:", X[indeks].partai)
						}
					} else if flag == "3" {
						indeks = searchByPemilih(Z, W)
						if indeks == -1 {
							fmt.Print("===== Maaf Data Calon Tidak Ditemukan ====")
						} else {
							pil := Z[indeks].pilihan
							searchByPemilih2(&X, Y, pil)
						}
					}
				} else if petugas == 6 {
					searchingNamaPemilih(&Z, W, nama)
				} else {
					fmt.Println("<====== Terima Kasih Atas Partisipasi Anda ======>")
					ulangi3 = false
				}
			}

		} else if pilihan == 0 {
			fmt.Println("<====== Terima Kasih Atas Partisipasi Anda ======>")
			fmt.Println("<======            Sampai Jumpa            ======>")
			ulangi1 = false
		}
	}
}

func menuTampilanAwal() {
	fmt.Println("")
	fmt.Println("<================================================>")
	fmt.Println("<===========| Aplikasi Pemilihan Umum |==========>")
	fmt.Println("<================================================>")
	fmt.Println("<===========|       MENU UTAMA        |==========>")
	fmt.Println("<================================================>")
	fmt.Println("<===========|       1. Pemilih        |==========>")
	fmt.Println("<===========|       2. Petugas        |==========>")
	fmt.Println("<===========|       0. Exit           |==========>")
	fmt.Println("<================================================>")
	fmt.Print("Pilih Opsi Anda: ")
}

func menuTampilanPemilihWaktu() {
	fmt.Println("")
	fmt.Println("<================================================>")
	fmt.Println("<==========| Selamat Datang Pemilih   |==========>")
	fmt.Println("<================================================>")
	fmt.Println("<==========| 1. Melakukan Pemilihan  |===========>")
	fmt.Println("<==========| 2. Tampilkan Data Calon |===========>")
	fmt.Println("<==========| 0. Exit                 |===========>")
	fmt.Println("<================================================>")
	fmt.Print("Masukkan Waktu Kondisi Anda Sekarang: ")
}

func menuTampilanPemilihOpsi() {
	fmt.Println("")
	fmt.Println("<================================================>")
	fmt.Println("<=========| Selamat Datang Pemilih  |============>")
	fmt.Println("<================================================>")
	fmt.Println("<=========|   MENU UTAMA PEMILIH    |============>")
	fmt.Println("<================================================>")
	fmt.Println("<=========| 1. Melakukan Pemilihan  |============>")
	fmt.Println("<=========| 2. Tampilkan Data Calon |============>")
	fmt.Println("<=========| 0. Exit                 |============>")
	fmt.Println("<================================================>")
	fmt.Print("Pilih Opsi Anda: ")
}

func menuTampilanPemilihOpsi2() {
	fmt.Println("")
	fmt.Println("<===============================================>")
	fmt.Println("       Maaf Waktu Pemilihan Telah berakhir       ")
	fmt.Println("<===============================================>")
	fmt.Println("<========|   1. Tampilkan Data Calon     |======>")
	fmt.Println("<========|   0. Exit                     |======>")
	fmt.Println("<===============================================>")
	fmt.Print("Pilih Opsi Anda: ")
}

func menuTampilanPetugas() {
	fmt.Println("")
	fmt.Println("<==============================================>")
	fmt.Println("<=====| Selamat Datang Petugas KPU 2023 |======>")
	fmt.Println("<==============================================>")
	fmt.Println("<=====|        MENU UTAMA PETUGAS       |======>")
	fmt.Println("<==============================================>")
	fmt.Println("<=====|    1. Tambah Calon              |======>")
	fmt.Println("<=====|    2. Ganti Data Calon          |======>")
	fmt.Println("<=====|    3. Hapus Data Calon          |======>")
	fmt.Println("<=====|    4. Tampilkan Data Calon      |======>")
	fmt.Println("<=====|    5. Cari Data Calon           |======>")
	fmt.Println("<=====|    6. Cari Data Pemilih         |======>")
	fmt.Println("<=====|    0. Exit                      |======>")
	fmt.Println("<==============================================>")
	fmt.Print("Apa Yang Ingin Anda Lakukan: ")
}
func menuCariDataCalon() {
	fmt.Println("")
	fmt.Println("<===============================================>")
	fmt.Println("<=====|      Selamat Datang Petugas       |=====>")
	fmt.Println("<===============================================>")
	fmt.Println("<=====|   Cari Data Calon Berdasarkan ?  |======>")
	fmt.Println("<===============================================>")
	fmt.Println("<=====|    1.    Nama Calon              |======>")
	fmt.Println("<=====|    2.    Partai Calon            |======>")
	fmt.Println("<=====|    3.    Nama Pemilih            |======>")
	fmt.Println("<===============================================>")
	fmt.Print("Pilihan Anda: ")
}

func dataAwalCalon(T *tabCaleg) {
	T[0].namaCalon = "Zaidan"
	T[1].namaCalon = "Bagas"
	T[2].namaCalon = "Hamdi"
	T[3].namaCalon = "Dafa"

	T[0].noUrut = 1
	T[1].noUrut = 2
	T[2].noUrut = 3
	T[3].noUrut = 4

	T[0].thresehold = 3.0
	T[1].thresehold = 1.5
	T[2].thresehold = 3.0
	T[3].thresehold = 2.5

	T[0].partai = "PPP"
	T[1].partai = "PDIP"
	T[2].partai = "Golkar"
	T[3].partai = "Demokrat"
}

func printDataCalonn(T *tabCaleg, N int) {
	fmt.Println("")
	fmt.Println("|=================================================================|")
	fmt.Println("| No. |         Nama Calon         |     Treshold    |    Partai  |")
	fmt.Println("|=================================================================|")
	for i := 0; i < N; i++ {
		fmt.Printf("|  %d  | %-26s | %-15.2f | %-10s | \n", T[i].noUrut, T[i].namaCalon, T[i].thresehold, T[i].partai)
		fmt.Println("|=================================================================|")
	}
}

func tambahDataCalon(T *tabCaleg, N *int) {
	var loop bool = true
	for loop {
		fmt.Println("")
		fmt.Print("Masukkan Nama Calon       : ")
		fmt.Scan(&T[*N].namaCalon)
		fmt.Print("Masukkan Nomor Urut Calon : ")
		fmt.Scan(&T[*N].noUrut)
		fmt.Print("Masukkan thresehold Calon : ")
		fmt.Scan(&T[*N].thresehold)
		fmt.Print("Masukkan Nama Partai      : ")
		fmt.Scan(&T[*N].partai)
		*N++
		var input string = "YA"
		fmt.Println("Tambah Data Calon Lagi? (YA / TIDAK) ")
		fmt.Scan(&input)
		if input == "TIDAK" {
			loop = false
			fmt.Println("<===== Data Calon Baru Berhasil Disimpan!  =====>")
		}
	}
}

// Dipakai untuk fungsi rename
func searchingCalon2(T tabCaleg, N int, rename string) int {
	var i, indeks int

	i = 0
	indeks = -1
	for i < N && indeks == -1 {
		if T[i].namaCalon == rename {
			indeks = i
		}
		i = i + 1
	}
	return indeks
}

// Dipakai untuk fungsi delete
func searchingCalon3(T tabCaleg, N int, delete string) int {
	var i, found int

	i = 0
	found = -1
	for i < N && found == -1 {
		if T[i].namaCalon == delete {
			found = i
		}
		i = i + 1
	}
	return found
}

func searchByNama(T tabCaleg, N int, flag string) int {
	var i, indeks int

	i = 0
	indeks = -1
	fmt.Print("Masukkan Nama Calon yang ingin dicari: ")
	fmt.Scan(&flag)
	for i < N && indeks == -1 {
		if T[i].namaCalon == flag {
			indeks = i
		}
		i++
	}
	return indeks
}

func searchByPartai(T tabCaleg, N int, flag string) int {
	var i, indeks int

	i = 0
	indeks = -1
	fmt.Print("Masukkan Nama Partai yang ingin dicari: ")
	fmt.Scan(&flag)
	for i < N && indeks == -1 {
		if T[i].partai == flag {
			indeks = i
		}
		i++
	}
	return indeks
}

func searchByPemilih(T tabPemilih, N int) int {
	var i, indeks int
	var inputan string

	i = 0
	indeks = -1
	fmt.Print("Masukkan Nama Pemilih: ")
	fmt.Scan(&inputan)
	for i < N && indeks == -1 {
		if T[i].namaPemilih == inputan {
			indeks = i
		}
		i++
	}
	return indeks
}

// Dipakai untuk melanjutkan pencarian dengan menggunakan pemilih dan dipanggil di func main
func searchByPemilih2(T *tabCaleg, N, X int) {
	var i, indeks int

	i = 0
	indeks = -1
	for i < N && indeks == -1 {
		if T[i].noUrut == X {
			indeks = i
			fmt.Println("Nomor Urut.", T[indeks].noUrut, "Nama Calon:", T[indeks].namaCalon, "Thresehold:", T[indeks].thresehold, "Partai:", T[indeks].partai)
		}
		i++
	}
}

func pemilihanSuara(T *tabCaleg, S *tabPemilih, N int, nPemilih *int) {
	var indeks int = -1
	var i, inputSuara int

	fmt.Print("Masukkan No Urut Calon Untuk Memilih Calon: ")
	fmt.Scan(&inputSuara)
	fmt.Print("Masukkan Nama Anda: ")
	fmt.Scan(&S[*nPemilih].namaPemilih)
	for indeks == -1 && inputSuara < N {
		if T[i].noUrut == inputSuara {
			indeks = i
			if T[indeks].thresehold >= 3 {
				S[*nPemilih].pilihan = inputSuara
				*nPemilih++
				T[indeks].jumlahSuara++
				fmt.Println("<======== Terima Kasih Telah Memilih Calon ========>")
			} else {
				fmt.Println("<= Maaf, Calon yang Anda Pilih Tidak Memenuhi Threshold =>")
			}
		}
		i++
	}
}

func deleteDataCalon(T *tabCaleg, N *int, delete string) {
	var found int

	found = searchingCalon3(*T, *N, delete)
	if found == -1 {
		fmt.Println("Data Calon yang Anda Masukkan Tidak Sesuai!")
	} else {
		for found <= *N-2 {
			T[found].noUrut = T[found+1].noUrut
			T[found].namaCalon = T[found+1].namaCalon
			T[found].thresehold = T[found+1].thresehold
			T[found].partai = T[found+1].partai
			found++
		}
		fmt.Println("Data Calon Berhasil Dihapus")
	}
	*N--
}

func renameDataCalon(T *tabCaleg, N int, rename string) {
	var indeks int

	indeks = searchingCalon2(*T, N, rename)
	if indeks == -1 {
		fmt.Println("Nama Calon Tidak Ditemukan")
	} else {
		fmt.Print("Ubah Nama Calon Menjadi: ")
		fmt.Scan(&T[indeks].namaCalon)
	}
}

func searchingNamaPemilih(T *tabPemilih, N int, nama string) {
	var i, indeks int

	fmt.Print("Masukkan Nama Pemilih yang Ingin Dicari: ")
	fmt.Scan(&nama)
	i = 0
	indeks = -1
	for i < N && indeks == -1 {
		if T[i].namaPemilih == nama {
			indeks = i
			fmt.Println("Nama Pemilih yang Anda Cari:", T[indeks].namaPemilih)
		}
		i++
	}
	if indeks == -1 {
		fmt.Println("Nama yang Anda Cari Tidak Ada")
	}
}

func sortingDataCalon(T *tabCaleg, N int) {
	var i, pass int
	var temp calon
	var flag int

	fmt.Println("")
	fmt.Println("<===============================================>")
	fmt.Println("<=====|   Cari Data Calon Berdasarkan ?  |======>")
	fmt.Println("<===============================================>")
	fmt.Println("<=====|    1.    Perolehan Suara         |======>")
	fmt.Println("<=====|    2.    Partai Calon            |======>")
	fmt.Println("<=====|    3.    Nama Calon              |======>")
	fmt.Println("<===============================================>")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&flag)
	pass = 1
	if flag == 1 {
		for pass < N {
			i = pass
			temp = T[pass]
			for i > 0 && temp.jumlahSuara > T[i-1].jumlahSuara {
				T[i] = T[i-1]
				i--
			}
			T[i] = temp
			pass++
		}
	}
	if flag == 2 {
		for pass < N {
			i = pass
			temp = T[pass]
			for i > 0 && temp.partai < T[i-1].partai {
				T[i] = T[i-1]
				i--
			}
			T[i] = temp
			pass++
		}
	}
	if flag == 3 {
		for pass < N {
			i = pass
			temp = T[pass]
			for i > 0 && temp.namaCalon < T[i-1].namaCalon {
				T[i] = T[i-1]
				i--
			}
			T[i] = temp
			pass++
		}
	}
}
