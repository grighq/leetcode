package designhashset

type MyHashSet struct {
	buckets [][]int
	size    int
}

func Constructor() MyHashSet {
	size := 20011
	return MyHashSet{buckets: make([][]int, size), size: size}
}

func (hs *MyHashSet) hash(key int) int {
	hash := key % hs.size
	if hash < 0 {
		return hs.size + hash
	}

	return hash
}

func (hs *MyHashSet) Add(key int) {
	h := hs.hash(key)
	for _, v := range hs.buckets[h] {
		if v == key {
			return
		}
	}

	hs.buckets[h] = append(hs.buckets[h], key)
}

func (hs *MyHashSet) Remove(key int) {
	h := hs.hash(key)
	for i, v := range hs.buckets[h] {
		if v == key {
			last := len(hs.buckets[h]) - 1
			hs.buckets[h][i] = hs.buckets[h][last]
			hs.buckets[h] = hs.buckets[h][:last]
			return
		}
	}
}

func (hs *MyHashSet) Contains(key int) bool {
	h := hs.hash(key)
	for _, v := range hs.buckets[h] {
		if v == key {
			return true
		}
	}

	return false
}
