package main

import (
	"bitbucket.org/kardianos/osext"
	"github.com/FraBle/WikidataQuiz/handler"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	if err := chdirToBinary(); err != nil {
		log.Printf("Error changing directory: %v", err)
	}

	initializeLogger()

	// if err := config.ReadConfig(); err != nil {
	//     log.Printf("Error reading config file: %v", err)
	// }

	router := mux.NewRouter().StrictSlash(true)

	// router.NotFoundHandler = http.HandlerFunc(handler.NotFound)
	router.HandleFunc("/", handler.HomeHandler).Methods("GET")
	router.HandleFunc("/question", handler.QuestionHandler).Methods("GET")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../src/github.com/FraBle/WikidataQuiz/static"))))
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func chdirToBinary() error {
	path, err := osext.ExecutableFolder()
	if err != nil {
		return err
	}
	return os.Chdir(path)
}

func initializeLogger() {
	if err := os.MkdirAll("../log", os.ModeDir); err != nil {
		log.Fatalf("Error creating log directory: %v", err)
	}
	logfile, err := os.OpenFile("../log/wikidatachat.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
}
