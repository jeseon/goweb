package main

import (
	"fmt"
	"net/http"
	"html/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

// Person represents person struct
type Person struct {
	Name  string
	Phone string
}

// GetFormDataHandler is html form
func GetFormDataHandler(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(res, req)
	// res.Write([]byte(fmt.Sprintf(html)))
}

// ReadFormDataHandler is action that process form-data
func ReadFormDataHandler(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	if err != nil {
		fmt.Println(err)
	}

	person := new(Person)
	decoder := schema.NewDecoder()
	err = decoder.Decode(person, req.PostForm)

	if err != nil {
		fmt.Println(err)
	}

	res.Write([]byte(fmt.Sprintf("Name is %v \n", person.Name)))
	res.Write([]byte(fmt.Sprintf("Phone is %v \n", person.Phone)))
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", GetFormDataHandler)
	router.HandleFunc("/index", GetFormDataHandler)
	router.HandleFunc("/process_form_data", ReadFormDataHandler)

	http.ListenAndServe(":8080", router)
}
