package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"os"
	//"io"
	"io/ioutil"
	//"time"
)


var endpoints TCPflag
var Interval = flag.Int("interval", 2, "interval in seconds between each retry. if retry is not set, interval will be ignored")
var Retry = flag.Int("retry", 3, "times to retry")
var debug = flag.Bool("debug", false, "debug mode")
var Timeout = flag.Int("timeout", 5, "I/O timeout in seconds for both each connection and read for the first byte")
// --tcp=ip:port
// --udp=ip:port
// --ping=ip

type TCPflag struct {
	Endpoint []TCP
}

type TCP struct {
	IP   string
	Port int
}

//converts struct as a string
func (endpoint *TCPflag) String() string {
	return fmt.Sprint(endpoint.Endpoint)
}

// Set called by flag.Parse function. Initializes the flag
func (endpoint *TCPflag) Set(value string) error {
	// --tcp ip:port,ip:port,ip2:port2
	for _, item := range strings.Split(value, ",") {
		arr := strings.Split(item, ":")
		ip := arr[0]
		port, err := strconv.Atoi(arr[1])
		if err != nil {
			return err
		}
		instance := TCP{IP: ip, Port: port}
		//instance := TCP{IP: arr[0], PORT: strconv.Atoi(arr[1])}
		endpoint.Endpoint = append(endpoint.Endpoint, instance)
	}
	return nil
}

func init() {
	flag.Var(&endpoints, "tcp", "list of tcp endpoint, for example 127.0.0.1:22. if multiple endpoint, use ',' to seperate")
}

func main() {
	flag.Parse()

	if *debug {
		SetUpLogger(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
		Debug.Println("Debug mode is turned on")
	} else {
		SetUpLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	}
	Debug.Printf("retry: %v times, Interval between each retry; %vs, Timeout: %vs", *Retry, *Interval, *Timeout)

	//resp, in millisecond
	var resp float64
	for _, endpoint := range endpoints.Endpoint {
		//fmt.Println(endpoint.IP, endpoint.Port)
		resp = TCPCheck(endpoint.IP, endpoint.Port)
	}
	fmt.Println(resp)

	//fmt.Println(endpoints.Endpoint[1].IP, endpoints.Endpoint[1].Port)
	//fmt.Println(*Interval, *Retry, *Timeout)
	//ip := endpoints.Endpoint[0].IP
	//c := endpoints.Endpoint[0].Port
	//resp := 0
	//count := 0
	//for (resp == 0) && (count < 3) {
	//	resp = TCPCheckOnce(ip, c)
	//	count += 1
	//}
}

/*
// CommandOptions are options client provided
type CommandOptions struct {
	Addr     string
	Port     int
	Protocol string
	Interval time.Duration
}

// InstallFlags are function initalize the flag value
func (commandOpts *CommandOptions) InstallFlags() {
	flag.StringVar(&commandOpts.Addr, "IP", "127.0.0.1", "remote IP address")
	flag.IntVar(&commandOpts.Port, "Port", 22, "remote port")
	flag.StringVar(&commandOpts.Protocol, "TCP", "tcp", "Protocol to use")
}
*/
