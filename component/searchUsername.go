package component

func SearchUsername(A AccountList, x string) bool {
	for _, account := range A {
		if account.Username == x {
			return true
		}
	}
	return false
}
