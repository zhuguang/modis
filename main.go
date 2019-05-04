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
		fmt.Println(len(cmd))
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
		fmt.Println("准备")
		conn, _ := listener.Accept()
		go handleMessage(conn)
		//conn.Close()
	}
}

func handleCommand(cmd [][]byte) ([]string, error) {
	//fmt.Printf("len:%v,%v,%v", len(cmd), len(cmd[len(cmd)-1]), cmd[len(cmd)-1])
	fmt.Println("keys" + string(cmd[0]))
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
		fmt.Println("handle..")
		num, err := conn.Read(line)
		re, _ := regexp.Compile("\\s+")
		line = re.ReplaceAll(line, []byte(" "))

		fmt.Printf("收到命令：%s", line)
		cmd := bytes.Split(line, []byte(" "))
		fmt.Println(len(cmd))
		cmd, v := validCmd(cmd)
		if v {
			continue
		}
		fmt.Println(len(cmd))
		response, errCmd := handleCommand(cmd)
		if response != nil {
			fmt.Println("ifrespon")
			printConn(response, conn)
		} else {
			fmt.Println("elseconn")
			conn.Write([]byte(" 1\r\n"))
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

//这里二维数组用指针修改不生效？
func validCmd(cmd [][]byte) ([][]byte, bool) {
	last := cmd[len(cmd)-1]
	//fmt.Printf("%v", last)
	if last == nil || string(last) == "" {
		cmd = cmd[0 : len(cmd)-1]
		fmt.Println("if?")
		fmt.Println(len(cmd))
		return cmd, false
	}
	return cmd, false
}

func printConn(rep []string, conn net.Conn) {
	fmt.Printf("printConn%v", len(rep))
	for _, info := range rep {
		fmt.Println(info)
		conn.Write([]byte(info))
	}
}
