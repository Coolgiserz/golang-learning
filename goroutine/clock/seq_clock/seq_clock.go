package main

import (
	"io"
	"log"
	"net"
	"time"
)

//程序目标：构建一个时钟服务器，每秒钟向客户端发送当前时间
// 该服务器顺序执行，一次只能处理一个请求
func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
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
		handleConn(conn) //处理连接

	}
}

func handleConn(c net.Conn) {
	defer c.Close() //延迟调用，确保处理完或者该函数退出时连接会关闭
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
