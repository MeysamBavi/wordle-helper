package sort

import (
	"container/list"
)

func Sort(l *list.List, less func(a, b *list.Element) bool) *list.List {
	return mergeSort(l.Front(), l.Len(), less)
}

func mergeSort(e *list.Element, length int, less func(a, b *list.Element) bool) *list.List {
	if length <= 1 {
		newList := list.New()
		newList.PushBack(e.Value)
		return newList
	}

	mid := length / 2

	midElement := e
	for i := 0; i < mid; i++ {
		midElement = midElement.Next()
	}

	a := mergeSort(e, mid, less)
	b := mergeSort(midElement, length-mid, less)
	return merge(a, b, less)
}

func merge(aList, bList *list.List, less func(a, b *list.Element) bool) *list.List {
	newList := list.New()
	a, b := aList.Front(), bList.Front()

	for a != nil && b != nil {
		var min *list.Element

		if less(a, b) {
			min = a
			a = a.Next()
		} else {
			min = b
			b = b.Next()
		}

		newList.PushBack(min.Value)
	}

	for a != nil {
		newList.PushBack(a.Value)
		a = a.Next()
	}

	for b != nil {
		newList.PushBack(b.Value)
		b = b.Next()
	}

	return newList
}
