package component

import "fmt"

func AddAndEditAccount(data *AccountList, i int) {
	var j int
	var title string

	fmt.Print("Masukkan Nama Pengguna: ")
	fmt.Scanln(&data[i].Username)

	fmt.Print("Masukkan Nama Channel YouTube: ")
	fmt.Scanln(&data[i].YouTubeChannel)

	fmt.Print("Masukkan Nama Email: ")
	fmt.Scanln(&data[i].Email)

	fmt.Print("Masukkan Jumlah Subscribers: ")
	fmt.Scanln(&data[i].Subscribers)

	data[i].Status = CheckStatus(data[i].Subscribers)

	j = 0

	fmt.Print("Masukkan Judul Video (atau ketik 'selesai' untuk menyelesaikan): ")
	fmt.Scanln(&title)

	for title != "selesai" && j < Nmax {
		data[i].Videos[j].Title = title
		fmt.Print("Masukkan Durasi Video: ")
		fmt.Scanln(&data[i].Videos[j].Duration)

		fmt.Print("Masukkan Jumlah Penonton: ")
		fmt.Scanln(&data[i].Videos[j].ViewCount)

		fmt.Print("Masukkan Tanggal Upload (gunakan format DDMMYYYY): ")
		fmt.Scanf("%02d%02d%04d", &data[i].Videos[j].PublishDate.Day, &data[i].Videos[j].PublishDate.Month, &data[i].Videos[j].PublishDate.Year)
		fmt.Scanln()

		j++

		fmt.Print("Masukkan Judul Video (atau ketik 'selesai' untuk menyelesaikan): ")
		fmt.Scanln(&title)
	}

	InsertionSort(&data[i].Videos, j)
	data[i].Balance = BalanceCheck(data[i].Videos, data[i].Subscribers, i)

	SelectionSort(&*data, &i)
}
