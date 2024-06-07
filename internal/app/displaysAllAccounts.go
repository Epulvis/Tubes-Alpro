package app

import (
	"YouTubeAdSenseAPP/internal/component"
	"bufio"
	"fmt"
	"os"
)

func ShowDisplaysAllAccounts(data *component.AccountList, index *int) {
	var input, x string

	scanner := bufio.NewScanner(os.Stdin)
	component.ClearScreen()
	component.DisplaysAccounts(*data, *index)

	fmt.Println("\nPilih metode pencarian data akun YouTube:")
	fmt.Println("1. Berdasarkan Channel")
	fmt.Println("2. Berdasarkan Status")
	fmt.Print("Masukkan pilihan Anda: ")

	scanner.Scan()
	input = scanner.Text()
	component.ClearScreen()

	if input == "1" {
		ShowDisplaysAccounts(&*data, &*index)

	} else if input == "2" {
		component.ClearScreen()
		fmt.Print("Masukkan status: ")
		fmt.Scanln(&x)
		component.ClearScreen()
		component.SequentialSearch(*data, *index, x)

		fmt.Println("Tekan Enter untuk kembali ke menu utama...")
		scanner.Scan()
		component.ClearScreen()
	} else {
		fmt.Print("Pilihan tidak valid, silakan coba lagi.\n")
	}
}
