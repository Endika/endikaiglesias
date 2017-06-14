// todos.go
package main

import (
	"html/template"
	"net/http"
)

func main() {
	tmpl := template.Must(template.ParseFiles("/home/endikaiglesias/index.html"))
	data := struct {
		Name     string
		Email    string
		Linkedin string
		Github   string
		Twitter  string
		Blog     string
	}{
		Name:     "Endika Iglesias",
		Email:    "me@endikaiglesias.com",
		Linkedin: "https://www.linkedin.com/in/endika-iglesias-0420a731/",
		Github:   "https://github.com/Endika",
		Twitter:  "https://twitter.com/Endika_Iglesias",
		Blog:     "http://optimizarsinmas.blogspot.com.es/",
	}

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("/home/endikaiglesias/static"))))

	http.Handle("/img/",
		http.StripPrefix("/img/",
			http.FileServer(http.Dir("/home/endikaiglesias/img"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8002", nil)
}
