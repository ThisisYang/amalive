package main

import (
	"fmt"

	"net"
	"time"
)

type checkInfo struct {
	endpoint  string
	retry     int
	retryIntv int
	timeout   int
}

// NewCheckInfo return a checkInfo struct
func NewCheckInfo(ip string, port, timeout, retry, retryIntv int) *checkInfo {
	endpoint := fmt.Sprintf("%s:%v", ip, port)
	if (retry == 0) && (retryIntv > 0) {
		panic("retry interval is set but retry is 0")
	}
	return &checkInfo{
		endpoint:  endpoint,
		retry:     retry,
		retryIntv: retryIntv,
		timeout:   timeout,
	}
}

func (c *checkInfo) tcpCheck() {
	var totalSec time.Duration
	for i := 1; i <= c.retry+1; i++ {
		t := time.Now()
		succ := checkOnce(c.endpoint, c.timeout)
		totalSec += time.Now().Sub(t)
		if succ == true {
			avg := totalSec.Seconds() * 1000 / float64(i)
			fmt.Println(avg)
			Info.Printf("Success on %vst try, Avg connection time: %f ms\n", i, avg)
			return
		}
	}
	Error.Printf("Failed to connect all %v times. Dail timeout: %v seconds\n", c.retry+1, c.timeout)
}

// checkOnce will try dial with timeout. return true if connected, otherwise false
func checkOnce(endpoint string, timeout int) bool {
	conn, err := net.DialTimeout("tcp", endpoint, time.Duration(timeout)*time.Second)
	if err != nil {
		Debug.Println("not connected")
		return false
	}
	conn.Close()

	return true
}
