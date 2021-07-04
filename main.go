// Copyright 2021 Mark Mandriota. All right reserved.
// Use of this source code is governed by a Apache 2.0-style
// license that can be found in the LICENSE file.

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
