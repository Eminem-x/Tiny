package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
	"tinyRpc"
	"tinyRpc/codec"
)

func main() {
	// 设置选项可在每条输出的文本前增加一些额外信息，如日期时间、文件名等
	// 可以参考: https://darjun.github.io/2020/02/07/godailylib/log/
	log.SetFlags(0)
	addr := make(chan string)
	go startServer(addr)

	// in fact, following code is like a simple tinyRpc client
	conn, _ := net.Dial("tcp", <-addr)
	defer func() { _ = conn.Close() }()

	time.Sleep(time.Second)
	// send options
	_ = json.NewEncoder(conn).Encode(tinyRpc.DefaultOption)
	cc := codec.NewGobCodec(conn)

	// send request & receive response
	// 有序请求
	for i := 0; i < 5; i++ {
		mockConnect(uint64(i), cc)
	}
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

func mockConnect(seqID uint64, cc codec.Codec) {
	h := &codec.Header{
		ServiceMethod: "Foo.Sum",
		Seq:           seqID,
	}
	_ = cc.Write(h, fmt.Sprintf("tinyRpc req %d", h.Seq))
	_ = cc.ReadHeader(h)
	var reply string
	_ = cc.ReadBody(&reply)
	log.Println("reply:", reply)
}
