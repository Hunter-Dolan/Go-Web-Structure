package controllers

import (
  "net/http"
  "fmt"
  
  "github.com/gorilla/mux"
  "../models"
  "../../config/globals"
)

func UsersIndex(w http.ResponseWriter, r *http.Request) {
  headers := w.Header()
  headers.Set("Content-Type", "text/html; charset=UTF-8")
  w.WriteHeader(200)
  
  count := 0
  
  model := global.DB.Connection.Model(models.User{})
  
  users, _ := model.Select("name").Rows()
  
  defer users.Close()

  model.Count(&count)
  
  content := ""
  name := ""
  for users.Next() {
    users.Scan(&name)
    content += "<br>"
    content += "User: "
    content += "<a href='/users/"
    content += name
    content += "'>"
    content += name
    content += "</a><br>"
  }

	fmt.Fprintf(w, "There are %i users <form method='post' action='/users/new'><input name='name'><input type='submit'></form> let's talk about it: <br> %q", count, content)
}

func UsersShow(w http.ResponseWriter, r *http.Request) {
  headers := w.Header()
  headers.Set("Content-Type", "text/html; charset=UTF-8")
  w.WriteHeader(200)
  
  vars := mux.Vars(r)
  name := vars["name"]
  
  user := models.FindUserByName(name)
  
  fmt.Fprintf(w, "%q", user)
}

func UsersNew(w http.ResponseWriter, r *http.Request) {
  name := r.FormValue("name")
	
  user := models.User{}
  user.Name = name
  
  global.DB.Connection.Create(&user)

  headers := w.Header()
  headers.Set("Location", "/users")
  w.WriteHeader(302)
}
