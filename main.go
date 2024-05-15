package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)


func main(){
  http.HandleFunc("/", handle_request)
  fmt.Println("Server started and serving at: 8090")
  http.ListenAndServe(":8090", nil)
}

func handle_request(w http.ResponseWriter, r *http.Request){
  log.Println(r.URL.Path)
  cleaned_path := filepath.Clean(r.URL.Path)
  method := strings.ToLower(r.Method)
   
  handler_file := fmt.Sprintf(".%s/%s.go", cleaned_path, method)

  if _, err := os.Stat(handler_file); os.IsNotExist(err){
    log.Println(handler_file)
    http.NotFound(w, r)
    return
  }


  fmt.Fprintf(w, "handler found: %s", handler_file)
}
