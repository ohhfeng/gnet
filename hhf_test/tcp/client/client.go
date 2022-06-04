package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "10.226.10.69:9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	n, err := conn.Write([]byte("hello"))
	fmt.Println(n, err)
	bys := make([]byte, 0, 10)
	read, err := conn.Read(bys)
	fmt.Println(read, err)
	fmt.Println(bys)
	//server := common.EchoServer{}
	//
	//client, err := gnet.NewClient(&server)
	//if err != nil {
	//	panic(err)
	//}
	//conn, err := client.Dial("tcp", "10.226.10.69:9000")
	//if err != nil {
	//	panic(err)
	//}
	//defer conn.Close()
	////go func() {
	//conn.AsyncWrite([]byte("Hello"))
	////}()
	//readBytes := conn.Read()
	//fmt.Println(readBytes)
}
