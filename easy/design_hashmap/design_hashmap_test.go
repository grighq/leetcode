package designhashmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyHashMap(t *testing.T) {
	hm := Constructor()

	assert.Equal(t, -1, hm.Get(1), "empty map should return -1")

	hm.Put(1, 10)
	assert.Equal(t, 10, hm.Get(1), "should get 10 after Put(1, 10)")

	hm.Put(2, 20)
	assert.Equal(t, 10, hm.Get(1), "should still get 10")
	assert.Equal(t, 20, hm.Get(2), "should get 20 after Put(2, 20)")

	hm.Remove(1)
	assert.Equal(t, -1, hm.Get(1), "should return -1 after Remove(1)")
	assert.Equal(t, 20, hm.Get(2), "should still get 20")

	hm.Remove(2)
	assert.Equal(t, -1, hm.Get(1), "should return -1")
	assert.Equal(t, -1, hm.Get(2), "should return -1 after Remove(2)")
}

func TestMyHashMapUpdate(t *testing.T) {
	hm := Constructor()

	hm.Put(1, 10)
	assert.Equal(t, 10, hm.Get(1))
	assert.Equal(t, -1, hm.Get(2))

	hm.Put(1, 100)
	assert.Equal(t, 100, hm.Get(1), "value should be updated to 100")
}

func TestMyHashMapNonExistent(t *testing.T) {
	hm := Constructor()

	assert.Equal(t, -1, hm.Get(0), "empty map should return -1 for 0")
	assert.Equal(t, -1, hm.Get(-1), "empty map should return -1 for -1")
	assert.Equal(t, -1, hm.Get(100), "empty map should return -1 for 100")
}

func TestMyHashMapNegativeKeys(t *testing.T) {
	hm := Constructor()

	hm.Put(-1, 10)
	hm.Put(-5, 50)
	assert.Equal(t, 10, hm.Get(-1), "should get 10 for key -1")
	assert.Equal(t, 50, hm.Get(-5), "should get 50 for key -5")
	assert.Equal(t, -1, hm.Get(-2), "should return -1 for missing key -2")

	hm.Remove(-1)
	assert.Equal(t, -1, hm.Get(-1), "should return -1 after Remove(-1)")
	assert.Equal(t, 50, hm.Get(-5), "should still get 50")
}

func TestMyHashMapZero(t *testing.T) {
	hm := Constructor()

	hm.Put(0, 99)
	assert.Equal(t, 99, hm.Get(0), "should get 99 for key 0")

	hm.Remove(0)
	assert.Equal(t, -1, hm.Get(0), "should return -1 after Remove(0)")
}

func TestMyHashMapLargeKeys(t *testing.T) {
	hm := Constructor()

	pairs := map[int]int{1000: 1, 100000: 2, 1000000: 3}
	for k, v := range pairs {
		hm.Put(k, v)
	}
	for k, v := range pairs {
		got := hm.Get(k)
		assert.Equal(t, v, got, "key %d should map to %d", k, v)
	}

	hm.Remove(100000)
	assert.Equal(t, -1, hm.Get(100000), "should return -1 after removing key 100000")
	assert.Equal(t, 1, hm.Get(1000), "key 1000 should still map to 1")
	assert.Equal(t, 3, hm.Get(1000000), "key 1000000 should still map to 3")
}
