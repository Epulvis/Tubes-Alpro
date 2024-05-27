package component

func ValidatePassword(A AccountList, username, password string) bool {

	for _, account := range A {
		if account.Username == username && account.Password == password {
			return true
		}
	}
	return false
}
