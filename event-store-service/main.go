package main

import (
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/cmd"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Run()
}
