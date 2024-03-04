package main

import (
	"fmt"

	"github.com/minjoo0729/pencoin/person"
)

func main() {
	mj := person.Person{}
	mj.SetDetails("minjoo", 23)
	fmt.Println("Main : ", mj)
}

// Method :  {minjoo 23}
// Main :  { 0}