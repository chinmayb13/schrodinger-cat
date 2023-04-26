package linkedlist3

/*
Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and set.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
set(key, value) - Set or insert the value if the key is not already present. When the cache reaches its capacity, it should invalidate the least recently used item before inserting the new item.
The LRUCache will be initialized with an integer corresponding to its capacity. Capacity indicates the maximum number of unique keys it can hold at a time.

Definition of "least recently used" : An access to an item is defined as a get or a set operation of the item. "Least recently used" item is the one with the oldest access time.

NOTE: If you are using any global variables, make sure to clear them in the constructor.

Example :

Input :

	capacity = 2
	set(1, 10)
	set(5, 12)
	get(5)        returns 12
	get(1)        returns 10
	get(10)       returns -1
	set(6, 14)    this pushes out key = 5 as LRU is full.
	get(5)        returns -1
*/
type dListNode struct {
	value int
	next  *dListNode
	prev  *dListNode
}

func dListNode_new(val int) *dListNode {
	var node *dListNode = new(dListNode)
	node.value = val
	return node
}

type LRUCache struct {
	capacity int
	head     *dListNode
	tail     *dListNode
	hashMap  map[int]int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		hashMap:  make(map[int]int),
	}
}

func (this *LRUCache) Get(key int) int {
	val, ok := this.hashMap[key]
	if ok {
		//if key is found, move the current key node to the tail
		node := dListNode_new(key)
		this.handleExistingKey(node)
		return val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node := dListNode_new(key)
	_, ok := this.hashMap[key]
	if ok {
		//if key is found, move the current key node to the tail
		this.handleExistingKey(node)
	} else if len(this.hashMap) == 0 {
		//if the map is empty
		this.head, this.tail = node, node
	} else {
		//update the tail
		node.prev = this.tail
		this.tail.next = node
		this.tail = this.tail.next
		//if map is at capacity
		if len(this.hashMap) == this.capacity {
			delete(this.hashMap, this.head.value)
			this.head = this.head.next
			this.head.prev.next = nil
			this.head.prev = nil

		}
	}
	this.hashMap[key] = value
}

func (lru *LRUCache) handleExistingKey(node *dListNode) {
	//if the given key is already present at tail, do nothing
	//if not, then perform below
	if lru.tail.value != node.value {
		//if key is present at head
		if lru.head.value == node.value {
			lru.head = lru.head.next
			lru.head.prev.next = nil
			lru.head.prev = nil
		} else {
			//else iterate and find
			cur := lru.head
			for cur != nil {
				if cur.value == node.value {
					cur.prev.next = cur.next
					cur.next.prev = cur.prev
					break
				}
				cur = cur.next
			}
		}
		//update tail
		node.prev = lru.tail
		lru.tail.next = node
		lru.tail = lru.tail.next
	}
}
