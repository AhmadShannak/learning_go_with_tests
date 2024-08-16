package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const Greeting = "Hello, "

func Greet(writer io.Writer, format, name string) {
	fmt.Fprintf(writer, Greeting+format, name)
}

func main() {
	Greet(os.Stdout, "%s", "Saber")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "%s", "world")
}
