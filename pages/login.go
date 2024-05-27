package pages

import (
	"YouTubeAdSenseAPP/component"
	"fmt"
	"log"
)

func showLogin() {
	var name, password string

	component.ClearScreen()

	for {

		fmt.Print("username: ")
		fmt.Scanln(&name)

		data, err := component.ReadData()
		if err != nil {
			log.Fatalf("Failed to read data: %v", err)
		}

		if component.SearchUsername(data, name) {
			fmt.Print("password: ")
			fmt.Scanln(&password)

			if component.ValidatePassword(data, name, password) {
				showDisplaysAccounts()
				return
			} else {
				component.ClearScreen()
				fmt.Println("Password salah, coba lagi.")
			}
		} else {
			component.ClearScreen()
			fmt.Println("Username tidak ditemukan, coba lagi.")
		}
	}
}
