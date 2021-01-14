package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"../proto"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {

		reader := bufio.NewReader(conn)
		msg, err := proto.Decode(reader)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Read Error. Err: ", err)
		}

		if msg == "q" {
			break
		}
		fmt.Println(msg)
	}

}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:5000")
	if err != nil {
		fmt.Println("Start TCP Server Failed. err: ", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Get Connected Failed. err: ", err)
		}
		go process(conn)
	}
}

// func process(conn net.Conn) {
// 	defer conn.Close()
// 	reader := bufio.NewReader(conn)
// 	for {
// 		msg, err := proto.Decode(reader)
// 		if err == io.EOF {
// 			return
// 		}
// 		if err != nil {
// 			fmt.Println("decode msg failed, err:", err)
// 			return
// 		}
// 		fmt.Println("收到client发来的数据：", msg)
// 	}
// }

// func main() {

// 	listen, err := net.Listen("tcp", "127.0.0.1:30000")
// 	if err != nil {
// 		fmt.Println("listen failed, err:", err)
// 		return
// 	}
// 	defer listen.Close()
// 	for {
// 		conn, err := listen.Accept()
// 		if err != nil {
// 			fmt.Println("accept failed, err:", err)
// 			continue
// 		}
// 		go process(conn)
// 	}
// }
