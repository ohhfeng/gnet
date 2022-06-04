package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			conn, err := net.Dial("tcp", "127.0.0.1:9000")
			if err != nil {
				panic(err)
			}
			defer conn.Close()
			n, err := conn.Write([]byte("h1111111"))
			fmt.Println(n, err)
			bys := make([]byte, 10)
			read, err := conn.Read(bys)
			fmt.Println(read, err)
			fmt.Println(bys)
		}()

	}

	wg.Wait()
}
