package component

func CheckIndex(data AccountList) int {
	for i, account := range data {
		if account.Username == "" {
			return i
		}
	}
	return -1
}
