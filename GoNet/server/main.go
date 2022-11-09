package main

import (
	"bufio"
	"fmt"
	"net"
)

// TCP server端
// 一个TCP服务端可以同时连接很多个客户端，例如世界各地的用户使用自己电脑上的浏览器访问淘宝网。
// 因为Go语言中创建多个goroutine实现并发非常方便和高效，所以我们可以每建立一次链接就创建一个goroutine去处理
// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		recvStr = "recev:" + recvStr
		fmt.Println("收到client端发来的数据 ：", recvStr)
		conn.Write([]byte(recvStr)) // 发送数据
	}
}

func main() {
	// 	TCP服务端程序的处理流程：

	// 监听端口
	// 接收客户端请求建立链接
	// 创建goroutine处理链接。
	listen, err := net.Listen("tcp", "127.0.0.1:20000") //监听端口信息
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接，接收客户端请求连接信息
		fmt.Println(conn.RemoteAddr().String())
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接，创建goroutine处理链接

	}
}
