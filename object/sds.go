package object

import "bytes"

//redis sds 中用c语言的 char 数组，c中 char 占一个字节（一个字节怎么能表示中文？）
//go 中 rune表示一个unicode码点，占4字节，虽然这样有些浪费内存，但可以利用 go 内置的可以处理rune的函数
//这里用 go 自带slice，自动扩容所以 Len,Alloc 可以不用
type Sds struct {
	Len   int
	Alloc int
	Data  []rune
}

func BytesToSds(arr []byte) *Sds {
	len := len(arr)
	data := bytes.Runes(arr)
	return &Sds{len, len, data}
}

func BytesToString(s []byte) string {
	return string(s)
}
