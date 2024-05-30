package app

import (
	"YouTubeAdSenseAPP/internal/component"
	"bufio"
	"fmt"
	"log"
	"os"
)

func showDeleteAccount() {
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
			component.DeleteAccount(&data, component.CheckIndex(data), component.CheckUsernameIndex(data, name))

			err = component.SaveData(component.DataFile, data)
			if err != nil {
				log.Fatalf("gagal dalam menyimpan data: %v", err)
			}

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
