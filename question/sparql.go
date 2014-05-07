package question

import (
	// standard library
	"log"
	"strconv"
	"time"

	// external packages
	"github.com/knakk/fenster/sparql"

	// internal packages
	"github.com/FraBle/WikidataQuiz/config"
)

func getCount(result *sparql.Results) int {
	log.Print(result)
	for _, value := range result.Results.Bindings[0] {
		i, _ := strconv.Atoi(value.Value)
		return i
	}
	return -1
}

func query(query string) (result *sparql.Results, err error) {
	timeout := time.Duration(config.CONFIG.DBpediaEndpointTimeout) * time.Second
	byteResult, err := sparql.Query(config.CONFIG.DBpediaEndpoint, query, "json", timeout, timeout)
	if err != nil {
		return
	}
	result, err = sparql.ParseJSON(byteResult)
	return
}
