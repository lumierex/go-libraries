package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	sfpxmrpc "sfpxm-rpc"
	"sfpxm-rpc/codec"
)

func startServer(addr chan string) {
	//net.Listener()
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Println("startServer net listen err: ", err)
	}
	log.Printf("start rpc server on %s", l.Addr())
	addr <- l.Addr().String()
	sfpxmrpc.Accept(l)
}

func main() {

	addr := make(chan string)
	go startServer(addr)

	conn, _ := net.Dial("tcp", <-addr)
	defer func() {
		_ = conn.Close()
	}()

	// 1. 首先发送Option
	err := json.NewEncoder(conn).Encode(sfpxmrpc.DefaultOption)
	if err != nil {
		log.Println("send option error：", err)
	}
	cc := codec.NewGobCodec(conn)
	for i := 0; i < 5; i++ {
		// 发送消息头和消息体
		h := &codec.Header{
			ServiceMethod: "Foo.Sum",
			// request sequence number
			Seq: uint64(i),
		}
		_ = cc.Write(h, fmt.Sprintf("sfpxm rpc req %d", h.Seq))
		_ = cc.ReadHeader(h)
		var reply string
		_ = cc.ReadBody(&reply)
		log.Println("reply:", reply)
	}
}
