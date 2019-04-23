package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var kv map[string]string

func main() {
	kv = make(map[string]string)
	fmt.Printf("modis>")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		re, _ := regexp.Compile("\\s+")
		line = string(re.ReplaceAll([]byte(line), []byte(" ")))
		cmd := strings.Split(line, " ")
		var err error
		var value string
		if cmd == nil || len(cmd) < 2 {
			fmt.Printf("modis>输入错误，请输入 set 或 get 指令\n")
		} else if cmd[0] == "set" {
			err = cmdSet(cmd)
		} else if cmd[0] == "get" {
			value, err = cmdGet(cmd)
			fmt.Printf("modis>%s\n", value)
		}
		if err != nil {
			fmt.Printf("modis>" + err.Error() + "\n")
		}
		fmt.Printf("modis>")
	}
}

func cmdGet(cmd []string) (string, error) {
	if len(cmd) != 2 {
		return "", errors.New("输入错误，请输入一个 key")
	}
	value, ok := kv[cmd[1]]
	if !ok {
		return "", errors.New("key 不存在")
	}
	return value, nil
}

func cmdSet(cmd []string) error {
	if len(cmd) != 3 || cmd[1] == "" || cmd[2] == "" {
		return errors.New("输入错误，请输入一个 key, 一个 value")
	} else {
		kv[cmd[1]] = cmd[2]
	}
	return nil
}
