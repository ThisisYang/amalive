package main

import (
	//"bufio"
	"fmt"
	"io"
	//"io/ioutil"
	//"log"
	"net"
	//"os"
	"time"
	//"unicode/utf8"
)

func TCPCheck(IP string, Port int) float64 {
	endpoint := fmt.Sprintf("%s:%v", IP, Port)
	var totalSecond time.Duration
	for i := 0; i<*Retry; i++{
		if i != 0 {
			time.Sleep(time.Duration(*Interval)*time.Second)
		}
		Debug.Printf("try: %v", i)
		one := make([]byte, 1)
		t := time.Now()
		conn, err := net.DialTimeout("tcp", endpoint, time.Duration(*Timeout)*time.Second)
		if err != nil {
			Debug.Println("no connection", err)
			totalSecond += time.Duration(*Timeout)*time.Second
			continue
		}
		t1 := time.Now()
		Debug.Printf("%s:%v connected", IP, Port)		
		Debug.Printf("connection time: %s", time.Now().Sub(t))
		conn.SetReadDeadline(time.Now().Add(time.Duration(*Timeout) * time.Second))
		num, err := conn.Read(one)
		if err == io.EOF {
			Debug.Println("EOF", num)
			totalSecond += time.Duration(*Timeout)*time.Second
			continue
		}
		totalSecond += time.Now().Sub(t)
		Debug.Printf("Read time: %v", time.Now().Sub(t1))
		Debug.Printf("Summary: try: %v, time took: %v, totalSecond: %s", i, time.Now().Sub(t), totalSecond)
		conn.Close()

	}
	return (totalSecond.Seconds())*1000/float64(*Retry)
}


func TCPCheckOnce(IP string, Port int) int {
	endpoint := fmt.Sprintf("%s:%v", IP, Port)

	one := make([]byte, 1)
	conn, err := net.DialTimeout("tcp", endpoint, 1*time.Second)
	Debug.Println("here")
	if err == nil {
		Debug.Println("connected")
		defer conn.Close()
		conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		num, err := conn.Read(one)
		if err == io.EOF {
			Debug.Println("no read")
			return 0
		}
		Debug.Println("read")
		return num
	} else {
		Debug.Println("not connected")
		return 0
	}

}

func TCPCheckForever(IP string, Port int, interval int) int {
	endpoint := fmt.Sprintf("%s:%v", IP, Port)
	//fmt.Println(endpoint)
	for {
		one := make([]byte, 1)
		time.Sleep(1 * time.Second)
		conn, err := net.DialTimeout("tcp", endpoint, 1*time.Second)
		// t := time.Now()
		if err == nil {
			defer conn.Close()
			//fmt.Fprint(conn, "GET / HTTP/1.1\r\nHOST:localhost:8001\r\n\r\n")
			conn.SetReadDeadline(time.Now().Add(1 * time.Second))
			//data, err := ioutil.ReadAll(conn)
			num, err := conn.Read(one)
			if err == io.EOF {
				return num
			}
			return num
			//fmt.Printf("%s", data)
		} else {
			return 0
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
