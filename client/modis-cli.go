package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

const (
	Pre = "modis>"
)

type Args struct {
	Host string
	Port string
}

var args Args

func main() {
	flag.StringVar(&args.Host, "h", "127.0.0.1", "Server hostname (default: 127.0.0.1).")
	flag.StringVar(&args.Port, "p", "10001", "Server port (default: 10001).")
	flag.Parse()
	initLocalTerminal()
}

//初始化本地终端
func initLocalTerminal() {
	fmt.Printf(Pre)
	address := args.Host + ":" + args.Port
	conn, err := net.Dial("tcp", address)
	defer conn.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Bytes()
		conn.Write(line)
		ret := make([]byte, 1000)
		_, err := conn.Read(ret)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println(string(ret))
		fmt.Printf(Pre)
	}
}
