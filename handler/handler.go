package handler

import (
	"encoding/json"
	"github.com/FraBle/WikidataQuiz/arduino"
	"github.com/FraBle/WikidataQuiz/question"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func HomeHandler(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadFile("../src/github.com/FraBle/WikidataQuiz/static/html/index.html")
	if err != nil {
		log.Printf("%v", err)
	}
	rw.Write(body)
}

func QuestionHandler(rw http.ResponseWriter, req *http.Request) {
	rand.Seed(time.Now().UnixNano())
	newNumber := rand.Intn(3)

	var response []byte
	var err error
	switch newNumber {
	case 0:
		response, err = json.Marshal(question.CapitalQuestion())
	case 1:
		response, err = json.Marshal(question.NobelPrzWinnersDiedAfter2000Question())
	case 2:
		response, err = json.Marshal(question.WorldCupQuestion())
	}
	if err != nil {
		log.Printf("%v", err)
	}
	rw.Write(response)
}

func ColorHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	color := vars["color"]
	switch color {
	case "black":
		if err := arduino.SetColor("<02>"); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	case "red":
		if err := arduino.SetColor("<03>"); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	case "green":
		if err := arduino.SetColor("<04>"); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	case "blue":
		if err := arduino.SetColor("<05>"); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	}
}
