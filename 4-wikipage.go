package main 

import (
    _ "github.com/lib/pq"
    "database/sql"
    "net/http"
    "html/template"
    "regexp"
    "log"
    "strings"
)

// parses all templates
// template.Must panics when passed a non-nil error value
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

// regexp.MustCompile will parse and compile the regular expression
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// creates a page with title and body text
type Page struct {
    Title string
    Body  string
}

// create a db and a table:
// sudo su _postgres -c \ "psql -c \"CREATE ROLE mydatabase LOGIN password 'averysecurepassword'\";"
// psql -d Pages -a -f 4-wikipage.sql
var db *sql.DB

func init() {
    var err error
    db, err = sql.Open("postgres", "user=snleu dbname=Pages sslmode=disable")
    if err != nil {
        log.Fatal(err)
  }
}

// this function calls the correct html template for the handlers
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}


// wrapper that takes view/edit/save handler function and returns func http.HandlerFunc
func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w,r)
            return
        }
        fn(w, r, m[2])
    }
}

func selectRow(title string, w http.ResponseWriter, r *http.Request) (*Page) {
    row := db.QueryRow("SELECT * FROM pages WHERE title = $1", title)
    p := new(Page)
    err := row.Scan(&p.Title, &p.Body)
    p.Title = strings.Trim(p.Title, " ")
    if err == sql.ErrNoRows {
        _, err1 := db.Exec("INSERT INTO pages VALUES($1, $2)", title, "")
        if err1 != nil {
            http.Error(w, err1.Error(), http.StatusInternalServerError)
        }
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
    } else if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    return p
}

// calls the title and body, and formats it in html to view the page
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
    p := selectRow(title, w, r)
    renderTemplate(w, "view", p)
}

// edits the body of a page, formatted in html
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
    p := selectRow(title, w, r)
    renderTemplate(w, "edit", p)
}

// saves edits made to a page and redirects to /view/ page
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
    body := r.FormValue("body")
    p := &Page{Title: title, Body: body}
    _, err := db.Exec("UPDATE pages SET body = $1 WHERE title = $2", p.Body, p.Title)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
    http.HandleFunc("/view/", makeHandler(viewHandler))
    http.HandleFunc("/edit/", makeHandler(editHandler))
    http.HandleFunc("/save/", makeHandler(saveHandler))
    http.ListenAndServe(":8080", nil)
}