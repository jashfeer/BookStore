package initialization

import (
	"fmt"

	"gorm.io/gorm"
	"github.com/gorilla/sessions"

	"github.com/jashfeer/RESTfulAPI/database"
)

var Db *gorm.DB
var Store= sessions.NewCookieStore([]byte("super-secret"))


func Init() {
	var err error
	Db, err = database.OpenDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("you connected to your database.")
	
}
