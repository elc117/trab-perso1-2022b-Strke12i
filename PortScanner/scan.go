package PortScanner

import (
	"net"
	"strconv"
	"time"
)

func OpenPorts(protocol string, hostname string, p1 int, p2 int) []int {
    var openPorts []int
    rng := p2 - p1
    ports := make(chan int, rng)
    result := make(chan int)

    for i := 0; i < cap(ports); i++{
        go PortScanner(protocol, hostname, ports, result)
    }

    go func() {
        for i := p1; i <= p2; i++{
            ports <- i
        }
    }()

    for i := p1; i <= p2; i++ {
        port := <- result
        if port != 0 {
            openPorts = append(openPorts,port)
        }
    }

    return openPorts


}


func PortScanner(protocol string, hostname string, ports chan int, result chan int) {
    for i := range ports {
        conn, err := net.DialTimeout(protocol, hostname + ":" + strconv.Itoa(i), time.Second)
        if err != nil {
            result <- 0
            continue
        }

        result <- i
        conn.Close()
    }
}


