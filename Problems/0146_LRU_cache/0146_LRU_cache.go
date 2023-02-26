package _146_LRU_cache

type LRUCache struct {
	memo map[int]*Node
	head *Node // 头尾虚节点
	tail *Node
	cap  int
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.next, tail.prev = tail, head
	return LRUCache{memo: make(map[int]*Node), head: head, tail: tail, cap: capacity}
}

func (this *LRUCache) Get(key int) int {
	node := this.memo[key]
	if node == nil {
		return -1
	}
	// 更新为最新访问
	this.makeRecently(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if node, isPresent := this.memo[key]; isPresent {
		node.value = value
		// 更新为最新访问
		this.makeRecently(node)
	} else {
		if len(this.memo) >= this.cap {
			// 达到最大容量
			// 移除首个节点
			this.remove(this.head.next)
		}
		this.addLast(&Node{key: key, value: value})
	}
}

func (this *LRUCache) makeRecently(node *Node) {
	// 移除
	this.remove(node)
	// 放到队尾
	this.addLast(node)
}

func (this *LRUCache) remove(node *Node) {
	delete(this.memo, node.key)
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) addLast(node *Node) {
	this.memo[node.key] = node
	prev := this.tail.prev
	// node前置节点
	prev.next = node
	node.prev = prev
	// node后置节点
	node.next = this.tail
	this.tail.prev = node
}

type Node struct {
	prev, next *Node
	key        int
	value      int
}
