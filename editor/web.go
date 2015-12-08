package main

import (
	"crypto/sha1"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

const dfltYAML = `
models:
  galaxie:
    nom:
      type: text
    position:
      type: point
    luminosité:
      type: number
      quantity: 2 or 4
      constraints:
        - float
        - in:
            - 0 to 1
            - 5 to 6
      erreur:
        type: number
        constraints:
            - float
            - in: 0 to 1
        quantity: 0 to 1
    images:
      type: image
      quantity: 0 to 5
  image:
    fichier:
      type: file
    bande:
      type: text
      constraints:
        - in:
            - u
            - g
            - r
            - i
            - z
  amas:
    nom:
      type: text
    nombre_de_galaxies:
      type: number
      constraints:
        - positive
        - integer
      quantity: 0 or 1
    galaxies:
        type: galaxie
        quantity: 0 to n
        parent: 1
`

func randomPlaygroundPath() string {
	b := sha1.Sum([]byte(fmt.Sprintf("%d", time.Now().Unix())))
	// Path is 3 bytes / 6 chars length
	return strings.ToUpper(fmt.Sprintf("%x", b[len(b)-3:]))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Moved temporarily
	http.Redirect(w, r, randomPlaygroundPath(), 302)
}

func playgroundHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			if e := server.templates.ExecuteTemplate(w, "error", err); e != nil {
				log.Println("[recover] an error occured: " + e.Error())
				fmt.Fprintf(w, e.Error())
			}
		}
	}()

	content := dfltYAML
	s, err := findSketchByName(r.URL.Path[1:])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if s != nil {
		content = s.content
	}

	// Records the content and render it only if no error occured
	wtmp := httptest.NewRecorder()
	type data struct {
		Context       context
		EditorContent string
	}
	args := &data{server.GetContext(), content}
	if err := server.templates.ExecuteTemplate(wtmp, "home", args); err != nil {
		panic(err)
	}
	fmt.Fprintf(w, wtmp.Body.String())
}
