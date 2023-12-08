package items

// Map вспомогательная не потокобезопасная структура для фильтрации мап
type Map[K comparable, V any] map[K]V

// Filter фильтрует значения слайса, для которых функция fn возвращает true
func (i Map[K, V]) Filter(fn func(key K, value V, index int) bool) Map[K, V] {
	res := Map[K, V]{}

	k := 0
	for key, val := range i {
		if fn(key, val, k) {
			res[key] = val
		}
		k++
	}

	return res
}

// Each позволяет проходиться про всем элементам слайса. Работает до первого возвращаемого функцией значения true
func (i Map[K, V]) Each(fn func(key K, value V, index int) bool) Map[K, V] {
	k := 0
	for key, val := range i {
		if fn(key, val, k) {
			break
		}
		k++
	}
	return i
}
