package lock_free_queue

import (
	"sync/atomic"
	"unsafe"
)

// lock-free的queue
type LKQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

// 通过链表实现，这个数据结构代表链表中的节点
type node struct {
	value interface{}
	next  unsafe.Pointer
}

func NewLKQueue() *LKQueue {
	n := unsafe.Pointer(&node{})
	return &LKQueue{head: n, tail: n} // 头尾都是 node指针
}

// 入队
func (q *LKQueue) Enqueue(v interface{}) {
	n := &node{value: v}
	for {
		tail := load(&q.tail)      // 原子的读取到内存   目前为空node     队列queue的尾
		next := load(&tail.next)   // 看尾的下一个 （node struct）       尾node 的 next
		if tail == load(&q.tail) { // 尾还是尾
			if next == nil { // 还没有新数据入队
				if cas(&tail.next, next, n) { //  增加到队尾    原子替换  node.next  =》  n     node 的next 是n      给node的 next赋值
					cas(&q.tail, tail, n) //入队成功，移动尾巴指针   队列的尾指针  =》 n   队列加入n    head还是原来的node   tail 是 n   给queue的尾赋值
					return
				}
			} else { // 已有新数据加到队列后面，需要移动尾指针
				cas(&q.tail, tail, next) // 队列 尾指针  =》 尾节点 node 的next 属性     给queue的尾赋值 为  node的next  他的next又为空
			}
		}
	}
}

// 出队，没有元素则返回nil
func (q *LKQueue) Dequeue() interface{} {
	for {
		head := load(&q.head)      //
		tail := load(&q.tail)      //  有
		next := load(&head.next)   //  头队列的next 是有的
		if head == load(&q.head) { // head还是那个head
			if head == tail { // head和tail一样
				if next == nil { // 说明是空队列   //
					return nil
				}
				// 只是尾指针还没有调整，尝试调整它指向下一个
				cas(&q.tail, tail, next)
			} else {
				// 读取出队的数据
				v := next.value //    node.value
				// 既然要出队了，头指针移动到下一个
				if cas(&q.head, head, next) { // head 指向 head。next  也是一个node
					return v // Dequeue is done.  return
				}
			}
		}
	}
}

// 将unsafe.Pointer原子加载转换成node
func load(p *unsafe.Pointer) (n *node) {
	return (*node)(atomic.LoadPointer(p))
}

// 封装CAS,避免直接将*node转换成unsafe.Pointer
func cas(p *unsafe.Pointer, old, new *node) (ok bool) {
	return atomic.CompareAndSwapPointer(
		p, unsafe.Pointer(old), unsafe.Pointer(new))
}
