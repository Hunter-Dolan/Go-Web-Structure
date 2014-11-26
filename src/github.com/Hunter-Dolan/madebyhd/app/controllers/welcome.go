package controllers

import (
  "net/http"
  "fmt"
  
  _"github.com/gorilla/mux"
  _"../models"
  _"../../config/globals"
)

func WelcomeShow(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "test")
}