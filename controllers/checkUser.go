package controllers

import (
	"net/http"

	"github.com/jashfeer/RESTfulAPI/initialization"
)

func AlreadyLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	session, _ := initialization.Store.Get(r, "session")
	_, ok := session.Values["name"]
	return ok
}