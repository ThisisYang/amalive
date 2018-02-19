package main

import (
	"flag"
	"io/ioutil"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var ip = flag.String("ip", "127.0.0.1", "IP to connect")
var port = flag.Int("port", 22, "port to check")
var cinterval = flag.Int("cintv", 5, "check interval in seconds between tcp check")
var rinterval = flag.Int("rintv", 2, "retry interval in seconds between each retry. if retry is not set, interval will be ignored")
var retry = flag.Int("retry", 3, "times to retry")
var debug = flag.Bool("debug", false, "debug mode")
var timeout = flag.Int("timeout", 1, "I/O timeout in seconds for both each connection and read for the first byte")

func main() {
	flag.Parse()
	ci := NewCheckInfo(*ip, *port, *timeout, *rinterval, *retry)

	if *debug {
		setUpLogger(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
		Debug.Println("Debug mode")
	} else {
		setUpLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	}
	Debug.Printf("check ednpiont: %v, retry: %v times, Interval between each retry; %vs, Timeout: %vs", ci.endpoint, ci.retry, ci.retryIntv, ci.timeout)

	done := make(chan struct{})
	var wg sync.WaitGroup
	go waitForSignal(done)

	//resp, in millisecond
	wg.Add(1)
	go func() {
		t := time.NewTicker(time.Duration(*cinterval) * time.Second)
		defer t.Stop()
		defer wg.Done()
		for {
			select {
			case <-done:
				Info.Println("received termination signal, stopping all goroutines")
				return
			case <-t.C:
				ci.tcpCheck()
			}
		}
	}()

	wg.Wait()

}

// waitForSignal will close done channel when received signal
func waitForSignal(done chan struct{}) {
	sigs := make(chan os.Signal, 1)
	// register sigs channel to receive SIGINT and SIGTERM signal
	signal.Notify(sigs, os.Interrupt, os.Kill, syscall.SIGTERM) //syscall.SIGINT, syscall.SIGTERM
	<-sigs
	close(done)
}
