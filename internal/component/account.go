package component

const DataFile = "accounts.json"

const Nmax = 200

type Account struct {
	Username       string    `json:"username"`
	YouTubeChannel string    `json:"youtubeChannel"`
	Email          string    `json:"email"`
	Status         string    `json:"status"`
	Subscribers    int       `json:"subscribers"`
	Balance        int       `json:"balance"`
	Videos         VideoList `json:"videos"`
}

type Video struct {
	Title       string `json:"title"`
	Duration    int    `json:"duration"`
	ViewCount   int    `json:"viewCount"`
	PublishDate Date   `json:"publishDate"`
}

type Date struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type VideoList [Nmax]Video

type AccountList [10]Account
