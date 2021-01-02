// Tutorial found here: https://golang.org/doc/articles/wiki/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
    Title string
    Body  []byte
}
// []byte means "a byte slice"
// []byte rather than string because that is the type expected by the io libraries we will use

// SAVE method
// "This is a method named save that takes as its receiver p, a pointer to Page . It takes no parameters, and returns a value of type error."
// save method returns an error value because that is the return type of WriteFile 
//  If all goes well, Page.save() will return nil (the zero-value for pointers, interfaces, and some other types)
//The octal integer literal 0600, passed as the third parameter to WriteFile, indicates that the file should be created with read-write permissions for the current user only. 
func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

// LOADPAGE function
func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

// func main() {
//     p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
//     p1.save()
//     p2, _ := loadPage("TestPage")
//     fmt.Println(string(p2.Body))
// }

// Using net/http package create a handler that will allow users to view a wiki page. It will handle URLs prefixed with "/view/".
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
    http.HandleFunc("/view/", viewHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
