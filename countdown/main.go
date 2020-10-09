package main

import (
	"fmt"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/rishi-org-stack/count/schema"

	"html/template"
	"net/http"
)

var u schema.User
var e schema.Event

func home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("home.html")
	t.Execute(w, "home.html")
	// fmt.Fprintf(w,"ok")
}

func login(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("login.html")
	t.Execute(w, "login.html")
}
func logined(w http.ResponseWriter, r *http.Request) {

	u.Name = r.FormValue("name")
	u.Email = r.FormValue("email")
	u.Password = r.FormValue("password")
	data := u.Get()
	if data != nil {
		u.ID = data["_id"].(primitive.ObjectID)
		if data["event"] != nil {
			t, err := template.ParseFiles("uhome.html")
			if err != nil {
				log.Fatal(err)
			}
			//interface is oftype primitive.A
			Events := data["event"].(primitive.A) //(primitive.A)
			fmt.Println(Events)
			t.Execute(w, Events)
		} else {
			t, _ := template.ParseFiles("un.html")
			t.Execute(w, "un.html")
		}

	} else {
		// fmt.Fprintf(w,"you are not present in our database")
		t, _ := template.ParseFiles("signin.html")
		t.Execute(w, "signin.html")
	}

}

func add(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("add.html")
	t.Execute(w, "ok")
}

func added(w http.ResponseWriter, r *http.Request) {
	e.Name = r.FormValue("name")
	e.Date, _ = strconv.Atoi(r.FormValue("date"))
	u.Addevent(e)
	schema.Update(u)
}
func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logined", logined)
	// http.HandleFunc("/signin",signin)
	http.HandleFunc("/add", add)
	http.HandleFunc("/added", added)
	http.ListenAndServe(":8000", nil)
}
