package component

import "fmt"

func DisplaysAccount(data AccountList, a int) {
	var format string

	format = "%-20s: %s\n"
	fmt.Printf(format, "Username", data[a].Username)
	fmt.Printf(format, "YouTube Channel", data[a].YouTubeChannel)
	fmt.Printf(format, "Email", data[a].Email)
	fmt.Printf(format, "Status", data[a].Status)
	fmt.Printf(format, "Subscribers", fmt.Sprintf("%d", data[a].Subscribers))
	fmt.Printf(format, "Balance", fmt.Sprintf("%d", data[a].Balance))
	displayVideos(data[a].Videos)
}
