package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

//程序目标：构建一个时钟服务器，接收端口号作为参数，每秒钟向客户端发送当前时间
// 该服务器并发执行，一次能处理多个请求
func main() {
	args := os.Args
	if len(args) < 2 {
		panic("Must specify the PORT!")
	}
	port := args[1]
	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		//Accept()会阻塞直到收到连接请求
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) //Print, Printf, Println
			continue
		}
		go handleConn(conn, port) //并发处理连接：通过go关键字启用goroutine

	}
}

func handleConn(c net.Conn, port string) {
	defer c.Close() //延迟调用，确保处理完或者该函数退出时连接会关闭
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n")+" from "+port+"\n")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
