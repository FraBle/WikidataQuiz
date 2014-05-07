package question

import (
	"encoding/json"
	"fmt"
	"github.com/FraBle/WikidataQuiz/model"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
)

type peopleDeathDateResponse struct {
	Status struct {
		Error        string
		Items        int
		Querytime    string
		Parsed_query string
	}
	Items []int
	Props struct {
		P570 []interface{} `json:"570"`
	}
}

// TODO: start date and end data saved in variable/ constant

// NobelPrizeWinnersQuestion() generates a question about the death date of a nobel prize winner who died after 2000.
func NobelPrizeWinnersQuestion() *model.Question {
	result := new(model.Question)

	deathYears, peopleIDs := getPeopleIDsAndDeathYears()

	selectedPersonIndex := rand.Intn(len(peopleIDs))

	result.RightAnswer = rand.Intn(4)
	var answers [4]string

	answers[result.RightAnswer] = strconv.Itoa(deathYears[selectedPersonIndex])

	offsets := fourRandomNumbersIn(15) // one number not used

	for i := 0; i < 3; i++ {
		newAnswer := 2000 + offsets[i]
		if newAnswer == deathYears[selectedPersonIndex] {
			newAnswer = 1999
		}

		if i >= result.RightAnswer {
			answers[i+1] = strconv.Itoa(newAnswer)
		} else {
			answers[i] = strconv.Itoa(newAnswer)
		}
	}

	result.Answers = answers
	result.Phrase = "When did Nobel prize winner " + titleFromID(peopleIDs[selectedPersonIndex]) + " die?"

	return result
}

func getPeopleIDsAndDeathYears() (deathYears []int, peopleIDs []int) {
	resp, err := http.Get("http://wdq.wmflabs.org/api?q=CLAIM[166:(TREE[7191][][279])]%20AND%20BETWEEN[570,2000,2015]&props=570")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var items peopleDeathDateResponse
	if err := json.Unmarshal(body, &items); err != nil {
		fmt.Println("error:", err)
	}

	peopleIDs = items.Items

	actualItems := items.Props.P570

	for i := 0; i < len(actualItems); i++ {
		timeString := actualItems[i].([]interface{})[2].(string)

		t, err := strconv.Atoi(timeString[8:12])
		if err != nil {
			fmt.Printf("%v", err)
		}
		deathYears = append(deathYears, t)
	}
	return
}
