package component

func CheckIndexAccount(A AccountList, x string) int {
	var i int
	i = 0
	for i, account := range A {
		if account.Username == x {
			return i
		}
		i++
	}
	return i
}
