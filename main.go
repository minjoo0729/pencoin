package main

import (
	explorer "github.com/minjoo0729/pencoin/explorer/templates"
	"github.com/minjoo0729/pencoin/rest"
)


func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
