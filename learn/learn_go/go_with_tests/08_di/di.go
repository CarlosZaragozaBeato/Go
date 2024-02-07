package main 

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

// Greet send a personalized greetin to writer 
func Greet(writer io.writer, name string) {
  fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
  Greet(w, "world")
}

func main () {
  log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}
