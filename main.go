// Server
package main

import (
	"module1/modules"
)

func main() {

	go modules.InitServer()
	go modules.InitClient()

}
