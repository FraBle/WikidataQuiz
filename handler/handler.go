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
)

// The HomeHandler returns the index.html
func HomeHandler(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadFile("../src/github.com/FraBle/WikidataQuiz/static/html/index.html")
	if err != nil {
		log.Printf("%v", err)
	}
	rw.Write(body)
}

// The QuestionHandler returns a randomly chosen question as json.
func QuestionHandler(rw http.ResponseWriter, req *http.Request) {
	var (
		response []byte
		err      error
	)
	switch rand.Intn(3) {
	case 0:
		response, err = json.Marshal(question.CapitalQuestion())
	case 1:
		response, err = json.Marshal(question.NobelPrizeWinnersQuestion())
	case 2:
		q, err := question.WorldCupQuestion()
		if err != nil {
			log.Printf("Error creating world cup question", err)
			// use capital question as fall-back if dbpedia errored (time-out...)
			response, err = json.Marshal(question.CapitalQuestion())
		} else {
			response, err = json.Marshal(q)
		}
	}
	if err != nil {
		log.Printf("Error creating question: %v", err)
	}
	rw.Write(response)
}

// The ColorHandler sets the LED color of the LED stripe connected to the Arduino.
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
