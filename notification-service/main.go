package main

import (
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/cmd"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Run()
}

//TODO:
// - add migration migration
// - add table email template
