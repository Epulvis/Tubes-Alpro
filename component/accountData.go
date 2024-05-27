package component

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const DataFile = "accounts.json"

const Nmax = 200

type Account struct {
	Username       string    `json:"username"`
	YouTubeChannel string    `json:"youtube_channel"`
	Password       string    `json:"password"`
	Status         string    `json:"status"`
	Subscribers    int       `json:"subscribers"`
	Balance        int       `json:"balance"`
	Videos         VideoList `json:"videos"`
}

type Video struct {
	Title     string `json:"title"`
	Duration  int    `json:"duration"`
	ViewCount int    `json:"view_count"`
	publish   Date   `json:"date"`
}

type Date struct {
	date  int
	month int
	year  int
}

type VideoList [Nmax]Video
type AccountList [10]Account

func ReadAccountsFromFile(filename string) (AccountList, error) {
	var list AccountList

	// Baca file JSON
	file, err := os.Open(filename)
	if err != nil {
		return list, err
	}
	defer file.Close()

	// Baca isi file
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return list, err
	}

	// Unmarshal JSON ke dalam struct
	err = json.Unmarshal(byteValue, &list)
	if err != nil {
		return list, err
	}

	return list, nil
}

func SaveAccountsToFile(filename string, data AccountList) error {
	byteValue, err := json.MarshalIndent(data, "", "")
	if err != nil {

		return err
	}

	err = ioutil.WriteFile(filename, byteValue, 0644)
	if err != nil {
		return err
	}

	return nil
}

func CheckIndex(data AccountList) int {
	for i, account := range data {
		if account.Username == "" {
			return i
		}
	}
	return -1
}
