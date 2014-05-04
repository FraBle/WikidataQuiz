package question

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"os"
	"time"
	"math/rand"
)

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

func fourRandomNumbersIn(area int) *[]int {  // actually [4]int
	var result []int
	for len(result) != 4  {
		rand.Seed(time.Now().UnixNano())
		newNumber := rand.Intn(area)
		if intInArray(newNumber, result) == false {
			result = append(result, newNumber)
		}
	}
	return &result
}

