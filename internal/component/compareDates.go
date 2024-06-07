package component

func compareDates(d1, d2 struct{ Day, Month, Year int }) int {
	if d1.Year != d2.Year {
		return d1.Year - d2.Year
	}
	if d1.Month != d2.Month {
		return d1.Month - d2.Month
	}
	return d1.Day - d2.Day
}
