package main

import (
	"bufio"
	"bytes"
	"fmt"
	"modis/command"
	"os"
	"regexp"
)

const (
	Pre = "modis>"
)

func main() {
	fmt.Printf(Pre)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Bytes()
		re, _ := regexp.Compile("\\s+")
		line = re.ReplaceAll(line, []byte(" "))
		cmd := bytes.Split(line, []byte(" "))
		//fmt.Printf("cmd:%v", cmd)
		var err error
		var value string
		if cmd == nil || len(cmd) < 1 {
			fmt.Printf("输入错误，请输入 set 或 get 指令\n")
		} else if bytes.EqualFold(cmd[0], []byte("set")) {
			value, err = command.Set(cmd)
		} else if bytes.EqualFold(cmd[0], []byte("get")) {
			value, err = command.Get(cmd)
		} else if bytes.EqualFold(cmd[0], []byte("keys")) {
			command.Keys()
		} else if bytes.EqualFold(cmd[0], []byte("incr")) {
			value, err = command.Incr(cmd)
		}
		if value != "" {
			fmt.Printf("%s\n", value)
		}
		if err != nil {
			fmt.Printf(err.Error() + "\n")
		}
		fmt.Printf(Pre)
	}
}

type dict struct {
}
