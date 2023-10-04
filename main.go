package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("index.html"))

    tmpl.Execute(w, nil)
  })

  fmt.Println("Listening on :8000...")

  log.Fatal(
    http.ListenAndServe(":8000", nil),
  )
}
