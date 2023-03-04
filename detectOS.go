package main

import (
	"runtime"
)

func detectOS() string {
	//fmt.Printf("Running on %s\n", runtime.GOOS)

	//OS = runtime.GOOS

	return runtime.GOOS
}
