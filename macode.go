package macode

import (
	"sort"
)

type Ordered interface {
	~string |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type Flatmap[K Ordered, V any] interface {
	Encode(map[K]V) []flatmap[K, V]
	Decode([]flatmap[K, V]) map[K]V
}

type flatmap[K Ordered, V any] struct {
	Key   K
	Value V
}

func New[K Ordered, V any](m map[K]V) Flatmap[K, V] {
	return flatmap[K, V]{}
}

func (p flatmap[K, V]) Encode(m map[K]V) []flatmap[K, V] {
	fm := make([]flatmap[K, V], 0, len(m))
	for k, v := range m {
		fm = append(fm, flatmap[K, V]{Key: k, Value: v})
	}
	sort.Slice(fm, func(i, j int) bool {
		return fm[i].Key < fm[j].Key
	})
	return fm
}

func (p flatmap[K, V]) Decode(ps []flatmap[K, V]) (ret map[K]V) {
	ret = make(map[K]V, len(ps))
	for _, pair := range ps {
		ret[pair.Key] = pair.Value
	}
	return
}
