package command

import (
	"fmt"
	"strconv"
	"testing"
)

func TestGet(t *testing.T) {
	number := 123
	byt := []rune(strconv.FormatInt(int64(number), 10))
	fmt.Println(byt)
}
