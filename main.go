package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	return os.WriteFile(p.Title+".txt", p.Body, 0600)
}

func loadPage(title string) *Page {
	body, _ := os.ReadFile(title + ".txt")
	fmt.Printf("body: %v\n", body)

	return &Page{Title: title, Body: body}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	page := loadPage(title)
	fmt.Println(page)
	fmt.Fprintf(w, "<h1>%s</h1> <p>%s</p>", page.Title, page.Body)

}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: fileName, Body: []byte(body)}
	p.save()

	fmt.Fprintf(w, "File saved!")
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p := loadPage(title)
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>", p.Title, p.Title, p.Body)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, i Love %s", r.URL.Path[1:])
}

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/edit/", editHandler)
	fmt.Println("Server starting in PORT: 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
