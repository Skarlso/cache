package cache

import (
	"fmt"
	"hash/fnv"
	"reflect"
	"sync"
)

// Cacheable defines the function to call. They would have to do something like
// func Bla(a, b string) Cacheable[string] { return func() string { return a + b } }
// but I need the arguments... maybe something like
// func Bla(a, b string) Cacheable[string] { return func(args ...any) string { return a + b }(a, b) }
type Cacheable[T any] func() T

// What if we would use the cache together with Tuple to achieve any kind of return result?

// Cache contains the values within a tuple.
type Cache[T any] struct {
	values map[uint64]T

	mu sync.RWMutex
}

func New[T any]() *Cache[T] {
	return &Cache[T]{
		values: make(map[uint64]T),
	}
}

// WithCache takes a function to call and cache the results for and args contains any arguments
// that the cache should be based on.
func (c *Cache[T]) WithCache(f Cacheable[T], args ...any) (value T) {
	key := c.generateHash(args)

	c.mu.Lock()
	defer c.mu.Unlock()

	if v, ok := c.values[key]; ok {
		return v
	}

	value = f()
	c.values[key] = value

	return value
}

func (c *Cache[T]) generateHash(args ...any) uint64 {
	h := fnv.New64a()
	v := reflect.ValueOf(args)

	// Iterate over the elements of the slice
	for i := 0; i < v.Len(); i++ {
		element := v.Index(i).Interface()
		// Hash each element
		if _, err := h.Write([]byte(fmt.Sprintf("%#v", element))); err != nil {
			panic(err)
		}
	}

	return h.Sum64()
}
