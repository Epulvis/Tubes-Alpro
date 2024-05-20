package component

import "log"

func ReadData(A AccountList) AccountList {
	list, err := ReadAccountsFromFile(DataFile)
	if err != nil {
		log.Fatalf("Failed to read accounts from file: %s", err)
	}

	for i, account := range list {
		if account.Username != "" {
			A[i] = account
		}
	}

	return A
}
