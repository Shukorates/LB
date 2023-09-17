package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var port = os.Args[2]
var serverName = os.Args[1]

func main() {
	http.HandleFunc("/", forwardRequest)
	log.Fatal(http.ListenAndServe(port, nil))

}
func forwardRequest(res http.ResponseWriter, req *http.Request) {

	fmt.Fprint(res, serverName)
}
