package pages

import (
	"YouTubeAdSenseAPP/component"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func showDisplaysAccounts() {
	reader := bufio.NewReader(os.Stdin)
	component.ClearScreen()
	fmt.Println("Menampilkan akun : ")
	fmt.Println(strings.Repeat("=", 89))
	fmt.Printf("| %-30s | %-20s | %-15s |  %-10s |\n", "Nama Pengguna", "ChannelYouTube", "Status", "Saldo")
	fmt.Println(strings.Repeat("=", 89))

	list, err := component.ReadAccountsFromFile("accounts.json")
	if err != nil {
		log.Fatalf("Failed to read accounts from file: %s", err)
	}

	for i, account := range list {
		if account.Username != "" {
			fmt.Printf("| %-30s | %-20s | %-15s |  %-10d |\n", account.Username, account.YouTubeChannel, account.Status, account.Balance)
			fmt.Println(strings.Repeat("-", 89))
			i++
		}
	}

	fmt.Println("\nTekan 'Enter' untuk kembali ke menu...")
	reader.ReadString('\n')
	component.ClearScreen()
}
