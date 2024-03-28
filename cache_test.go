package cache

import "testing"

func TestCacheParameters(t *testing.T) {
	c := New[string]()
	callCount := 0

	var f = func(a, b string) Cacheable[string] {
		return func() string {
			callCount++

			return a + b
		}
	}

	for range 10 {
		c.WithCache(f("1", "2"), "1", "2")
	}

	if callCount != 1 {
		t.Fatalf("call cound should have been 1 but was %d", callCount)
	}
}

func TestCacheChangingParameters(t *testing.T) {
	c := New[int]()
	callCount := 0

	var f = func(a, b int) Cacheable[int] {
		return func() int {
			callCount++

			return a + b
		}
	}

	var result int
	for i := range 10 {
		result = c.WithCache(f(i, i+1), i, i+1)
	}

	if callCount != 10 {
		t.Fatalf("call cound should have been 1 but was %d", callCount)
	}

	if result != 19 {
		t.Fatalf("result should have been 19 but was %d", result)
	}
}
