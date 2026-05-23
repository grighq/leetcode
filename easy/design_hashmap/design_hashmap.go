package designhashmap

type Pair struct {
	key int
	val int
}

type MyHashMap struct {
	buckets [][]Pair
	size    int
}

func Constructor() MyHashMap {
	size := 20011
	return MyHashMap{make([][]Pair, size), size}
}

func (hm *MyHashMap) hash(key int) int {
	hash := key % hm.size
	if hash < 0 {
		return hash + hm.size
	}

	return hash
}

func (hm *MyHashMap) Put(key, val int) {
	h := hm.hash(key)
	for i, v := range hm.buckets[h] {
		if v.key == key {
			hm.buckets[h][i].val = val
			return
		}
	}

	hm.buckets[h] = append(hm.buckets[h], Pair{key, val})
}

func (hm *MyHashMap) Get(key int) int {
	h := hm.hash(key)
	for _, v := range hm.buckets[h] {
		if v.key == key {
			return v.val
		}
	}

	return -1
}

func (hm *MyHashMap) Remove(key int) {
	h := hm.hash(key)
	for i, v := range hm.buckets[h] {
		if v.key == key {
			last := len(hm.buckets[h]) - 1
			hm.buckets[h][i] = hm.buckets[h][last]
			hm.buckets[h] = hm.buckets[h][:last]
			return
		}
	}
}
