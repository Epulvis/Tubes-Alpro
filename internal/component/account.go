package component

const Nmax = 20

type Account struct {
	Username       string
	YouTubeChannel string
	Email          string
	Status         string
	Subscribers    int
	Balance        int
	Videos         VideoList
}

type Video struct {
	Title       string
	Duration    int
	ViewCount   int
	PublishDate Date
}

type Date struct {
	Day   int
	Month int
	Year  int
}

type VideoList [Nmax]Video

type AccountList [Nmax]Account
