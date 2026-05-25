package lrucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCacheBasic(t *testing.T) {
	cache := Constructor(2)

	assert.Equal(t, -1, cache.Get(1), "empty cache should return -1")

	cache.Put(1, 10)
	assert.Equal(t, 10, cache.Get(1), "should get 10 after Put(1, 10)")

	cache.Put(2, 20)
	assert.Equal(t, 10, cache.Get(1), "should still get 10")
	assert.Equal(t, 20, cache.Get(2), "should get 20 after Put(2, 20)")
}

func TestLRUCacheUpdate(t *testing.T) {
	cache := Constructor(2)

	cache.Put(1, 10)
	assert.Equal(t, 10, cache.Get(1))

	cache.Put(1, 100)
	assert.Equal(t, 100, cache.Get(1), "value should be updated to 100")
}

func TestLRUCacheEviction(t *testing.T) {
	cache := Constructor(2)

	cache.Put(1, 10)
	cache.Put(2, 20)
	cache.Put(3, 30)

	assert.Equal(t, -1, cache.Get(1), "key 1 should be evicted")
	assert.Equal(t, 20, cache.Get(2), "key 2 should still exist")
	assert.Equal(t, 30, cache.Get(3), "key 3 should exist")
}

func TestLRUCacheRecentlyUsed(t *testing.T) {
	cache := Constructor(2)

	cache.Put(1, 10)
	cache.Put(2, 20)
	cache.Get(1)
	cache.Put(3, 30)

	assert.Equal(t, 10, cache.Get(1), "key 1 should still exist (was recently used)")
	assert.Equal(t, -1, cache.Get(2), "key 2 should be evicted (least recently used)")
	assert.Equal(t, 30, cache.Get(3), "key 3 should exist")
}

func TestLRUCacheCapacityOne(t *testing.T) {
	cache := Constructor(1)

	cache.Put(1, 10)
	assert.Equal(t, 10, cache.Get(1))

	cache.Put(2, 20)
	assert.Equal(t, -1, cache.Get(1), "key 1 should be evicted")
	assert.Equal(t, 20, cache.Get(2), "key 2 should exist")

	cache.Put(2, 200)
	assert.Equal(t, 200, cache.Get(2), "key 2 should have updated value")
}

func TestLRUCacheSequence(t *testing.T) {
	cache := Constructor(3)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)

	assert.Equal(t, -1, cache.Get(1), "key 1 should be evicted")
	assert.Equal(t, 2, cache.Get(2))
	assert.Equal(t, 3, cache.Get(3))
	assert.Equal(t, 4, cache.Get(4))

	cache.Put(5, 5)

	assert.Equal(t, -1, cache.Get(2), "key 2 should be evicted (least recently used)")
	assert.Equal(t, 3, cache.Get(3), "key 3 should still exist")
	assert.Equal(t, 4, cache.Get(4), "key 4 should still exist")
	assert.Equal(t, 5, cache.Get(5))
}
