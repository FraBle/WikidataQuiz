package question

import (
	// standard library
	"strconv"

	// internal packages
	"github.com/FraBle/WikidataQuiz/model"
	"github.com/FraBle/WikidataQuiz/utility"
)

// WorldCupQuestion() generates a question about the winner of the world cup in a random year.
func WorldCupQuestion() (result model.Question, err error) {
	q := `PREFIX dbpedia2: <http://dbpedia.org/property/>
SELECT COUNT(DISTINCT ?t)
WHERE {
  ?t dbpedia2:tourneyName ?name .
  FILTER (STR(?name) = "FIFA World Cup")
}`
	offsetResult, err := utility.Query(q)
	if err != nil {
		return
	}
	count := utility.GetCount(offsetResult)
	offset := utility.Random(0, count)
	q = `PREFIX dbpedia2: <http://dbpedia.org/property/>
SELECT ?year ?first ?second ?third ?fourth
WHERE {
  ?t dbpedia2:tourneyName ?name .
  FILTER (STR(?name) = "FIFA World Cup")
  ?t dbpedia2:champion ?first .
  ?t dbpedia2:year ?year .
  ?t dbpedia2:second ?second .
  ?t dbpedia2:third ?third .
  ?t dbpedia2:fourth ?fourth .
}
LIMIT 1
OFFSET ` + strconv.Itoa(offset)
	results, err := utility.Query(q)
	if err != nil {
		return
	}
	year := results.Results.Bindings[0]["year"].Value
	first := results.Results.Bindings[0]["first"].Value
	second := results.Results.Bindings[0]["second"].Value
	third := results.Results.Bindings[0]["third"].Value
	fourth := results.Results.Bindings[0]["fourth"].Value
	list := [...]string{second, third, fourth}

	indexes := utility.FourRandomNumbersIn(4)
	result.RightAnswer = indexes[0]
	result.Answers[indexes[0]] = first
	for i := 1; i < 4; i++ {
		result.Answers[indexes[i]] = list[i-1]
	}
	result.Phrase = "Who won the FIFA world championship in " + year + "?"
	return
}
