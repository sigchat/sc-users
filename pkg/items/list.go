package items

import "slices"

// List вспомогательная не потокобезопасная структура для фильтрации слайсов
type List[T any] []T

// Filter фильтрует значения слайса, для которых функция fn возвращает true
func (i List[T]) Filter(fn func(item T, index int) bool) List[T] {
	res := List[T]{}

	for index, val := range i {
		if fn(val, index) {
			res = append(res, val)
		}
	}

	return res
}

// Each позволяет проходиться про всем элементам слайса. Работает до первого возвращаемого функцией значения true
func (i List[T]) Each(fn func(item T, index int) bool) List[T] {
	for index, val := range i {
		if fn(val, index) {
			break
		}
	}
	return i
}

func (i List[T]) Sorted(cmp func(a T, b T) int) List[T] {
	res := make(List[T], len(i))
	copy(res, i)
	slices.SortFunc(i, cmp)
	return res
}

func (i List[T]) Indexes() List[int] {
	res := make(List[int], len(i))
	i.Each(func(item T, index int) bool {
		res = append(res, index)
		return false
	})

	return res
}

func (i List[T]) ByIndexes(indexes ...int) List[T] {
	res := make(List[T], 0, len(indexes))
	for _, index := range indexes {
		res = append(res, i[index])
	}
	return res
}

func (i List[T]) Slice() []T {
	slice := make([]T, 0, len(i))
	i.Each(func(item T, index int) bool {
		slice = append(slice, item)
		return false
	})
	return slice
}

func (i List[T]) Trim(index int) List[T] {
	return i[:index]
}
