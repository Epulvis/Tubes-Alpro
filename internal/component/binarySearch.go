package component

func BinarySearch(T AccountList, n int, x string) int {
	var kr, kn, med, found int
	kr = 0
	kn = n - 1
	found = -1
	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if T[med].Username < x {
			kr = med + 1
		} else if T[med].Username > x {
			kn = med - 1
		} else {
			found = med
		}
	}
	return found
}
