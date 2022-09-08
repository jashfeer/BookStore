package controllers

import (
	"fmt"
	"net/http"

	"github.com/jashfeer/RESTfulAPI/initialization"
)

func Logout(w http.ResponseWriter,r *http.Request){
	if !AlreadyLoggedIn(w,r) {
		fmt.Fprintf(w,"You are NOT a user ")
		return
	}
	session,_:=initialization.Store.Get(r,"session")
	delete(session.Values,"name")
	session.Save(r,w)

	fmt.Fprintf(w,"LogOut Successful ")
	
}