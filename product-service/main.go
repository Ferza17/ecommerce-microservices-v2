package main

import (
	"github.com/ferza17/ecommerce-microservices-v2/product-service/cmd"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Run()
}
