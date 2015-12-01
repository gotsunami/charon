package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

func playgroundHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			if e := server.templates.ExecuteTemplate(w, "error", err); e != nil {
				log.Println("[recover] an error occured: " + e.Error())
				fmt.Fprintf(w, e.Error())
			}
		}
	}()

	// Records the content and render it only if no error occured
	wtmp := httptest.NewRecorder()
	type data struct {
		Context context
	}
	args := &data{server.GetContext()}
	if err := server.templates.ExecuteTemplate(wtmp, "home", args); err != nil {
		panic(err)
	}
	fmt.Fprintf(w, wtmp.Body.String())
}
