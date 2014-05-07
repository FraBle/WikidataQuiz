package question

import (
	// standard library
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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
		return (((value["entities"]).(map[string]interface{})["Q"+strconv.Itoa(ID)]).(map[string]interface{})["labels"]).(map[string]interface{})["en"].(map[string]interface{})["value"].(string)
	}

	return ""
}
