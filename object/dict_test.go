package object

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestDictCreate(t *testing.T) {
	dict := dictCreate(&dictT{})
	dict.dictAdd("1","haha")
	fmt.Println(dict)
}

type dictT struct {
}

func (d *dictT) hashFunc(k interface{}) int64 {
	return rand.Int63()
}
func (d *dictT) keyCompare(key1 interface{}, key2 interface{}) int {
	return 0
}
