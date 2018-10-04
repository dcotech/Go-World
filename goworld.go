package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!") // Format what's specificed: ResponseWriter and Request
}

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":1911", nil) // ListenAndServe needs server functoinality in the 2nd parameter. Not needed here, so nill is used

}
