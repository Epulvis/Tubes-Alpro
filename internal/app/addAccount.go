package app

import (
	"YouTubeAdSenseAPP/internal/component"
	"fmt"
)

func ShowAddAccount(data *component.AccountList, n *int) {
	var i int

	i = *n
	if i == -1 {
		fmt.Println("Data account sudah penuh.")
		return
	}

	component.ClearScreen()
	fmt.Println("Tambah akun baru:\n")
	component.AddAndEditAccount(&*data, i)

	*n = *n + 1
	component.ClearScreen()
	fmt.Print("Akun baru berhasil ditambahkan.\n")

}
