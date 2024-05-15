package main

import (
	"fmt"
	"net/http"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request)  {
  fmt.Fprintf(w, "Welcome to the user list")
}
