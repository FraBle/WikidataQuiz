package question

import (
	//standard library
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	// internal packages
	"github.com/FraBle/WikidataQuiz/model"
	"github.com/FraBle/WikidataQuiz/utility"
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
func NobelPrizeWinnersQuestion() (result model.Question, err error) {
	deathYears, peopleIDs := getPeopleIDsAndDeathYears()
	selectedPersonIndex := utility.Random(0, len(peopleIDs))

	result.RightAnswer = utility.Random(0, 4)
	var answers [4]string

	answers[result.RightAnswer] = strconv.Itoa(deathYears[selectedPersonIndex])

	offsets := utility.FourRandomNumbersIn(15) // one number not used

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

	person, err := titleFromID(peopleIDs[selectedPersonIndex])
	if err != nil {
		log.Printf("Error getting nobel prize winner title: %v", err)
		return
	}
	result.Phrase = "When did Nobel prize winner " + person + " die?"

	return
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
