package lrucache

type Pair struct {
	Key int
	Val int
}

type ListNode struct {
	Pair Pair
	Prev *ListNode
	Next *ListNode
}

type LRUCache struct {
	cache    map[int]*ListNode
	capacity int
	head     *ListNode
	tail     *ListNode
}

func Constructor(capacity int) LRUCache {
	c := LRUCache{make(map[int]*ListNode, capacity), capacity, &ListNode{}, &ListNode{}}
	c.head.Next = c.tail
	c.tail.Prev = c.head
	return c
}

func (c *LRUCache) remove(node *ListNode) {
	prev := node.Prev
	next := node.Next

	prev.Next = next
	next.Prev = prev
}

func (c *LRUCache) addToHead(node *ListNode) {
	next := c.head.Next
	node.Next = next
	node.Prev = c.head
	c.head.Next = node
	next.Prev = node
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.cache[key]; ok {
		c.remove(node)
		c.addToHead(node)
		return node.Pair.Val
	}

	return -1
}

func (c *LRUCache) Put(key, val int) {
	if node, ok := c.cache[key]; ok {
		node.Pair.Val = val
		c.remove(node)
		c.addToHead(node)
	} else {
		if len(c.cache) == c.capacity {
			oldNode := c.tail.Prev
			delete(c.cache, oldNode.Pair.Key)
			c.remove(oldNode)
		}
		newNode := &ListNode{Pair: Pair{Key: key, Val: val}}
		c.cache[key] = newNode
		c.addToHead(newNode)
	}
}
