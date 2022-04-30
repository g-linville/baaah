package typed

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func Keys[K comparable, T any](data map[K]T) (result []K) {
	for k := range data {
		result = append(result, k)
	}
	return
}

func SortedKeys[K constraints.Ordered, T any](data map[K]T) (result []K) {
	for k := range data {
		result = append(result, k)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return
}

type Entry[K, V any] struct {
	Key   K
	Value V
}

func Sorted[K constraints.Ordered, V any](data map[K]V) []Entry[K, V] {
	var result []Entry[K, V]
	for _, key := range SortedKeys(data) {
		result = append(result, Entry[K, V]{
			Key:   key,
			Value: data[key],
		})
	}
	return result
}
