package main

import "fmt"

const N = 1000

type login struct {
	usn   string
	pw    string
	email string
}

type question struct {
	ask     string
	tag     string
	title   string
	answers [N]string
}

type tabAnswer [N]question
type tabAsk [N]question
type tabAkun [N]login

func user() {
	fmt.Println("----------------------------------")
	fmt.Println(" USER PASIEN / DOKTER / GUEST")
	fmt.Println("----------------------------------")
	fmt.Println("1. Pasien")
	fmt.Println("2. Dokter")
	fmt.Println("3. Guest")
	fmt.Println("4. Keluar Aplikasi")
	fmt.Println("----------------------------------")
}

func backToUserMenu() {
	fmt.Println("\n----------------------------------")
	fmt.Println("Kembali ke menu pemilihan user...")
	fmt.Println("----------------------------------")
	user()
}

func availaccount() {
	fmt.Println("----------------------------------")
	fmt.Println("Sudah daftar / belum")
	fmt.Println("----------------------------------")
	fmt.Println("1. Sudah")
	fmt.Println("2. Belum")
	fmt.Println("3. Kembali")
	fmt.Println("----------------------------------")
}

func regpasien(id *tabAkun, index *int) {
	for {
		fmt.Println("----------------------------------")
		fmt.Println("  SILAHKAN DAFTAR SEBAGAI PASIEN")
		fmt.Println("----------------------------------")
		var usn, pw, email string
		fmt.Print("Masukkan username: ")
		fmt.Scan(&usn)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&pw)
		fmt.Print("Masukkan email: ")
		fmt.Scan(&email)

		usernameTerdaftar := false
		for i := 0; i < *index; i++ {
			if (*id)[i].usn == usn {
				usernameTerdaftar = true
				break
			}
		}

		if usernameTerdaftar {
			fmt.Println("Username sudah terdaftar. Silakan gunakan username lain.")
		} else {
			// Tambahkan akun baru jika username belum terdaftar
			(*id)[*index] = login{usn: usn, pw: pw, email: email}
			*index++

			fmt.Println("-----------------------------------")
			fmt.Println("  BERHASIL, SELAMAT DATANG", usn)
			break
		}
	}
}
func logpasien(id tabAkun, index int) {
	for {
		fmt.Println("----------------------------------")
		fmt.Println("      SILAHKAN LOGIN PASIEN")
		fmt.Println("----------------------------------")
		var usn, pw string
		fmt.Print("Masukkan username: ")
		fmt.Scan(&usn)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&pw)

		valid := false
		for i := 0; i < index; i++ {
			if usn == id[i].usn && pw == id[i].pw {
				valid = true
				fmt.Println("-----------------------------------")
				fmt.Println("  BERHASIL, WELCOME BACK", usn)
				return // Menghentikan fungsi setelah login berhasil
			}
		}

		if !valid {
			fmt.Println("-----------------------------------")
			fmt.Println("  LOGIN GAGAL, USERNAME ATAU PASSWORD SALAH")
			fmt.Println("-----------------------------------")
			fmt.Println("1. Coba lagi")
			fmt.Println("2. Kembali ke halaman sebelumnya")
			fmt.Print("Pilihan Anda: ")
			var pilihan int
			fmt.Scan(&pilihan)

			if pilihan == 2 {
				return
			}
		}
	}
}
func regdokter(id *tabAkun, index *int) {
	for {
		fmt.Println("----------------------------------")
		fmt.Println("  SILAHKAN DAFTAR SEBAGAI DOKTER")
		fmt.Println("----------------------------------")
		var usn, pw, email string
		fmt.Print("Masukkan username: ")
		fmt.Scan(&usn)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&pw)
		fmt.Print("Masukkan email: ")
		fmt.Scan(&email)

		usernameTerdaftar := false
		for i := 0; i < *index; i++ {
			if (*id)[i].usn == usn {
				usernameTerdaftar = true
				break
			}
		}

		if usernameTerdaftar {
			fmt.Println("Username sudah terdaftar. Silakan gunakan username lain.")
		} else {
			// Tambahkan akun baru jika username belum terdaftar
			(*id)[*index] = login{usn: usn, pw: pw, email: email}
			*index++

			fmt.Println("-----------------------------------")
			fmt.Println("  BERHASIL, SELAMAT DATANG", usn)
			break
		}
	}
}

func logdokter(id tabAkun, index int) {
	for {
		fmt.Println("----------------------------------")
		fmt.Println("         LOGIN DOKTER")
		fmt.Println("----------------------------------")
		var usn, pw string
		fmt.Print("Masukkan username: ")
		fmt.Scan(&usn)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&pw)

		valid := false
		for i := 0; i < index; i++ {
			if usn == id[i].usn && pw == id[i].pw {
				valid = true
				fmt.Println("-----------------------------------")
				fmt.Println("  BERHASIL, WELCOME BACK DOK!", usn)
				return // Menghentikan fungsi setelah login berhasil
			}
		}

		if !valid {
			fmt.Println("-----------------------------------")
			fmt.Println("  LOGIN GAGAL, USERNAME ATAU PASSWORD SALAH")
			fmt.Println("-----------------------------------")
			fmt.Println("1. Coba lagi")
			fmt.Println("2. Kembali ke menu pendaftaran/login")
			fmt.Print("Pilihan Anda: ")
			var pilihan int
			fmt.Scan(&pilihan)

			if pilihan == 2 {
				availaccount() // Kembali ke menu pendaftaran/login
				return
			}
		}
	}
}
func menuUser() {
	fmt.Println("----------------------------------")
	fmt.Println(" SELAMAT DATANG DI GHAZY KONSUL ")
	fmt.Println(" SILAHKAN PILIH MENU")
	fmt.Println("----------------------------------")
	fmt.Println("1. Konsultasi Kesehatan")
	fmt.Println("2. Forum Konsultasi")
	fmt.Println("3. Diskusi Pertanyaan")
	fmt.Println("4. Cari Pertanyaan berdasarkan tag")
	fmt.Println("5. Tampilkan Tag Terurut berdasarkan jumlah Pertanyaan")
	fmt.Println("6. Log out")
	fmt.Println("----------------------------------")
}

func menuGuest() {
	fmt.Println("----------------------------------")
	fmt.Println(" SELAMAT DATANG DI GHAZY KONSUL ")
	fmt.Println(" SILAHKAN PILIH MENU")
	fmt.Println("----------------------------------")
	fmt.Println("1. Forumm Konsultasi")
	fmt.Println("2. Cari Pertanyaan berdasarkan tag")
	fmt.Println("3. Log out")
	fmt.Println("----------------------------------")
}

func confirm() int {
	var option int
	fmt.Println("-----------------------------------")
	fmt.Println(" Ingin mengakses halaman ini? ")
	fmt.Println("-----------------------------------")
	fmt.Println("1. YA")
	fmt.Println("2. Tidak")
	fmt.Println("-----------------------------------")
	fmt.Scan(&option)
	fmt.Println("-----------------------------------")
	return option
}

func askquestion(Ask *tabAsk, index int) {
	fmt.Println(" SILAHKAN POSTING KONSULTASI PENYAKITMU")
	fmt.Println("-----------------------------------")
	fmt.Print(" MASUKKAN JUDUL : ")
	fmt.Scan(&(*Ask)[index].title)
	fmt.Print(" APA PERTANYAAN MU?: ")
	fmt.Scan(&(*Ask)[index].ask)
	fmt.Print(" MASUKKAN TAG TERKAIT PERTANYAANMU, CONTOH #covid19: ")
	fmt.Scan(&(*Ask)[index].tag)
	fmt.Println("-----------------------------------")
	fmt.Println(" BERHASIL, SILAHKAN TUNGGU HINGGA SESEORANG MERESPON")
}

func historykonsul(id tabAkun, Ask tabAsk, Ans tabAnswer, index int) {
	for {
		fmt.Println(" RIWAYAT PERTANYAAN & JAWABAN")
		fmt.Println("-----------------------------------")
		noPertanyaan := true
		for i := 0; i < index; i++ {
			if Ask[i].title != "" {
				fmt.Printf("%d. Dari: %s, Judul: %s, Pertanyaan: %s, Tag: %s\n", i+1, id[i].usn, Ask[i].title, Ask[i].ask, Ask[i].tag)
				noDiskusi := true
				for j := 0; j < N; j++ {
					if Ans[i].answers[j] != "" {
						fmt.Println("   Jawaban dari ", id[j].usn, ":", Ans[i].answers[j])
						noDiskusi = false
					}
				}
				if noDiskusi {
					fmt.Println("   Belum ada Diskusi apapun")
				}
				noPertanyaan = false
			}
		}

		if noPertanyaan {
			fmt.Println(" Belum Ada pertanyaan apapun ")
		}

		fmt.Println("-----------------------------------")
		fmt.Println("1. Lihat lagi riwayat konsultasi")
		fmt.Println("2. Kembali ke menu sebelumnya")
		fmt.Println("-----------------------------------")
		var choice int
		fmt.Scan(&choice)
		if choice == 2 {
			fmt.Println("-----------------------------------")
			fmt.Println(" Kembali ke menu sebelumnya ")
			fmt.Println("-----------------------------------")
			return
		}
	}
}

func cariPertanyaan(id tabAkun, Ask *tabAsk, Ans *tabAnswer, index int) {
	for {
		var tag string
		fmt.Println(" CARI PERTANYAAN BERDASARKAN TAG")
		fmt.Println("-----------------------------------")
		fmt.Scan(&tag)

		found := false

		fmt.Println("-----------------------------------")
		fmt.Println(" BERIKUT ADALAH PERTANYAAN YANG SESUAI DENGAN TAG YANG DICARI ")
		fmt.Println("-----------------------------------")

		// Gunakan fungsi caritag untuk menemukan indeks
		for i := 0; i < index; i++ {
			if (*Ask)[i].tag == tag && (*Ask)[i].title != "" {
				found = true
				fmt.Printf("%d. Dari: %s, Judul: %s, Pertanyaan: %s, Tag: %s\n", i+1, id[i].usn, (*Ask)[i].title, (*Ask)[i].ask, (*Ask)[i].tag)

				noDiskusi := true
				for j := 0; j < N; j++ {
					if (*Ans)[i].answers[j] != "" {
						fmt.Println("   Jawaban dari ", id[j].usn, ":", (*Ans)[i].answers[j])
						noDiskusi = false
					}
				}
				if noDiskusi {
					fmt.Println("   Belum ada Diskusi apapun")
				}
			}
		}

		if !found {
			fmt.Println("Tidak ada pertanyaan dengan tag yang sesuai.")
		}

		fmt.Println("-----------------------------------")
		fmt.Println("Apakah Anda sudah selesai mencari?")
		fmt.Println("1. Ya")
		fmt.Println("2. Tidak")
		var choice int
		fmt.Scan(&choice)
		if choice == 1 {
			fmt.Println("-----------------------------------")
			fmt.Println(" Kembali ke menu sebelumnya ")
			fmt.Println("-----------------------------------")
			return
		}
	}
}

func diskusikonsul(id tabAkun, Ask *tabAsk, Ans *tabAnswer, index int) {
	for {
		var tag string
		fmt.Println(" CARI TAG YANG INGIN KAMU DISKUSIKAN")
		fmt.Println("-----------------------------------")
		fmt.Scan(&tag)

		found := false

		fmt.Println("-----------------------------------")
		fmt.Println(" BERIKUT ADALAH PERTANYAAN YANG SESUAI DENGAN TAG YANG DICARI ")
		fmt.Println("-----------------------------------")

		// Gunakan fungsi caritag untuk menemukan indeks
		for i := 0; i < index; i++ {
			if (*Ask)[i].tag == tag && (*Ask)[i].title != "" {
				if !found {
					found = true
				}
				fmt.Printf("%d. Dari: %s, Judul: %s, Pertanyaan: %s, Tag: %s\n", i+1, id[i].usn, (*Ask)[i].title, (*Ask)[i].ask, (*Ask)[i].tag)

				noDiskusi := true
				for j := 0; j < N; j++ {
					if (*Ans)[i].answers[j] != "" {
						fmt.Println("   Jawaban dari ", id[j].usn, ":", (*Ans)[i].answers[j])
						noDiskusi = false
					}
				}
				if noDiskusi {
					fmt.Println("   Belum ada Diskusi apapun")
				}
			}
		}

		if !found {
			fmt.Println("-----------------------------------")
			fmt.Println("Tidak ada pertanyaan dengan tag yang sesuai.")
			fmt.Println("1. Masukkan tag lagi")
			fmt.Println("2. Kembali")
			var choice int
			fmt.Scan(&choice)
			if choice == 2 {
				fmt.Println("-----------------------------------")
				fmt.Println(" Kembali ke menu sebelumnya ")
				fmt.Println("-----------------------------------")
				return
			}
			continue
		}
		fmt.Println("-----------------------------------")
		fmt.Println("PILIH NOMOR PERTANYAAN YANG INGIN DIJAWAB")
		fmt.Println("-----------------------------------")
		var choice int
		fmt.Scan(&choice)
		if choice > 0 && choice <= index {
			fmt.Println("-----------------------------------")
			fmt.Println(" SILAHKAN DISKUSIKAN JAWABAN MU DIBAWAH ")
			fmt.Println("-----------------------------------")
			var answer string
			fmt.Scan(&answer)
			for k := 0; k < N; k++ {
				if (*Ans)[choice-1].answers[k] == "" {
					(*Ans)[choice-1].answers[k] = answer
					break
				}
			}
			fmt.Println("-----------------------------------")
			fmt.Println(" BERHASIL, MERESPON KONSULTASI")
			fmt.Println("-----------------------------------")
			fmt.Println("Apakah Anda sudah selesai berdiskusi?")
			fmt.Println("1. Ya")
			fmt.Println("2. Tidak")
			fmt.Scan(&choice)
			if choice == 1 {
				fmt.Println("-----------------------------------")
				fmt.Println(" Kembali ke menu sebelumnya ")
				fmt.Println("-----------------------------------")
				return
			}
		} else {
			fmt.Println("Nomor pertanyaan tidak valid.")
			fmt.Println("1. Pilih nomor pertanyaan lagi")
			fmt.Println("2. Kembali")
			fmt.Scan(&choice)
			if choice == 2 {
				fmt.Println("-----------------------------------")
				fmt.Println(" Kembali ke menu sebelumnya ")
				fmt.Println("-----------------------------------")
				return
			}
		}
	}
}

func caritag(Ask tabAsk, index int, x string) int {
	var found int = -1
	var j int = 0
	for j < index && found == -1 {
		if Ask[j].tag == x {
			found = j
		}
		j = j + 1
	}
	return found
}
func sortTagsByFrequency(Ask tabAsk, index int) {
	var tags [N]string
	var frequencies [N]int
	var tagCount int

	// Hitung frekuensi tag
	for i := 0; i < index; i++ {
		if Ask[i].tag != "" {
			found := false
			for j := 0; j < tagCount; j++ {
				if tags[j] == Ask[i].tag {
					frequencies[j]++
					found = true
					break
				}
			}
			if !found {
				tags[tagCount] = Ask[i].tag
				frequencies[tagCount] = 1
				tagCount++
			}
		}
	}

	// Bubble sort berdasarkan frekuensi
	for i := 0; i < tagCount-1; i++ {
		for j := 0; j < tagCount-i-1; j++ {
			if frequencies[j] < frequencies[j+1] {
				// Tukar frekuensi
				tempFreq := frequencies[j]
				frequencies[j] = frequencies[j+1]
				frequencies[j+1] = tempFreq

				// Tukar tag
				tempTag := tags[j]
				tags[j] = tags[j+1]
				tags[j+1] = tempTag
			}
		}
	}

	// Tampilkan hasil
	fmt.Println("-----------------------------------")
	fmt.Println(" TAG TERURUT BERDASARKAN JUMLAH PERTANYAAN ")
	fmt.Println("-----------------------------------")
	for i := 0; i < tagCount; i++ {
		fmt.Printf("%d. Tag: %s, Jumlah Pertanyaan: %d\n", i+1, tags[i], frequencies[i])
	}
}

func main() {
	var id tabAkun
	var option int
	var pertanyaan tabAsk
	var jawab tabAnswer
	var questionIndex int
	var pasienindex, dokterindex int

	for {
		user()
		fmt.Scan(&option)
		if option == 1 {
			availaccount()
			fmt.Scan(&option)
			if option == 2 {
				regpasien(&id, &pasienindex)
			} else if option == 1 {
				logpasien(id, pasienindex)
			} else if option == 3 {
				continue
			}
			for {
				menuUser()
				fmt.Scan(&option)
				if option == 1 {
					if confirm() == 1 {
						askquestion(&pertanyaan, questionIndex)
						questionIndex++
					}
				} else if option == 2 {
					if confirm() == 1 {
						historykonsul(id, pertanyaan, jawab, questionIndex)
					}
				} else if option == 3 {
					if confirm() == 1 {
						diskusikonsul(id, &pertanyaan, &jawab, questionIndex)
					}
				} else if option == 4 {
					if confirm() == 1 {
						cariPertanyaan(id, &pertanyaan, &jawab, questionIndex)
					}
				} else if option == 5 {
					if confirm() == 1 {
						sortTagsByFrequency(pertanyaan, questionIndex)
					}
				} else if option == 6 {
					if confirm() == 1 {
						break
					}
				}
			}
		} else if option == 2 {
			availaccount()
			fmt.Scan(&option)
			if option == 2 {
				regdokter(&id, &dokterindex)
			} else if option == 1 {
				logdokter(id, dokterindex)
			} else if option == 3 {
				continue
			}
			for {
				menuUser()
				fmt.Scan(&option)
				if option == 1 {
					if confirm() == 1 {
						askquestion(&pertanyaan, questionIndex)
						questionIndex++
					}
				} else if option == 2 {
					if confirm() == 1 {
						historykonsul(id, pertanyaan, jawab, questionIndex)
					}
				} else if option == 3 {
					if confirm() == 1 {
						diskusikonsul(id, &pertanyaan, &jawab, questionIndex)
					}
				} else if option == 4 {
					if confirm() == 1 {
						cariPertanyaan(id, &pertanyaan, &jawab, questionIndex)
					}
				} else if option == 5 {
					if confirm() == 1 {
						sortTagsByFrequency(pertanyaan, questionIndex)
					}
				} else if option == 6 {
					if confirm() == 1 {
						break
					}
				}
			}
		} else if option == 3 {
			for {
				menuGuest()
				fmt.Scan(&option)
				if option == 1 {
					if confirm() == 1 {
						historykonsul(id, pertanyaan, jawab, questionIndex)
					}
				} else if option == 2 {
					if confirm() == 1 {
						cariPertanyaan(id, &pertanyaan, &jawab, questionIndex)
					}
				} else if option == 3 {
					if confirm() == 1 {
						break
					}
				}
			}
		} else {
			break
		}
	}
}
