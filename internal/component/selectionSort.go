package component

func SelectionSort(T *AccountList, n *int) {
	var i, idxMin, j int
	for i = 0; i < *n-1; i++ {
		idxMin = i
		for j = i + 1; j < *n; j++ {
			if (*T)[j].Username < (*T)[idxMin].Username {
				idxMin = j
			}
		}
		(*T)[i], (*T)[idxMin] = (*T)[idxMin], (*T)[i]
	}
}
