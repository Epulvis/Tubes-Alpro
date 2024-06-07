package component

import "fmt"

func SequentialSearch(T AccountList, n int, X string) {
	var newT AccountList
	var j, i int

	for j < n {
		if T[j].Status == X {
			newT[i] = T[j]
			i++
		}
		j++
	}

	if i == 0 {
		fmt.Print("status tidak ditemukan\n")
		return
	}

	DisplaysAccounts(newT, i)
}
