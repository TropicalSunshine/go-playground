package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

var (
	count    = flag.Int("c", 4, "stop after count packets")
	interval = flag.Duration("i", 1*time.Second, "interval between packets")
	size     = flag.Int("s", 56, "packet size in bytes")
	verbose  = flag.Bool("v", false, "verbose output")
)

func main() {
	flag.Parse()
	fmt.Println(count, interval, size, verbose)
	fmt.Println(net.ResolveIPAddr("ip", "www.google.com"))
}
