package config

import (
  "github.com/gorilla/mux"
  "../app/controllers"
)

func Route(r *mux.Router) {
  //
  // Welcome
  //
	r.HandleFunc("/", controllers.WelcomeShow)
  
  //
  // Users
  //
	r.HandleFunc("/users", controllers.UsersIndex)
	r.HandleFunc("/users/new", controllers.UsersNew)  
	r.HandleFunc("/users/{name}", controllers.UsersShow)    
}