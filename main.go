package main

import (
	"YouTubeAdSenseAPP/internal/app"
	"YouTubeAdSenseAPP/internal/component"
	"fmt"
)

func main() {
	var input string
	var nData int
	var data component.AccountList

	data = component.SnowData()
	nData = component.CheckIndex(data)
	component.SelectionSort(&data, &nData)

	fmt.Println("Selamat datang di aplikasi pengelola akun YouTube!")
	fmt.Println("Aplikasi ini akan membantu Anda menampilkan informasi akun dan video.")
	fmt.Println("Silakan pilih opsi dari menu di bawah ini.")

	for {
		app.ShowMenu()

		fmt.Scanln(&input)

		if input == "1" {
			app.ShowDisplaysAllAccounts(&data, &nData)
		} else if input == "2" {
			app.ShowDisplaysAccounts(&data, &nData)
		} else if input == "3" {
			app.ShowAddAccount(&data, &nData)
		} else if input == "4" {
			app.ShowDeleteAccount(&data, &nData)
		} else if input == "5" {
			return
		} else {
			component.ClearScreen()
			fmt.Print("Pilihan tidak valid, silakan coba lagi.\n")
		}

	}
}
