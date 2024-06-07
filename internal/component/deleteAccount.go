package component

func DeleteAccount(T *AccountList, n *int, idx int) {
	var i int
	for i = idx; i < *n; i++ {
		T[i] = T[i+1]
	}
	*n = *n - 1

}
