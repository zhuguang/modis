package object

const (
	ITE_DIR = 1
)

type listNode struct {
	pre   *listNode
	next  *listNode
	value interface{}
}

type list struct {
	head *listNode
	tail *listNode
	len  uint64
}

type listIterator struct {
	next      *listNode
	direction int8
}

func (ite *listIterator) getNext() *listNode {
	temp := ite.next
	ite.next = temp.next
	return temp
}
func (ite *listIterator) hasNext() bool {
	return ite.next != nil
}

func listCreate() *list {
	return &list{nil, nil, 0}
}

func (l *list) listAddNodeTail(val interface{}) {
	tailP := l.tail
	if tailP == nil {
		tailNode := &listNode{nil, nil, val}
		l.head = tailNode
		l.tail = tailNode
	} else {
		newNodeP := &listNode{tailP, nil, val}
		tailP.next = newNodeP
		l.tail = newNodeP
	}
	l.len += 1
}

func (l *list) listAddNodeHead(val interface{}) {
	headP := l.head
	if headP == nil {
		headNode := &listNode{nil, nil, val}
		l.head = headNode
		l.tail = headNode
	} else {
		newNodeP := &listNode{nil, headP, val}
		l.head = newNodeP
		headP.pre = l.head
	}
	l.len ++
}

func (l *list) listLength() uint64 {
	return l.len
}

func (l *list) listFirst() *listNode {
	return l.head
}

func (l *list) listLast() *listNode {
	return l.tail
}

func (l *list) listPrevNode(node *listNode) *listNode {
	return node.pre
}

func (l *list) listNextNode(node *listNode) *listNode {
	return node.next
}

func (l *list) listDelNode(node *listNode) {
	if node.pre != nil {
		node.pre.next = node.next
	} else {
		l.head = node.next
	}
	if node.next != nil {
		node.next.pre = node.pre
	} else {
		l.tail = node.pre
	}
	l.len --
}

// direction 迭代器方向：1从头开始，否则从尾开始
func (l *list) listIteratorGen(direction int8) *listIterator {
	listIterator := &listIterator{nil, direction}
	if direction == ITE_DIR {
		listIterator.next = l.head
	} else {
		listIterator.next = l.head
	}
	return listIterator
}

//todo 线程不安全
func (l *list) listIndex(index uint64) *listNode {
	listIterator := l.listIteratorGen(ITE_DIR)
	a := 0
	for listIterator.hasNext() {
		node := listIterator.getNext()
		if uint64(a) == index {
			return node
		}
		a++
	}
	return nil
}
