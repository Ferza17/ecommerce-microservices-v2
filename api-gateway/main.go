package main

import (
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/cmd"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Run()
}
