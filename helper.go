package amalive

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

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

//use Value to dynamicly store value in flag
//https://golang.org/pkg/flag/#Value
//type Value interface {
//    String()
//    Set(string) error
//}

//converts struct as a string
func (endpoint *TCPflag) String() string {
	return fmt.Sprint(endpoint.Endpoint)
}

// Set called by flag.Parse function. Initializes the flag
func (endpoint *TCPflag) Set(value string) error {
    //fmt.Println(len(endpoint.Endpoint))
    for _, item := range strings.Split(value, ",") {
        arr := strings.Split(item, ":")
        if port, err := strconv.Atoi(arr[1]): err != nill{
            return err
        }
        instance := TCP{IP: arr[0], PORT: port}
        //instance := TCP{IP: arr[0], PORT: strconv.Atoi(arr[1])}
        endpoint.Endpoint = append(endpoint.Endpoint, instance)
    }
    return nil
}

var try TCPflag
func init(){
    flag.Var(&try, "tcp", "try tcp, comma seperate")
}

func main(){
    flag.Parse()
    c := try.Endpoint[1].PORT
    fmt.Println(c)
}


type UDP struct {
	IP   string
	Port string
}

type PING struct {
	IP string
}


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
	flag.IntVar(&commandOpts.Port, name, value, usage)
	flag.IntVar(&commandOpts.Protocol, name, value, usage)
}
