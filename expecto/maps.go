package expecto

import "testing"

type AssertMap[K comparable, V any] struct {
	t *testing.T
	m map[K]V
}

func Map[K comparable, V any](t *testing.T, m map[K]V) *AssertMap[K, V] {
	return &AssertMap[K, V]{t: t, m: m}
}

func (a *AssertMap[K, V]) HasKey(msg string, key K) bool {
	_, ok := a.m[key]
	if !ok {
		a.t.Logf("  %s:", msg)
		a.t.Logf("     got: \033[0;31m%v\033[0m", a.m)
		a.t.Fatalf("Key not found")
	}
	return ok
}
