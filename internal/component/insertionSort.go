package component

func InsertionSort(A *VideoList, n int) {
	var i, j int

	for i = 0; i < n-1; i++ {
		maxIdx := i
		for j = i + 1; j < n; j++ {
			if A[j].PublishDate.Year < A[maxIdx].PublishDate.Year && A[j].PublishDate.Month < A[maxIdx].PublishDate.Month && A[j].PublishDate.Day < A[maxIdx].PublishDate.Day {
				maxIdx = j
			}
		}
		A[i], A[maxIdx] = A[maxIdx], A[i]
	}
}
