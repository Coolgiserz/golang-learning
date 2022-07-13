package main

import (
	"io"
	"log"
	"net"
	"os"
)

// TCP客户端程序
// 连接TCP服务器,从中读取输出并写到标准输出中，直到达到EOF或者出错

func main() {
	conn, err := net.Dial("tcp", "localhost:8080") //如果对应地址没有服务，则会报错"connect: connection refused"
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
