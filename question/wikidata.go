package question

import (
	// standard library
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func titleFromID(ID int) (title string, err error) {
	response, err := http.Get("https://www.wikidata.org/w/api.php?action=wbgetentities&ids=Q" + strconv.Itoa(ID) + "&format=json&languages=en&props=labels")
	if err != nil {
		log.Printf("Error calling Wikidata API: %v", err)
		return
	} else {
		defer response.Body.Close()
		body, e := ioutil.ReadAll(response.Body)
		if e != nil {
			log.Printf("Error reading Wikidata response: %v", err)
			return
		}
		responseMap := make(map[string]interface{})
		e = json.Unmarshal(body, &responseMap)
		if e != nil {
			log.Printf("Error unmarshaling Wikidata title: %v", err)
			return
		}
		title = (((responseMap["entities"]).(map[string]interface{})["Q"+strconv.Itoa(ID)]).(map[string]interface{})["labels"]).(map[string]interface{})["en"].(map[string]interface{})["value"].(string)
	}
	return
}
