package pages

import (
	"YouTubeAdSenseAPP/component"
	"bufio"
	"fmt"
	"log"
	"os"
)

func showDisplaysAccounts() {
	var name string

	reader := bufio.NewReader(os.Stdin)

	data, err := component.ReadData()
	if err != nil {
		log.Fatalf("Failed to read data: %v", err)
	}

	component.ClearScreen()
	for {

		fmt.Print("username: ")
		fmt.Scanln(&name)

		if component.SearchUsername(data, name) {
			displaysAccounts(data, component.CheckIndexAccount(data, name))

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
	fmt.Println(data[a].Username)
	fmt.Println(data[a].YouTubeChannel)
	fmt.Println(data[a].Subscribers)
	fmt.Println(data[a].Status)
	fmt.Println(data[a].Balance)
}
