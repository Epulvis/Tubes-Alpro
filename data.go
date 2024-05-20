p

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Data struct untuk menyimpan informasi akun
type Account struct {
	NamaPengguna   string  `json:"nama_pengguna"`
	ChannelYouTube string  `json:"channel_youtube"`
	Saldo          string  `json:"saldo"`
	StatusYouTube  string  `json:"status_youtube"`
	Video          []Video `json:"video"`
}

// Data struct untuk menyimpan informasi video
type Video struct {
	Judul  string `json:"judul"`
	Durasi string `json:"durasi"`
	Views  string `json:"views"`
}

// Fungsi untuk memuat data akun dari file
func LoadAccounts(filename string) ([]Account, error) {
	var accounts []Account
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return accounts, nil // File tidak ada, kembali dengan array kosong
		}
		return nil, err
	}
	err = json.Unmarshal(file, &accounts)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// Fungsi untuk menyimpan data akun ke file
func SaveAccounts(filename string, accounts []Account) error {
	data, err := json.MarshalIndent(accounts, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Fungsi untuk menghapus akun berdasarkan nama pengguna
func DeleteAccount(accounts []Account, namaPengguna string) ([]Account, bool) {
	for i, account := range accounts {
		if account.NamaPengguna == namaPengguna {
			accounts = append(accounts[:i], accounts[i+1:]...)
			return accounts, true
		}
	}
	return accounts, false
}
