package main

import (
	"fmt"
	"io/ioutil"
	// "math/rand"
	"net/http"
	// "strconv"
	// "time"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>Visiting: %s.txt</h1><div>Content: %s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8080", nil)
}

// func randGenerator(i int) string {
// 	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	return strconv.Itoa(seed.Intn(i))
// }
