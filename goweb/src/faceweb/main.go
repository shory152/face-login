package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/search", searchFace)

	log.Fatal(http.ListenAndServe(":9090", nil))
}

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func index(w http.ResponseWriter, req *http.Request) {
	buf, _ := ioutil.ReadFile("camera.html")
	w.Write(buf)
}
