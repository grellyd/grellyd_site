package handlers

import (
	"fmt"
	"net/http"
	"regexp"
	"grellyddotcom/pages"
	"grellyddotcom/templates"
)

var validPath = regexp.MustCompile("^/(index|edit|save|view|status)/([a-zA-Z0-9]*)$")

func MakeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func StatusHandler(w http.ResponseWriter, r *http.Request, title string) {
	fmt.Fprintf(w, "OK")
}

func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &pages.Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := pages.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

// will become edit blog post
func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := pages.LoadPage(title)
    if err != nil {
        p = &pages.Page{Title: title}
    }
	renderTemplate(w, "edit", p)
}

func BlogHandler(w http.ResponseWriter, r *http.Request, title string) {
}

func IndexHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := pages.LoadPage(title)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	renderTemplate(w, "index", p)
}

func renderTemplate(w http.ResponseWriter, tmplt string, p *pages.Page) {
	err := templates.All.ExecuteTemplate(w, tmplt+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
