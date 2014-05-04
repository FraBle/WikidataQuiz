package question

import (
	"github.com/FraBle/WikidataQuiz/model"
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"math/rand"
	"time"
	"os"
)

type countryCapitalResponse struct {
	Status struct{
		Error string 
		Items int
		Querytime string
		Parsed_query string
	}
	Items []int
	Props struct{
		P36 []interface{} `json:"36"`
	}
}

var countryIDs []int
var capitalIDs []int


func CapitalQuestion() *model.Question {
	result := new(model.Question)

	getCountryCapitalIDs()

	var indexes []int
	for len(indexes) != 4  {
		rand.Seed(time.Now().UnixNano())
		newIndex := rand.Intn(len(countryIDs))
		if intInArray(newIndex, indexes) == false {
			indexes = append(indexes, newIndex)			//optimize ? - write directly in result
		}
	}

	rand.Seed(time.Now().UnixNano())
	result.RightAnswer = rand.Intn(4)


	for j := 0; j < len(countryIDs); j++ {
		fmt.Printf("Land: %d - Hauptstadt: %d\n", countryIDs[j], capitalIDs[j])
	} 

	result.Phrase = "What is the capitol of " + titleFromID(countryIDs[indexes[result.RightAnswer]]) + "?"
	result.Answers = []string{titleFromID(capitalIDs[indexes[0]]),
							  titleFromID(capitalIDs[indexes[1]]),
							  titleFromID(capitalIDs[indexes[2]]),
							  titleFromID(capitalIDs[indexes[3]])}
	return result
}


func titleFromID(ID int) string {
	response, err := http.Get("https://www.wikidata.org/w/api.php?action=wbgetentities&ids=Q" + strconv.Itoa(ID) + "&format=json&languages=en&props=labels")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
 
		value := make(map[string]interface{})
 
		err = json.Unmarshal(body, &value)
		if err != nil {
			fmt.Println(err)
		}
		return (((value["entities"]).(map[string]interface{})["Q" + strconv.Itoa(ID)]).(map[string]interface{})["labels"]).(map[string]interface{})["en"].(map[string]interface{})["value"].(string)
	}

	return ""
}


func intInArray(elem int, array []int) bool {
    for _, b := range array {
        if b == elem {
            return true
        }
    }
    return false
}


func getCountryCapitalIDs()  {
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
	
	for i:=0; i < len(actualItems); i++ {
		countryIDs = append(countryIDs, int(actualItems[i].([]interface{})[0].(float64)))
		capitalIDs = append(capitalIDs, int(actualItems[i].([]interface{})[2].(float64)))
	}
}



/*func countryCapitalIDMap()  map[int]int {
	resp, err := http.Get("http://wdq.wmflabs.org/api?q=claim[31:(tree[6256][][279])]&props=36")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)	
	var items map[string]interface{}
	if err := json.Unmarshal(body, &items); err != nil {
		fmt.Println("error:", err)
	}

	var actualItems [][]interface{}

	fmt.Printf("%f", (((items["props"]).(map[string]interface{})["36"]).([]interface{})[0]).([]interface{})[0].(float64))

}
*/