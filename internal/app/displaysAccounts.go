package app

import (
	"YouTubeAdSenseAPP/internal/component"
	"bufio"
	"fmt"
	"os"
)

func ShowDisplaysAccounts(data *component.AccountList, n *int) {
	var name string
	var input byte
	var index int

	reader := bufio.NewReader(os.Stdin)
	component.ClearScreen()
	fmt.Print("\nMasukkan username: ")
	fmt.Scanln(&name)
	index = component.BinarySearch(*data, *n, name)
	component.ClearScreen()

	if index != -1 {
		component.DisplaysAccount(*data, component.BinarySearch(*data, *n, name))
		fmt.Println("Apakah Anda ingin mengedit data ini?(Y/n)")
		input, _ = reader.ReadByte()
		if input == 'Y' || input == 'y' {
			component.ClearScreen()
			component.AddAndEditAccount(&*data, index)
			component.ClearScreen()
			fmt.Print("Data berhasil diedit\n")

		} else if input == 'n' || input == 'N' {
			component.ClearScreen()
		} else {
			component.ClearScreen()
			fmt.Print("Pilihan tidak valid, silakan coba lagi.\n")
		}

	} else {
		fmt.Print("Username tidak ditemukan, coba lagi.\n")
	}
}
