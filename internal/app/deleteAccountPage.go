package app

import (
	"YouTubeAdSenseAPP/internal/component"
	"fmt"
)

func ShowDeleteAccount(data *component.AccountList, n *int) {
	var name string
	var index int

	component.ClearScreen()

	fmt.Print("\nMasukkan username: ")
	fmt.Scanln(&name)
	index = component.BinarySearch(*data, *n, name)

	if index != -1 {
		component.ClearScreen()
		component.DeleteAccount(*&data, *&n, index)
		fmt.Print("Data berhasil di hapus\n")

	} else {
		component.ClearScreen()
		fmt.Print("Username tidak ditemukan, coba lagi.\n")
	}
}
