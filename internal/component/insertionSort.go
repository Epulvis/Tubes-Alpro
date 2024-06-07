package component

func InsertionSort(A *VideoList, n int) {
	var i, j int
	var key Video

	for i = 1; i < n; i++ {
		key = (*A)[i]
		j = i - 1

		for j >= 0 && compareDates((*A)[j].PublishDate, key.PublishDate) > 0 {
			(*A)[j+1] = (*A)[j]
			j = j - 1
		}
		(*A)[j+1] = key
	}
}
