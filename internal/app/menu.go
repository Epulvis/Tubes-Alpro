package app

import "fmt"

func ShowMenu() {
	fmt.Println("\n--- Menu ---")
	fmt.Println("1. Tampilkan semua akun")
	fmt.Println("2. Tampilkan akun")
	fmt.Println("3. Tambah akun dan video")
	fmt.Println("4. Hapus akun")
	fmt.Println("5. Keluar")
	fmt.Print("Masukkan pilihan Anda: ")
}
