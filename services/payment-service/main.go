package main

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/cmd"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Run()
}
