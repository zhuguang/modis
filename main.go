package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"modis/command"
	"net"
	"os"
	"regexp"
	"strings"
)

const (
	Pre = "modis>"
)

func main() {
	go initServer()
	initLocalTerminal()
}

//初始化本地终端
func initLocalTerminal() {
	fmt.Printf(Pre)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Bytes()
		re, _ := regexp.Compile("\\s+")
		line = re.ReplaceAll(line, []byte(" "))
		cmd := bytes.Split(line, []byte(" "))
		//fmt.Println(len(cmd))
		cmd, v := validCmd(cmd)
		if v {
			continue
		}
		fmt.Println(len(cmd))
		var err error
		var response []string
		response, err = handleCommand(cmd)
		if response != nil {
			printTerminal(response)
		}
		if err != nil {
			fmt.Printf(err.Error() + "\n")
		}
		fmt.Printf(Pre)
	}
}

func printTerminal(rep []string) {
	for _, info := range rep {
		fmt.Printf("%s\n", info)
	}
}

//初始化网络服务器
func initServer() {
	listener, _ := net.Listen("tcp", "127.0.0.1:10001")
	defer listener.Close()
	for {
		fmt.Println("准备...")
		conn, _ := listener.Accept()
		go handleMessage(conn)
		//conn.Close()
	}
}

func handleCommand(cmd [][]byte) ([]string, error) {
	//fmt.Printf("len:%v,%v,%v", len(cmd), len(cmd[len(cmd)-1]), cmd[len(cmd)-1])
	var err error
	var value string
	var ret []string
	if cmd == nil || len(cmd) < 1 {
		ret = []string{"输入错误，请输入 set 或 get 指令\n"}
	} else if bytes.EqualFold(cmd[0], []byte("set")) {
		value, err = command.Set(cmd)
	} else if bytes.EqualFold(cmd[0], []byte("get")) {
		value, err = command.Get(cmd)
	} else if bytes.EqualFold(cmd[0], []byte("keys")) {
		fmt.Println("to keys")
		ret = command.Keys()
	} else if bytes.EqualFold(cmd[0], []byte("incr")) {
		value, err = command.Incr(cmd)
	}
	if ret == nil || len(ret) == 0 {
		ret = []string{value}
	}
	return ret, err
}

func handleMessage(conn net.Conn) {
	for {
		line := make([]byte, 1000)
		num, err := conn.Read(line)
		//可以输出实际字符，可以输出\n \r\n \u2318 等等这种特殊字符
		//fmt.Printf("%+q",line)
		//接收到的字节数组，不足自定义字节数组长度的会被补\u0000，也就是unicode 0x00
		re, _ := regexp.Compile("[\\s\u0000]+")
		line = re.ReplaceAll(line, []byte(" "))
		cmd := bytes.Split(line, []byte(" "))
		cmd, v := validCmd(cmd)
		if v {
			continue
		}
		response, errCmd := handleCommand(cmd)
		if response != nil {
			printConn(response, conn)
		} else {
			conn.Write([]byte("\n"))
		}
		if errCmd != nil {
			conn.Write([]byte(errCmd.Error() + "\n"))
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			if strings.Contains(err.Error(), "use of closed network connection") {
				log.Fatal("use of closed network connection")
				break
			}
			fmt.Printf("err:%v, %v", err, num)
		}

	}
}

//todo 二维数组指针
func validCmd(cmd [][]byte) ([][]byte, bool) {
	last := cmd[len(cmd)-1]
	if last == nil || bytes.EqualFold(last, nil) {
		cmd = cmd[0 : len(cmd)-1]
		fmt.Println(len(cmd))
		return cmd, false
	}
	return cmd, false
}

func printConn(rep []string, conn net.Conn) {
	for _, info := range rep {
		conn.Write([]byte(info + "\n"))
	}
}
