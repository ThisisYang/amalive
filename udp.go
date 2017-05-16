package main

import (
	//"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
	//"unicode/utf8"
)

func main() {
	one := []byte{}
	conn, err := net.Dial("ip4:1", "127.0.0.3")
	if err == nil {
		defer conn.Close()
		conn.SetReadDeadline(time.Now().Add(10 * time.Millisecond))
		data, err := conn.Read(one)
		if err == io.EOF {
			log.Fatal("err = EOF")
		}
		fmt.Println("data:", data)
		os.Exit(0)
	}

	//
	//status, _ := bufio.NewReader(conn).ReadString('\n')
	//buf := make([]byte, 1000000)
	//var n = 1
	//for n >= 1 {
	//n, _ = conn.Read(buf)
	//r, size := utf8.DecodeRune(buf)
	//fmt.Printf(string(buf))
	//}
	fmt.Println(err)
	log.Fatal("connection err")
}
