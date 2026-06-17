package minimumindexsumoftwolists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRestaurant(t *testing.T) {
	assert.Equal(t, []string{}, findRestaurant([]string{}, []string{"sad", "happy", "good"}))
	assert.Equal(t, []string{"sad", "happy"}, findRestaurant([]string{"happy", "sad", "good"}, []string{"sad", "happy", "good"}))
	assert.Equal(t, []string{"Shogun"}, findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}))
	assert.Equal(t, []string{"Shogun"}, findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Shogun", "Burger King"}))
}
