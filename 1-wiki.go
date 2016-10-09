package main

import (
	"fmt"
	"io/ioutil"
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

// creates a test page, saves it, loads it, and prints the body
func main() {
    p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page.")}
    p1.save()
    p2, _ := loadPage("TestPage")
    fmt.Println(string(p2.Body))
}
