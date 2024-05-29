package component

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadData(filename string) (AccountList, error) {
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
