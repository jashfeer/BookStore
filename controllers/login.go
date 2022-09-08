package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jashfeer/RESTfulAPI/initialization"
)

type User struct{
	Username string
	Password string
}



func Login(w http.ResponseWriter,r *http.Request){
	if AlreadyLoggedIn(w,r) {
		fmt.Fprintf(w,"You are already LoggedIn ")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	if (user.Username == "user" && user.Password == "123"){

		session,_:=initialization.Store.Get(r,"session")
		session.Values["name"]=user.Username
		session.Save(r,w)

		fmt.Fprintf(w,"LogIn Successful ")
		return
	}

	fmt.Fprintf(w,"Incorrect Username or Password")
}