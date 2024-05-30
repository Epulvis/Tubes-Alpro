package component

func CheckStatus(x int) string {
	var status string
	status = "Silver"

	if x > 1000 {
		if x > 100000 {
			if x > 10000000 {
				status = "Diamond"
			}
			status = "Platinum"
		}
		status = "Gold"
	}
	return status
}
