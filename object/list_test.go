package object

import (
	"fmt"
	"testing"
)

func TestListCreate(t *testing.T) {
	l := listCreate()
	l.listAddNodeTail("hello1")
	l.listAddNodeTail(2)

	fmt.Println(l)
}

func TestListAddNodeHead(t *testing.T) {
	l := listCreate()
	l.listAddNodeTail("tail1")
	l.listAddNodeTail("tail2")
	l.listAddNodeHead("head2")
	l.listAddNodeHead("head1")
	fmt.Printf("%v", l.listLength())

}

func TestListIndex(t *testing.T) {
	l := listCreate()
	l.listAddNodeTail("tail1")
	l.listAddNodeTail("tail2")
	l.listAddNodeHead("head2")
	l.listAddNodeHead("head1")
	n0 := l.listIndex(0)
	n1 := l.listIndex(1)
	n2 := l.listIndex(2)
	n3 := l.listIndex(3)
	fmt.Printf("0:%v,1:%v,2:%v,3：%v", n0, n1, n2, n3)
	l.listDelNode(n0)
	l.listDelNode(n2)
	fmt.Printf("%v",l)
}
