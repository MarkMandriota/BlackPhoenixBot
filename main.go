package main

import (
	"flag"
	"runtime"
)

var (
	conf = flag.String("conf", "config.mag", "Pass to magic parser config file.")
)

func init() {
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	_ = make([]byte, s.HeapSys/3)

	flag.Parse()
}

func main() {}
