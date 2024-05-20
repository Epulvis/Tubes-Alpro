package component

import (
	"encoding/json"
	"io/ioutil"
)

func SaveAccount(filename string, accounts AccountList) error {
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
