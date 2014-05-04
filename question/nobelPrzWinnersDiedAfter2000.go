package question

import (
	"github.com/FraBle/WikidataQuiz/model"
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"math/rand"
	"time"
)

type peopleDeathDateResponse struct {
	Status struct{
		Error string 
		Items int
		Querytime string
		Parsed_query string
	}
	Items []int
	Props struct{
		P570 []interface{} `json:"570"`
	}
}

// todo: start date and end data saved in variable/ constant


var peopleIDs []int
var deathYears []int

func NobelPrzWinnersDiedAfter2000Question() *model.Question {
	result := new(model.Question)

	getPeopleIDsAndDeathYears()

	rand.Seed(time.Now().UnixNano())
	selectedPersonIndex := rand.Intn(len(peopleIDs))

	rand.Seed(time.Now().UnixNano())
	result.RightAnswer = rand.Intn(4)

	answers := []string{"","","",""}

	answers[result.RightAnswer] = strconv.Itoa(deathYears[selectedPersonIndex])

	offsets := *fourRandomNumbersIn(15)    // one number not used

	for i:=0; i < 3; i++{
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

func getPeopleIDsAndDeathYears()  {
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
	
	for i:=0; i < len(actualItems); i++ {
		timeString := actualItems[i].([]interface{})[2].(string)

		t, err := strconv.Atoi(timeString[8:12])
		if err != nil {
			fmt.Printf("%v", err)
		}
		deathYears = append(deathYears, t)
	}
}
