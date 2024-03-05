package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/minjoo0729/pencoin/blockchain"
)

const port string = ":4000"

type homeData struct {
	UserName string
	Blocks   []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/pages/home.gohtml"))
	data := homeData{"Minjoo", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
