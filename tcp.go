package amalive

import (
	//"bufio"
	"fmt"
	"io"
	//"io/ioutil"
	"log"
	"net"
	//"os"
	"time"
	//"unicode/utf8"
)

func TCPCheck() {
	one := []byte{}
	for {
		time.Sleep(1 * time.Second)
		conn, err := net.DialTimeout("tcp", "216.200.116.24:22", 500*time.Millisecond)
		t := time.Now()
		if err == nil {
			defer conn.Close()
			//fmt.Fprint(conn, "GET / HTTP/1.1\r\nHOST:localhost:8001\r\n\r\n")
			conn.SetReadDeadline(time.Now().Add(10 * time.Millisecond))
			//data, err := ioutil.ReadAll(conn)
			_, err := conn.Read(one)
			if err == io.EOF {
				log.Fatal("err = EOF")
			}
			fmt.Println("connected", t)
			//fmt.Printf("%s", data)
		} else {
			fmt.Println(err, t)
			//log.Fatal("connection err")
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

		//when calling log.Fatel, os.Exit(1) happend

	}

}
