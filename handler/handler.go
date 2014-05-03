package handler

import (
	"encoding/json"
	"github.com/FraBle/WikidataQuiz/question"
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
