package app

import (
	"YouTubeAdSenseAPP/internal/component"
	"fmt"
	"log"
)

func showAddAccount() {
	var data component.AccountList
	var i, j int

	data, err := component.ReadData(component.DataFile)
	if err != nil {
		log.Fatalf("Gagal membaca data: %v", err)
	}

	i = component.CheckIndex(data)
	if i == -1 {
		fmt.Println("Data account sudah penuh.")
		return
	}

	component.ClearScreen()
	fmt.Println("Tambah akun baru:")
	fmt.Print("Masukan Nama Pengguna: ")
	fmt.Scanln(&data[i].Username)

	fmt.Print("Masukan Nama Channel YouTube: ")
	fmt.Scanln(&data[i].YouTubeChannel)

	fmt.Print("Masukan Nama Email: ")
	fmt.Scanln(&data[i].Email)

	fmt.Print("Masukan Jumlah Subscribers: ")
	fmt.Scanln(&data[i].Subscribers)

	data[i].Status = component.CheckStatus(data[i].Subscribers)

	j = 0

	for {
		fmt.Print("Masukan Judul Video (atau ketik 'selesai' untuk menyelesaikan): ")
		fmt.Scanln(&data[i].Videos[j].Title)

		if data[i].Videos[j].Title == "selesai" {
			break
		}

		fmt.Print("Masukan Durasi Video: ")
		fmt.Scanln(&data[i].Videos[j].Duration)

		fmt.Print("Masukan Jumlah Penonton: ")
		fmt.Scanln(&data[i].Videos[j].ViewCount)

		fmt.Print("Masukan Tanggal Upload (gunakan format DD): ")
		fmt.Scanln(&data[i].Videos[j].PublishDate.Day)

		fmt.Print("Masukan Bulan Upload (gunakan format MM): ")
		fmt.Scanln(&data[i].Videos[j].PublishDate.Month)

		fmt.Print("Masukan Tahun Upload (gunakan format YY): ")
		fmt.Scanln(&data[i].Videos[j].PublishDate.Year)
		j++
	}

	data[i].Videos = component.InsertionSort(&data[i].Videos, j)

	err = component.SaveData(component.DataFile, data)
	if err != nil {
		log.Fatalf("gagal dalam menyimpan data: %v", err)
	}

	component.ClearScreen()
	fmt.Println("Akun baru berhasil ditambahkan.")

}
