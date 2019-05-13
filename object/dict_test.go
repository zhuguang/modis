package object

import (
	"fmt"
	"testing"
)

func TestDictCreate(t *testing.T) {
	dict := dictCreate(&dictT{})
	dict.dictAdd("1","haha")
	dict.dictReplace("1","haha1")
	fmt.Println(dict)
}

type dictT struct {
}

func (d *dictT) hashFunc(k interface{}) int64 {
	return 100111
}
func (d *dictT) keyCompare(key1 interface{}, key2 interface{}) int {
	return 0
}
