package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	//var cpu int = runtime.NumCPU()
	fmt.Fprintf(w, "Hello, World! Here's some info about your system: \n \n") // Format what's specificed: ResponseWriter and Request
	fmt.Fprintf(w, "CPU:", runtime.NumCPU())
}

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":1911", nil) // ListenAndServe needs server functoinality in the 2nd parameter. Not needed here, so nill is used
}
