package designhashset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyHashSet(t *testing.T) {
	hs := Constructor()

	assert.False(t, hs.Contains(1), "empty set should not contain 1")

	hs.Add(1)
	assert.True(t, hs.Contains(1), "set should contain 1 after Add")

	hs.Add(2)
	assert.True(t, hs.Contains(1), "set should still contain 1")
	assert.True(t, hs.Contains(2), "set should contain 2 after Add")

	hs.Remove(1)
	assert.False(t, hs.Contains(1), "set should not contain 1 after Remove")
	assert.True(t, hs.Contains(2), "set should still contain 2")

	hs.Remove(2)
	assert.False(t, hs.Contains(1), "set should not contain 1")
	assert.False(t, hs.Contains(2), "set should not contain 2 after Remove")
}

func TestMyHashSetDuplicates(t *testing.T) {
	hs := Constructor()

	hs.Add(5)
	hs.Add(5)
	hs.Add(5)
	assert.True(t, hs.Contains(5), "set should contain 5 after duplicate Adds")

	hs.Remove(5)
	assert.False(t, hs.Contains(5), "set should not contain 5 after Remove")
}

func TestMyHashSetNonExistent(t *testing.T) {
	hs := Constructor()

	assert.False(t, hs.Contains(0), "empty set should not contain 0")
	assert.False(t, hs.Contains(-1), "empty set should not contain -1")
	assert.False(t, hs.Contains(100), "empty set should not contain 100")
}

func TestMyHashSetNegativeKeys(t *testing.T) {
	hs := Constructor()

	hs.Add(-1)
	hs.Add(-5)
	assert.True(t, hs.Contains(-1), "set should contain -1")
	assert.True(t, hs.Contains(-5), "set should contain -5")
	assert.False(t, hs.Contains(-2), "set should not contain -2")

	hs.Remove(-1)
	assert.False(t, hs.Contains(-1), "set should not contain -1 after Remove")
	assert.True(t, hs.Contains(-5), "set should still contain -5")
}

func TestMyHashSetZero(t *testing.T) {
	hs := Constructor()

	hs.Add(0)
	assert.True(t, hs.Contains(0), "set should contain 0")

	hs.Remove(0)
	assert.False(t, hs.Contains(0), "set should not contain 0 after Remove")
}

func TestMyHashSetLargeKeys(t *testing.T) {
	hs := Constructor()

	keys := []int{1000, 100000, 1000000}
	for _, k := range keys {
		hs.Add(k)
	}
	for _, k := range keys {
		assert.True(t, hs.Contains(k), "set should contain %d", k)
	}

	hs.Remove(100000)
	assert.False(t, hs.Contains(100000), "set should not contain 100000 after Remove")
	assert.True(t, hs.Contains(1000), "set should still contain 1000")
	assert.True(t, hs.Contains(1000000), "set should still contain 1000000")
}
