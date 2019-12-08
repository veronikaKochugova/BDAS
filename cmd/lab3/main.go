package main

import (

	// "io"

	"log"
	"net/http"
)

func helloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
	// fmt.Fprintf(w, "This is an example server.\n")
	//io.WriteString(w, "This is an example server.\n")
	//fmt.Fprintf(w, "Hi there, I love %s!", req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/hello", helloServer)
	err := http.ListenAndServeTLS(":443", "examples/server.crt", "examples/server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
