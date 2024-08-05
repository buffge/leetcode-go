package main

import (
	"fmt"
)

type ListKeyNode struct {
	Key  int
	Val  int
	Prev *ListKeyNode
	Next *ListKeyNode
}

/*
get要求是O(1) 那么底层存储就是数组或者hash
删除要求是O(1) 那么底层就是双向链表或hash
用 双向链表做序列,keyMapPtr做 get 映射
缓存中有 头尾指针,最新被获取或插入的放到尾
1.更新新值时删除头
2.更新旧值时将旧值放到尾
3.get存在时将值放到末尾
*/
type LRUCache struct {
	head, tail *ListKeyNode
	keySet     map[int]*ListKeyNode
}

func Constructor(capacity int) LRUCache {
	head := &ListKeyNode{}
	tail, curr := head, head
	for i := 0; i < capacity-1; i++ {
		curr.Next = &ListKeyNode{Prev: curr}
		curr, tail = curr.Next, curr.Next
	}
	return LRUCache{head: head, tail: tail, keySet: make(map[int]*ListKeyNode, capacity)}
}

func (ce *LRUCache) Get(key int) int {
	p, exist := ce.keySet[key]
	if !exist {
		return -1
	}
	// 将值放到尾部
	if p != ce.tail {
		// 如果值当前在头 将头右移
		if p.Prev == nil {
			if ce.head.Next != nil {
				ce.head = ce.head.Next
				ce.head.Prev = nil
			}
		} else { // 值不在头时将值从链表中移除
			p.Prev.Next = p.Next
			p.Next.Prev = p.Prev
		}
		// 将值移到尾
		p.Prev = ce.tail
		p.Next = nil
		ce.tail.Next = p
		ce.tail = p
		ce.keySet[key] = ce.tail
	}
	return p.Val
}

func (ce *LRUCache) Put(key int, value int) {
	p, exist := ce.keySet[key]
	if exist && ce.tail == p { // 更新时如果存在并且是尾直接更新
		ce.tail.Val = value
		return
	}
	// 如果不存在 或者不在尾 将值移动到尾
	// 不存在时将头移到尾
	newHead, newTail := ce.head, ce.head
	if exist { // 存在并且不在尾时
		if p == ce.head { //  如果p是头
			if ce.head.Next != nil { // 更新新头
				newHead = ce.head.Next
			}
		} else {
			p.Prev.Next = p.Next // 将p先移除链表
			p.Next.Prev = p.Prev
		}
		newTail = p // p 移到尾部
	} else { // 不存在时 删除头元素缓存并将头右移
		delete(ce.keySet, ce.head.Key)
		if ce.head.Next != nil { // 更新新头
			newHead = ce.head.Next
		}
		ce.keySet[key] = ce.head
	}

	ce.head = newHead // 更新头
	ce.head.Prev = nil
	newTail.Key = key // 更新尾
	newTail.Val = value
	newTail.Prev = ce.tail
	ce.tail.Next = newTail
	ce.tail = newTail
	ce.tail.Next = nil
}

// todo 优化代码冗余 组合到函数中
func main() {
	capacity := 3
	obj := Constructor(capacity)
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Put(3, 3)
	obj.Put(4, 4)
	obj.Put(3, 13)
	obj.Put(4, 14)
	obj.Put(2, 12)

	fmt.Println(obj.Get(4))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(1))
	obj.Put(5, 5)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
	fmt.Println(obj.Get(5))

}
