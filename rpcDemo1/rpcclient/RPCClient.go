package main

import (
	"net/rpc"
	"log"
	"fmt"
)

//实现客户端

type Param struct {
	Width,Height int

}

func main()  {
	//链接远程RPC
	rp,err := rpc.DialHTTP("tcp","127.0.0.1:9090")
	if err != nil {
		log.Fatal(err)
	}
    //从服务器接收过来的值
	ret := 0
	e := rp.Call("Rect.Area",Param{100,100},&ret)
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(ret)

}