package main 

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

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

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// the title of the page is from the url path excluding "/view/"
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    // title is formated as a header and body as documentation
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8080", nil)
}