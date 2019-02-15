package main

import (
	"fmt"
	"net/http"
	"os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintln(w, "hello , this is "+hostname)
}

func main() {
	http.HandleFunc("/", IndexHandler)
	err := http.ListenAndServe("0.0.0.0:9999", nil)
	fmt.Printf("err:%v\n", err)
}
