package component

import (
	"log"
)

func CheckIndex() int {
	var count int
	list, err := ReadAccountsFromFile(DataFile)
	if err != nil {
		log.Fatalf("Failed to read accounts from file: %s", err)
	}

	count = 0
	for _, account := range list {
		if account.Username != "" {
			count++
		}
	}

	return count
}
