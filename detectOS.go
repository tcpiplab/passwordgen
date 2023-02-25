package main

import (
	"fmt"
	"runtime"
)

func detectOS() string {
	fmt.Printf("Running on %s\n", runtime.GOOS)

	return runtime.GOOS
}
