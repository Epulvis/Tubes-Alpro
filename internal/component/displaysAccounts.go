package component

import (
	"fmt"
	"strings"
)

func DisplaysAccounts(data AccountList, index int) {
	fmt.Println("Menampilkan akun :")
	fmt.Println(strings.Repeat("=", 106))
	fmt.Printf("| %-20s | %-20s | %-30s | %-10s | %-10s |\n", "Nama Pengguna", "ChannelYouTube", "Email", "Status", "Saldo")

	for i := 0; i < index; i++ {
		PrintComplexSeparator()
		fmt.Printf("| %-20s | %-20s | %-30s | %-10s | %-10d |\n", data[i].Username, data[i].YouTubeChannel, data[i].Email, data[i].Status, data[i].Balance)
	}
	fmt.Println(strings.Repeat("=", 106))
}
