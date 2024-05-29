package component

func SearchUsername(A AccountList, x string) bool {
	var name bool

	name = false
	for _, account := range A {
		if account.Username == x {
			name = true
		}
	}
	return name
}
