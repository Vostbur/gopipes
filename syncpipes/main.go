package main

import (
	"fmt"
	"time"
)

var hosts = [...]string{
	"192.168.1.1",
	"192.168.1.2",
	"192.168.1.3",
	"192.168.1.4",
	"192.168.1.5",
	"192.168.1.6",
	"192.168.1.7",
	"192.168.1.8",
	"192.168.1.9",
	"192.168.1.10",
	"192.168.1.11",
	"192.168.1.12",
}

var commands = []string{
	"sh run",
	"sh int",
	"sh ip int br",
	"sh clock",
	"sh ip route",
}

func login(host string) {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("result " + host)
}

func exec(conn, cmd string) {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(conn + " result " + cmd)
}

func main() {
	start := time.Now()

	for _, host := range hosts {
		login(host)
		for _, cmd := range commands {
			exec(host, cmd)
		}
	}

	duration := time.Since(start)
	fmt.Printf("Duration: %d ms\n", duration.Milliseconds())
}
