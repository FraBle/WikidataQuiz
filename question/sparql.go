// You can edit this code!
// Click here and start typing.
package question

import (
	"github.com/knakk/fenster/sparql"
	"time"
	"io/ioutil"
	"strconv"
)

func getCount(result *sparql.Results) int {
	for _, value := range result.Results.Bindings[0] {
		i, _ := strconv.Atoi(value.Value)
		return i
	}
	return -1
}

func query(q string) *sparql.Results {
	endpoint := "http://dbpedia.org/sparql"
	tenSec := time.Duration(10)*time.Second
	byteRes, _ := sparql.Query(endpoint, q, "json", tenSec, tenSec)
	res, _ := sparql.ParseJSON(byteRes)
	return res
}

func readFile(filename string) string {
	byteRes, _ := ioutil.ReadFile(filename)
	return string(byteRes)
}