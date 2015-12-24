package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
)

const (
	DEFAULT_PORT = "8080"
)

func perror(msg string) {
	fmt.Fprintf(os.Stderr, "error: %s\n", msg)
	os.Exit(1)
}

type context map[string]interface{}

type webserver struct {
	router    *mux.Router
	templates *template.Template
	context   context
}

func (s webserver) GetContext() context {
	return s.context
}

var server *webserver

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("Usage: %s [options] configfile\n", os.Args[0]))
		fmt.Fprintf(os.Stderr, `Where configfile holds JSON data like:
{
    "dbURI": "mongodb://uri",
    "staticURI": "/static",
    "templatePath": "/path/to/templates/",
}

Options:
`)
		flag.PrintDefaults()
		os.Exit(2)
	}

	port := flag.String("p", DEFAULT_PORT, "http listening port")
	version := flag.Bool("version", false, "program version")

	flag.Parse()

	if *version {
		fmt.Printf("version: %s\n", appVersion)
		return
	}

	if len(flag.Args()) != 1 {
		fmt.Fprint(os.Stderr, "Config file is missing.\n\n")
		flag.Usage()
	}

	var err error
	conf, err := parseConfig(flag.Arg(0))
	if err != nil {
		perror(fmt.Sprintf("config parsing error: %s\n", err.Error()))
	}

	if err := dial(conf.DbURI, conf); err != nil {
		log.Println("can't connect to database: ", err)
	}
	log.Println("connected to database.")

	router := mux.NewRouter().StrictSlash(true)

	// Routes
	router.HandleFunc("/save/{id:[0-9A-Z]{6}}", saveSketchHandler).Methods("POST").Name("saveSketch")
	router.HandleFunc("/{id:[0-9A-Z]{6}}", playgroundHandler).Name("playground")
	router.HandleFunc("/example/{yamlFile:[A-Za-z]+}", exampleHandler).Name("example")
	router.HandleFunc("/", homeHandler).Name("home")
	router.PathPrefix(conf.StaticURI).Handler(http.StripPrefix(conf.StaticURI, http.FileServer(http.Dir("assets/"))))
	http.Handle("/", router)

	tmpl := template.Must(template.New("playground").
		ParseFiles(
		path.Join(conf.TemplatePath, "error.html"),
		path.Join(conf.TemplatePath, "example.html"),
		path.Join(conf.TemplatePath, "home.html")))

	server = &webserver{router, tmpl, context{"static": conf.StaticURI, "version": appVersion}}

	log.Printf("Listening on port %v, waiting for incoming requests...", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
