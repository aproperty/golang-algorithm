package main

type Item struct {
	iValie int
	pNext  *Item
}

func getItem(pHead *Item, m int) *Item {
	if pHead == nil {
		return nil
	}
	if m == 0 {
		m = 1
	}

	ahead := pHead

	for i := 0; i < m-1; i++ {
		if ahead.pNext != nil {
			ahead = ahead.pNext
		} else {
			return nil
		}
	}

	behind := pHead

	for ahead.pNext != nil {
		ahead = ahead.pNext
		behind = behind.pNext
	}

	return behind
}
