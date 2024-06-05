package app

import (
	"YouTubeAdSenseAPP/internal/component"
	"fmt"
)

func ShowHomepages() {
	var input string

	fmt.Println("Selamat datang di aplikasi pengelola akun YouTube!")
	fmt.Println("Aplikasi ini akan membantu Anda menampilkan informasi akun dan video.")
	fmt.Println("Silakan pilih opsi dari menu di bawah ini.")

	for {
		showMenu()

		fmt.Scanln(&input)

		if input == "1" {
			showDisplaysAllAccounts()
		} else if input == "2" {
			showDisplaysAccounts()
		} else if input == "3" {
			showAddAccount()
		} else if input == "4" {
			showDeleteAccount()
		} else if input == "5" {
			return
		} else {
			component.ClearScreen()
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}

	}
}

func showMenu() {
	fmt.Println("\n--- Menu ---")
	fmt.Println("1. Tampilkan semua akun")
	fmt.Println("2. Tampilkan akun")
	fmt.Println("3. Tambah akun dan video")
	fmt.Println("4. Hapus akun")
	fmt.Println("5. Keluar")
	fmt.Print("Masukkan pilihan Anda: ")
}
