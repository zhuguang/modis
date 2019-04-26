package command

import (
	"bytes"
	"errors"
	"fmt"
	"modis/object"
	"modis/utils"
	"strconv"
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

func SAdd(cmd [][]byte) (string, error) {
	if len(cmd) != 3 || cmd[1] == nil || len(cmd[1]) < 1 || cmd[2] == nil || len(cmd[2]) < 1 {
		return "", errors.New("输入错误，请输入sadd [key] [数字]")
	}
	key := string(cmd[1])
	num := bytes.Runes(cmd[2])
	if !utils.RunesIsDigit(num) {
		return "", errors.New("值必须是数字")
	}

	val, ok := KVGet(key)
	if !ok {
		return "", errors.New("key 不存在")
	}
	if !utils.RunesIsDigit(val.Data) {
		return "", errors.New("原有值不是数字")
	}
	data := val.Data
	dataDigit, err := strconv.Atoi(string(data))
	if err != nil {
		return "", err
	}
	addNum, err := strconv.Atoi(string(num))
	if err != nil {
		return "", err
	}
	dataDigit += addNum
	result := []rune(strconv.FormatInt(int64(dataDigit), 10))
	KVSet(key, object.RunesToSds(result))
	return "ok", nil
}
