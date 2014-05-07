package question

import (
	// standard library
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	// internal packages
	"github.com/FraBle/WikidataQuiz/model"
	"github.com/FraBle/WikidataQuiz/utility"
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

// CapitalQuestion() generates a question about the capital of a country.
// It offers different capitals whereby one is correct.
func CapitalQuestion() (result model.Question, err error) {
	countryIDs, capitalIDs, err := getCountryCapitalIDs()
	if err != nil {
		return
	}

	indexes := utility.FourRandomNumbersIn(len(countryIDs))

	result.RightAnswer = utility.Random(0, 4)

	country, err := utility.TitleFromID(countryIDs[indexes[result.RightAnswer]])
	if err != nil {
		log.Printf("Error getting country title: %v", err)
		return
	}

	result.Phrase = "What is the capital of " + country + "?"

	for i, val := range indexes {
		capital, e := utility.TitleFromID(capitalIDs[val])
		if e != nil {
			log.Printf("Error getting capital title: %v", err)
			return result, e
		}
		result.Answers[i] = capital
	}

	return
}

// getCountryCapitalIDs() is a helper function which calls the wmflabs API to aggregate the appropiate country IDs.
func getCountryCapitalIDs() (countryIDs, capitalIDs []int, err error) {
	resp, err := http.Get("http://wdq.wmflabs.org/api?q=claim[31:(tree[6256][][279])]&props=36")
	if err != nil {
		log.Printf("Error calling wmflabs API: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var items countryCapitalResponse
	if err := json.Unmarshal(body, &items); err != nil {
		log.Printf("Error unmarshaling wmflabs response: %v", err)
	}

	actualItems := items.Props.P36

	for i := 0; i < len(actualItems); i++ {
		countryIDs = append(countryIDs, int(actualItems[i].([]interface{})[0].(float64)))
		capitalIDs = append(capitalIDs, int(actualItems[i].([]interface{})[2].(float64)))
	}
	return
}
