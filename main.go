package main

import (
	"fmt"

	"github.com/minjoo0729/pencoin/rest"
)


func main() {
	fmt.Println("Starting the server")
	rest.Start(4000)
}
