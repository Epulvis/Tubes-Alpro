package component

import (
	"fmt"
	"strings"
)

func displayVideos(videos VideoList) {
	var header, row string
	header = "| %-30s | %-10s | %-20s | %-13s |\n"
	row = "| %-30s | %-10d | %-20d | %02d-%02d-%04d     |\n"

	fmt.Println("\nDaftar Video:")
	fmt.Println(strings.Repeat("-", 87))
	fmt.Printf(header, "Judul", "Durasi", "Jumlah Penonton", "Tanggal Upload")
	fmt.Println(strings.Repeat("-", 87))

	for _, video := range videos {
		if video.Title != "" {
			fmt.Printf(row, video.Title, video.Duration, video.ViewCount, video.PublishDate.Day, video.PublishDate.Month, video.PublishDate.Year)
		}
	}

	fmt.Println(strings.Repeat("-", 87))
}
