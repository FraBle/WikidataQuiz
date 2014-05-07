package main

import (
	// standard Library
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	// external packages
	"bitbucket.org/kardianos/osext"
	"github.com/gorilla/mux"

	// internal packages
	"github.com/FraBle/WikidataQuiz/config"
	"github.com/FraBle/WikidataQuiz/handler"
)

// The main function reads the config and starts the web server.
func main() {

	if err := chdirToBinary(); err != nil {
		log.Printf("Error changing directory: %v", err)
	}

	initializeLogger()

	if err := config.ReadConfig(); err != nil {
		log.Printf("Error reading config file: %v", err)
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handler.HomeHandler).Methods("GET")
	router.HandleFunc("/question", handler.QuestionHandler).Methods("GET")
	router.HandleFunc("/led/{color}", handler.ColorHandler).Methods("GET")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../src/github.com/FraBle/WikidataQuiz/static"))))
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// chdirToBinary() change the current working directory to the binary file so that relative paths work without any problems.
func chdirToBinary() error {
	path, err := osext.ExecutableFolder()
	if err != nil {
		return err
	}
	return os.Chdir(path)
}

// initializeLogger() enables dual logging: into a log file and to standard out.
func initializeLogger() {
	if err := os.MkdirAll("../log", os.ModeDir|os.ModePerm); err != nil {
		log.Fatalf("Error creating log directory: %v", err)
	}
	logfile, err := os.OpenFile("../log/wikidatachat.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
}

func initializeRandomNumberGenerator() {
	rand.Seed(time.Now().UTC().UnixNano())
}
