package question

import (
	"encoding/json"
	"fmt"
	"github.com/FraBle/WikidataQuiz/model"
	"io/ioutil"
	"math/rand"
	"net/http"
)

type countryCapitalResponse struct {
	Status struct {
		Error        string
		Items        int
		Querytime    string
		Parsed_query string
	}
	Items []int
	Props struct {
		P36 []interface{} `json:"36"`
	}
}

var countryIDs []int
var capitalIDs []int

func CapitalQuestion() *model.Question {
	result := new(model.Question)

	getCountryCapitalIDs()

	indexes := fourRandomNumbersIn(len(countryIDs))

	result.RightAnswer = rand.Intn(4)

	result.Phrase = "What is the capital of " + titleFromID(countryIDs[indexes[result.RightAnswer]]) + "?"
	result.Answers = [4]string{titleFromID(capitalIDs[indexes[0]]),
		titleFromID(capitalIDs[indexes[1]]),
		titleFromID(capitalIDs[indexes[2]]),
		titleFromID(capitalIDs[indexes[3]])}
	return result
}

func getCountryCapitalIDs() {
	resp, err := http.Get("http://wdq.wmflabs.org/api?q=claim[31:(tree[6256][][279])]&props=36")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var items countryCapitalResponse
	if err := json.Unmarshal(body, &items); err != nil {
		fmt.Println("error:", err)
	}

	actualItems := items.Props.P36

	for i := 0; i < len(actualItems); i++ {
		countryIDs = append(countryIDs, int(actualItems[i].([]interface{})[0].(float64)))
		capitalIDs = append(capitalIDs, int(actualItems[i].([]interface{})[2].(float64)))
	}
}
