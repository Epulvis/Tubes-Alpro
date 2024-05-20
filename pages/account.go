package pages

import (
	"bufio"
	"fmt"
	"github.com/inancgumus/screen"
	"os"
	"strings"
)

const dataFile = "accounts.json"

// Fungsi untuk mengubah slice string menjadi slice interface{}
func toInterfaceSlice(strs []string) []interface{} {
	ifaceSlice := make([]interface{}, len(strs))
	for i, v := range strs {
		ifaceSlice[i] = v
	}
	return ifaceSlice
}

// Fungsi untuk membuat header tabel
func printHeader(headers []string) {
	fmt.Printf("| %-20s | %-20s | %-10s | %-10s |\n", toInterfaceSlice(headers)...)
	fmt.Println(strings.Repeat("-", 72))
}

// Fungsi untuk mencetak baris data akun
func printAccountRow(account Account) {
	fmt.Printf("| %-20s | %-20s | %-10s | %-10s |\n",
		account.NamaPengguna, account.ChannelYouTube, account.Saldo, account.StatusYouTube)
	fmt.Println(strings.Repeat("-", 72))
	fmt.Println("Video:")
	for _, video := range account.Video {
		printVideoRow(video)
	}
	fmt.Println(strings.Repeat("-", 72))
}

// Fungsi untuk mencetak baris data video
func printVideoRow(video Video) {
	fmt.Printf("| %-30s | %-10s | %-20s |\n",
		video.Judul, video.Durasi, video.Views)
}

// Fungsi untuk menampilkan menu
func showMenu() {
	fmt.Println("\n--- Menu ---")
	fmt.Println("1. Tampilkan akun dan video")
	fmt.Println("2. Tambah akun")
	fmt.Println("3. Hapus akun")
	fmt.Println("4. Keluar")
	fmt.Print("Masukkan pilihan Anda: ")
}

// Fungsi untuk membersihkan layar terminal
func clearScreen() {
	screen.Clear()
	screen.MoveTopLeft()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Selamat datang di aplikasi pengelola akun YouTube!")
	fmt.Println("Aplikasi ini akan membantu Anda menampilkan informasi akun dan video.")
	fmt.Println("Silakan pilih opsi dari menu di bawah ini.")

	accounts, err := LoadAccounts(dataFile)
	if err != nil {
		fmt.Println("Gagal memuat data akun:", err)
		return
	}

	for {
		showMenu()
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // Menghilangkan newline dan spasi tambahan

		switch input {
		case "1":
			clearScreen()
			fmt.Println("Menampilkan akun dan video:\n")
			printHeader([]string{"Nama Pengguna", "Channel YouTube", "Saldo", "Status YouTube"})
			for _, account := range accounts {
				printAccountRow(account)
			}
			fmt.Println("\nTekan 'Enter' untuk kembali ke menu...")
			reader.ReadString('\n')
			clearScreen()
		case "2":
			clearScreen()
			fmt.Println("Tambah akun baru:")
			fmt.Print("Nama Pengguna: ")
			namaPengguna, _ := reader.ReadString('\n')
			namaPengguna = strings.TrimSpace(namaPengguna)

			fmt.Print("Channel YouTube: ")
			channelYouTube, _ := reader.ReadString('\n')
			channelYouTube = strings.TrimSpace(channelYouTube)

			fmt.Print("Saldo: ")
			saldo, _ := reader.ReadString('\n')
			saldo = strings.TrimSpace(saldo)

			fmt.Print("Status YouTube: ")
			statusYouTube, _ := reader.ReadString('\n')
			statusYouTube = strings.TrimSpace(statusYouTube)

			// Tambahkan video jika diinginkan
			var videos []Video
			for {
				fmt.Print("Judul Video (atau ketik 'selesai' untuk menyelesaikan): ")
				judul, _ := reader.ReadString('\n')
				judul = strings.TrimSpace(judul)
				if strings.ToLower(judul) == "selesai" {
					break
				}

				fmt.Print("Durasi: ")
				durasi, _ := reader.ReadString('\n')
				durasi = strings.TrimSpace(durasi)

				fmt.Print("Views: ")
				views, _ := reader.ReadString('\n')
				views = strings.TrimSpace(views)

				videos = append(videos, Video{Judul: judul, Durasi: durasi, Views: views})
			}

			accounts = append(accounts, Account{
				NamaPengguna:   namaPengguna,
				ChannelYouTube: channelYouTube,
				Saldo:          saldo,
				StatusYouTube:  statusYouTube,
				Video:          videos,
			})

			err := SaveAccounts(dataFile, accounts)
			if err != nil {
				fmt.Println("Gagal menyimpan data akun:", err)
			}

			fmt.Println("Akun berhasil ditambahkan!")
			fmt.Println("\nTekan 'Enter' untuk kembali ke menu...")
			reader.ReadString('\n')
			clearScreen()
		case "3":
			clearScreen()
			fmt.Println("Hapus akun:")
			fmt.Print("Masukkan nama pengguna yang ingin dihapus: ")
			namaPengguna, _ := reader.ReadString('\n')
			namaPengguna = strings.TrimSpace(namaPengguna)

			updatedAccounts, deleted := DeleteAccount(accounts, namaPengguna)
			if deleted {
				accounts = updatedAccounts
				err := SaveAccounts(dataFile, accounts)
				if err != nil {
					fmt.Println("Gagal menyimpan data akun:", err)
				}
				fmt.Println("Akun berhasil dihapus!")
			} else {
				fmt.Println("Akun tidak ditemukan.")
			}

			fmt.Println("\nTekan 'Enter' untuk kembali ke menu...")
			reader.ReadString('\n')
			clearScreen()
		case "4":
			fmt.Println("Terima kasih telah menggunakan aplikasi ini. Selamat tinggal!")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}
