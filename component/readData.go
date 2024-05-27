package component

func ReadData() (AccountList, error) {
	return ReadAccountsFromFile(DataFile)
}
