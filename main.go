package main

import "fmt"
import "os"
import "net/http"

func handler(w http.ResponseWriter, r *http.Request){

h, err := os.Hostname()
	if err == nil {
  fmt.Fprint(w, h)
  }


}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
