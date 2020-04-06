package lru

import "algorithm/linkedList/double"

type LRUCache struct {
	Size int
	Data map[string]*double.DoubleNode
	Link *double.Double
}

func Constructor(capacity int) *LRUCache {
	return &LRUCache{capacity, make(map[string]*double.DoubleNode), &double.Double{}}
}

func (this *LRUCache) Get(key string) interface{} {

	if _, ok := this.Data[key]; !ok {
		return nil
	}
	node := this.Data[key]
	this.Link.UNShiftNode(node)
	return node.Value

}

func (this *LRUCache) Put(key string, value interface{}) {
	//
	node, ok := this.Data[key]
	if ok {
		node.Value.Value = value
		this.Link.RemoveNode(node)
		this.Link.UNShiftNode(node)
	} else {
		node = &double.DoubleNode{double.Ele{key, value}, nil, nil}
		this.Data[key] = node
		if len(this.Data) > this.Size {
			key := this.Link.Tail.Value.Key
			delete(this.Data, key)
			this.Link.RemoveNode(this.Link.Tail)
		}
		this.Link.UNShiftNode(node)

	}

}
