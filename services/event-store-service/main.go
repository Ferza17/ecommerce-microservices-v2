package main

import (
	"runtime"

	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Run()
}
