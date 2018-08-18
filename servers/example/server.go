package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("Server starting\n")

	http.HandleFunc("/", handler)
	no(http.ListenAndServe(":9000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling Request from " + r.RemoteAddr + " For address: " + r.RequestURI)

	io.WriteString(w, "Hi there ..\n")
	name, err := os.Hostname()
	no(err)
	io.WriteString(w, fmt.Sprint("Thanks for visiting server: "+name+"\n"))
}

func no(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
