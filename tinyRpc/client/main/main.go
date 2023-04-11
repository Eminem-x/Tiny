package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
	"tinyRpc"
)

func main() {
	// 设置选项可在每条输出的文本前增加一些额外信息，如日期时间、文件名等
	// 可以参考: https://darjun.github.io/2020/02/07/godailylib/log/
	log.SetFlags(0)
	addr := make(chan string)
	go startServer(addr)
	client, _ := tinyRpc.Dial("tcp", <-addr)
	defer func() { _ = client.Close() }()

	time.Sleep(time.Second)
	// send request & receive response
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := fmt.Sprintf("tinyRpc req %d", i)
			var reply string
			if err := client.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Println("reply:", reply)
		}(i)
	}
	wg.Wait()
}

func startServer(addr chan string) {
	// pick a free port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	// 使用信道 addr, 确保服务端端口监听成功, 客户端再发起请求
	addr <- l.Addr().String()
	tinyRpc.Accept(l)
}
