package component

func BalanceCheck(A VideoList, T, n int) int {
	var balance, i int
	for i = 0; i <= n; i++ {
		balance += A[i].ViewCount * T
	}
	return balance
}
