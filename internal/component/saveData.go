package component

import (
	"encoding/json"
	"io/ioutil"
)

func SaveData(filename string, data AccountList) error {
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
