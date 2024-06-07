package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings" 
	"unicode/utf8" 
	"todomysql/pkg/models"
)

// creating a list function for the struct applictaion
func (app *application) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allowed", "POST")
	if r.Method != "GET" { //checking for the method id GET
		http.Error(w, "Method Not Allowed", 500)
		return
	}

	files := []string{
		"./ui/html/list.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
		"./ui/html/header.partial.tmpl",
	}

	//  panic("oops! something went wrong")

	s, err := app.todos.GetMultiple()
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}
	err = ts.Execute(w, s)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

}

func (app *application) ListSingle(w http.ResponseWriter, r *http.Request) {
	//need to take the id from the user
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	s, err := app.todos.GetSingle(id)
	if err == models.ErrNoRecord {
		http.Error(w, "Internal Server Ereeeeror", 500)
		return
	} else if err != nil {
		http.Error(w, "Eroooooror", 500)
		return
	}
	fmt.Fprintf(w, "%v", s)

}

// creating a custom add function to struct application(object app)
// this application has variable todos(which is a struct that has a custom insert function and variable DB(this DB has Exec) )
func (app *application) Add(w http.ResponseWriter, r *http.Request) {

	title := r.FormValue("name")

	errors := make(map[string]string)
	if strings.TrimSpace(title) == "" {
		errors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(title) > 100 {
		errors["title"] = "This field is too long (maximum is 100 characters)"
	}
	if len(errors) > 0 {
		fmt.Fprint(w, errors)
		return
	}

	_, err := app.todos.Insert(title) //name will be taken using r.FormValue
	// if error below lines work else insert function will be done on sql but cannot be seen in html home page
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err)
		return
	}
	// redirecting to "/" will call the GetMultiple function in List function
	// and thus the added row will be visisble
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func (app *application) Delete(w http.ResponseWriter, r *http.Request) {
	del := r.URL.Query().Get("id")
	idToDelete, _ := strconv.Atoi(del)

	errDel := app.todos.Delete(idToDelete)
	if errDel != nil {
		log.Println(errDel)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) Update(w http.ResponseWriter, r *http.Request) {
	updId := r.URL.Query().Get("id")
	updateValue := r.URL.Query().Get("updateValue")
	idToUpdate, _ := strconv.Atoi(updId)
	log.Println(idToUpdate)

	errUpdate := app.todos.Update(idToUpdate, updateValue)
	if errUpdate != nil {
		log.Println(errUpdate)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) UpdateForm(w http.ResponseWriter, r *http.Request) {
	upid, _ := strconv.Atoi(r.FormValue("id"))
	upval := app.todos.Upadateform(upid, r.FormValue("message"))
	log.Println()
	if upval != nil {
		log.Println(upval)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
