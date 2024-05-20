package pages

import (
	"YouTubeAdSenseAPP/component"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ShowHomepages() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Selamat datang di aplikasi pengelola akun YouTube!")
	fmt.Println("Aplikasi ini akan membantu Anda menampilkan informasi akun dan video.")
	fmt.Println("Silakan pilih opsi dari menu di bawah ini.")

	for {
		showMenu()
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			showDisplaysAccounts()
		case "2":
			//showLogin()
		case "3":
			showAddAccount()
		case "4":
			//showDeleteAccount()
		case "5":
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
			component.ClearScreen()
		}
	}

}

func showMenu() {
	fmt.Println("\n--- Menu ---")
	fmt.Println("1. Tampilkan akun")
	fmt.Println("2. Login")
	fmt.Println("3. Tambah akun dan video")
	fmt.Println("4. Hapus akun")
	fmt.Println("5. Keluar")
	fmt.Print("Masukkan pilihan Anda: ")
}
