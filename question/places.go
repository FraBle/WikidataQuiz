package question

import (
	"github.com/FraBle/WikidataQuiz/model"
)

func CapitalQuestion() *model.Question {
	return &model.Question{"Ich bin eine Beispielfrage?", []string{"Antwort1", "Antwort2", "Antwort3", "Antwort4"}, 2}
}
