package command

import (
	"errors"
	"fmt"
	"modis/object"
)

type SdsArr []object.Sds

var (
	kv     map[string]*object.Sds
	cmdArr SdsArr
)

func KVGet(key string) (*object.Sds, bool) {

	val, ok := kv[key]
	//fmt.Printf("key:%v, value: %v,ok %v", key, val, ok)
	return val, ok
}

func KVSet(key string, value *object.Sds) {
	kv[key] = value
}

func Keys() {
	for k := range kv {
		fmt.Println(k)
		//fmt.Println(string(v.Data))
	}
}
func init() {
	kv = make(map[string]*object.Sds)
}
func Get(cmd [][]byte) (string, error) {
	if len(cmd) != 2 {
		return "", errors.New("输入错误，请输入一个 key")
	}
	key := string(cmd[1])
	value, ok := KVGet(key)
	if !ok {
		return "", errors.New("key 不存在")
	}
	return string(value.Data), nil
}

func Set(cmd [][]byte) (string, error) {
	if len(cmd) != 3 || cmd[1] == nil || len(cmd[1]) < 1 || cmd[2] == nil || len(cmd[2]) < 1 {
		return "", errors.New("输入错误，请输入一个 key, 一个 value")
	} else {
		key := string(cmd[1])
		val := object.BytesToSds(cmd[2])
		KVSet(key, val)
		//fmt.Printf("key:%v,val:% v", key, val)
	}
	return "ok", nil
}
