package main

import (
	"fmt"
	"sync"
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

func login(host string) string {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("result " + host)
	return "result " + host
}

func exec(conn, cmd string) string {
	time.Sleep(1000 * time.Millisecond)
	return conn + " result " + cmd
}

func main() {
	start := time.Now()

	connections := make(chan string, len(hosts))
	wg := &sync.WaitGroup{}
	for _, host := range hosts {
		wg.Add(1)
		go func(c chan string, h string, wg *sync.WaitGroup) {
			defer wg.Done()
			connections <- login(h)
		}(connections, host, wg)
	}
	wg.Wait()
	close(connections)

	i := 0
	output := make([]chan string, len(hosts))
	for i := range output {
		output[i] = make(chan string, len(commands))
	}
	for conn := range connections {
		go func(o chan string, conn string, comms []string) {
			for _, cmd := range comms {
				o <- exec(conn, cmd)
			}
			close(o)
		}(output[i], conn, commands)
		i++
	}

	for _, perHost := range output {
		for out := range perHost {
			fmt.Println(out)
		}
	}

	duration := time.Since(start)
	fmt.Printf("Duration: %d ms\n", duration.Milliseconds())
}
