package handler

import (
	// standard library
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	// external packages
	"github.com/gorilla/mux"

	// internal packages
	"github.com/FraBle/WikidataQuiz/arduino"
	"github.com/FraBle/WikidataQuiz/question"
	"github.com/FraBle/WikidataQuiz/utility"
)

// The HomeHandler returns the index.html.
func HomeHandler(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadFile("../src/github.com/FraBle/WikidataQuiz/static/html/index.html")
	if err != nil {
		log.Printf("Error reading index.html: %v", err)
	}
	rw.Write(body)
}

// The QuestionHandler returns a randomly chosen question as JSON.
func QuestionHandler(rw http.ResponseWriter, req *http.Request) {
	var (
		response []byte
		err      error
	)
	switch utility.Random(0, 3) {
	case 0:
		q, err := question.CapitalQuestion()
		if err != nil {
			log.Printf("Error creating capital question", err)
		} else {
			response, err = json.Marshal(&q)
		}
	case 1:
		q, err := question.NobelPrizeWinnersQuestion()
		if err != nil {
			log.Printf("Error creating nobel prize winner question", err)
		} else {
			response, err = json.Marshal(&q)
		}
	case 2:
		q, err := question.WorldCupQuestion()
		if err != nil {
			log.Printf("Error creating world cup question", err)

			// use capital question as fall-back if dbpedia errored (time-out...)
			q, err = question.CapitalQuestion()
			if err != nil {
				log.Printf("Error creating capital question", err)
			} else {
				response, err = json.Marshal(&q)
			}
		} else {
			response, err = json.Marshal(&q)
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
