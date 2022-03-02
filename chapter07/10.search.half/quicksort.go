package main

func quickSort(arr *[]int64, start, end int) {
	pivotIdx := (start + end) / 2
	pivotVal := (*arr)[pivotIdx]
	l, r := start, end
	for l <= r {
		for (*arr)[l] < pivotVal {
			l++
		}
		for (*arr)[r] > pivotVal {
			r--
		}
		if l >= r {
			break
		}
		(*arr)[l], (*arr)[r] = (*arr)[r], (*arr)[l]
		l++
		r--
	}
	if l == r {
		l++
		r--
	}
	if r > start {
		quickSort(arr, start, r)
	}
	if l < end {
		quickSort(arr, l, end)
	}

}
