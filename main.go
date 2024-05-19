package main

const Nmax = 1000

type Account struct {
	Username, YouTubeChannel, Password, Subscribers, Status string
	Balance                                                 int
	Videos                                                  VideoList
}

type Video struct {
	Title     string
	Duration  int
	ViewCount int
}

type VideoList [Nmax]Video
type AccountList [Nmax]Account

func main() {
	// Your code here
}
