package component

func CheckUsernameIndex(A AccountList, x string) int {
	var i int
	i = 0
	for i, account := range A {
		if account.Username == x {
			break
		}
		i++
	}
	return i
}
