package utility

import (
	// standard library
	"strconv"
	"time"

	// external packages
	"github.com/knakk/fenster/sparql"

	// internal packages
	"github.com/FraBle/WikidataQuiz/config"
)

// GetCount() returns the amount of result values of a sparql result.
func GetCount(result *sparql.Results) int {
	for _, value := range result.Results.Bindings[0] {
		i, _ := strconv.Atoi(value.Value)
		return i
	}
	return -1
}

// Query() queries the  Dbpedia API.
func Query(query string) (result *sparql.Results, err error) {
	timeout := time.Duration(config.CONFIG.DBpediaEndpointTimeout) * time.Second
	byteResult, err := sparql.Query(config.CONFIG.DBpediaEndpoint, query, "json", timeout, timeout)
	if err != nil {
		return
	}
	result, err = sparql.ParseJSON(byteResult)
	return
}
