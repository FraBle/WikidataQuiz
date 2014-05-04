package handler

import (
	"encoding/json"
	"github.com/FraBle/WikidataQuiz/arduino"
	"github.com/FraBle/WikidataQuiz/question"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func HomeHandler(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadFile("../src/github.com/FraBle/WikidataQuiz/static/html/index.html")
	if err != nil {
		log.Printf("%v", err)
	}
	rw.Write(body)
}

func QuestionHandler(rw http.ResponseWriter, req *http.Request) {
	question, err := json.Marshal(question.CapitalQuestion())
	if err != nil {
		log.Printf("%v", err)
	}
	rw.Write(question)
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
