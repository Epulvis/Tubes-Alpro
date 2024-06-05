package app

import (
	"YouTubeAdSenseAPP/internal/component"
	"bufio"
	"fmt"
	"log"
	"os"
)

func showDisplaysAccounts() {
	var name string
	reader := bufio.NewReader(os.Stdin)

	data, err := component.ReadData(component.DataFile)
	if err != nil {
		log.Fatalf("Gagal membaca data: %v", err)
	}

	component.ClearScreen()
	for {

		fmt.Print("username: ")
		fmt.Scanln(&name)

		if component.SearchUsername(data, name) {
			displaysAccounts(data, component.CheckUsernameIndex(data, name))

			fmt.Println("\nTekan 'Enter' untuk kembali ke menu...")
			reader.ReadString('\n')
			component.ClearScreen()
			return

		} else {
			component.ClearScreen()
			fmt.Println("Username tidak ditemukan, coba lagi.")
		}
	}
}

func displaysAccounts(data component.AccountList, a int) {
	fmt.Printf("Username: %s\n", data[a].Username)
	fmt.Printf("YouTube Channel: %s\n", data[a].YouTubeChannel)
	fmt.Printf("Email: %s\n", data[a].Email)
	fmt.Printf("Status: %s\n", data[a].Status)
	fmt.Printf("Subscribers: %d\n", data[a].Subscribers)
	fmt.Printf("Balance: %d\n", data[a].Balance)
	for _, video := range data[a].Videos {
		if video.Title != "" {
			fmt.Printf("Video Title: %s\n", video.Title)
			fmt.Printf("Duration: %d\n", video.Duration)
			fmt.Printf("ViewCount: %d\n", video.ViewCount)
			fmt.Printf("Publish Date: %02d-%02d-%04d\n", video.PublishDate.Day, video.PublishDate.Month, video.PublishDate.Year)
		}
	}
}
