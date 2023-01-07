package heapsort

func createHeap(arr []int, end int) {

	for start := end / 2; start >= 0; start-- {
		maxHeap(start, end, arr)
	}
}

func HeapSort(c []int) []int {

	m := len(c) - 1

	createHeap(c, m)

	for i := m; i > 0; i-- {

		if c[0] > c[i] {
			c[0], c[i] = c[i], c[0]

			createHeap(c, i-1)
		}
	}

	return c
}
