package main

import (
	"fmt"
	"net"
	"sync"

	"../proto"
)

var wg sync.WaitGroup

// 好像只能收到几条？ 没办法收到一个goroutine发送的所有消息？？

func sendMsg(i int) {
	defer wg.Done()
	conn, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		fmt.Println("Connect Failed. Err: ", err)
	}
	defer conn.Close()

	// 4条消息
	msgList := []string{fmt.Sprintf("我是%d号通信，这是我的第一条信息", i), fmt.Sprintf("我是%d号通信，这是我的第二条信息", i), fmt.Sprintf("我是%d号通讯，这是我的第三条信息", i), "q"}
	// msgList := []string{fmt.Sprintf("我是%d号通信，这是我的第一条信息", i), "q"}

	// 遍历每条消息
	for _, msg := range msgList {
		msgInfo, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("Encode Msg Failed. Err: ", err)
		}

		_, err = conn.Write(msgInfo)
		if err != nil {
			fmt.Println("Send Msg Failed. Err: ", err)
		}
		fmt.Println("SEND: ", string(msg))

	}
}

func main() {
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go sendMsg(i)
	}

	wg.Wait()
}

// func main() {
// 	conn, err := net.Dial("tcp", "127.0.0.1:30000")
// 	if err != nil {
// 		fmt.Println("dial failed, err", err)
// 		return
// 	}
// 	defer conn.Close()
// 	for i := 0; i < 20; i++ {
// 		msg := `Hello, Hello. How are you?`
// 		data, err := proto.Encode(msg)
// 		if err != nil {
// 			fmt.Println("encode msg failed, err:", err)
// 			return
// 		}
// 		conn.Write(data)
// 	}
// }
