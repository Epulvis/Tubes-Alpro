package component

func CheckStatus(x int) string {
	if x > 1000 {
		if x > 100000 {
			if x > 10000000 {
				return "Diamond"
			}
			return "Platinum"
		}
		return "Gold"
	}
	return "Silver"
}
