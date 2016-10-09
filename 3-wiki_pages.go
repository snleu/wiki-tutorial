package main 

import (
	"io/ioutil"
	"net/http"
	"html/template"
)

// parses all templates
// template.Must panics when passed a non-nil error value
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

// creates a page with title and body text
type Page struct {
    Title string
    Body  []byte
}

// saves a new page as a txt file
func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

// loads the page from its file with name title and throws an error if no page is found
func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

// this function calls the correct html template for the handlers
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// calls the title and body, and formats it in html to view the page
func viewHandler(w http.ResponseWriter, r *http.Request) {
	// the title of the page is from the url path excluding "/view/"
    title := r.URL.Path[len("/view/"):]
    p, err := loadPage(title)
    // if the page called does not exist, there is a redirect to the /edit page
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    // title and body are formatted in html
    renderTemplate(w, "view", p)
}

// edits the body of a page, formatted in html
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    // error != nil if the page does not exist
    if err != nil {
    	// if the page does not exist then it shows an edit page with the entered title and blank text
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}

// saves edits made to a page and redirects to /view/ page
func saveHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/save/"):]
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    // errors that occur during p.save() will be reported to user
    err := p.save()
    if err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    	return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
    http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}