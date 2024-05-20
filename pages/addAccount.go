package pages

import (
	"YouTubeAdSenseAPP/component"
	"fmt"
)

func showAddAccount() {
	var data component.AccountList
	var i int

	i = component.CheckIndex()

	component.ClearScreen()
	fmt.Println("Tambah akun baru:")
	fmt.Print("Nama Pengguna: ")
	_, err := fmt.Scanln(&data[i].Username)
	if err != nil {
		return
	}

	fmt.Print("Nama Channel YouTube: ")
	_, err = fmt.Scanln(&data[i].YouTubeChannel)
	if err != nil {
		return
	}

	fmt.Print("Password: ")
	_, err = fmt.Scanln(&data[i].Password)
	if err != nil {
		return
	}

	fmt.Print("Jumlah Subscribers: ")
	_, err = fmt.Scanln(&data[i].Subscribers)
	if err != nil {
		return
	}

	for {
		j := 0
		fmt.Print("Judul Video (atau ketik 'selesai' untuk menyelesaikan): ")
		_, err := fmt.Scanln(&data[i].Videos[j].Title)
		if err != nil {
			return
		}
		if data[i].Videos[j].Title == "selesai" {
			break
		}

		fmt.Print("Durasi: ")
		_, err = fmt.Scanln(&data[i].Videos[j].Duration)
		if err != nil {
			return
		}

		fmt.Print("Jumlah Penonton: ")
		_, err = fmt.Scanln(&data[i].Videos[j].ViewCount)
		if err != nil {
			return
		}

		fmt.Print("Tanggal Upload (gunakan format DDMMYYYY): ")
		_, err = fmt.Scanln(&data[i].Videos[j].Date)
		if err != nil {
			return
		}

		j++
	}

	component.ClearScreen()

}
