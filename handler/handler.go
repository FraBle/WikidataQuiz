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
		arduino.SetColor("<02>")
	case "red":
		arduino.SetColor("<03>")
	case "green":
		arduino.SetColor("<04>")
	case "blue":
		arduino.SetColor("<05>")
	}
}
