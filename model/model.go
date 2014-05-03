package model

type Question struct {
	Phrase      string   `json:"phrase"`
	Answers     []string `json:"answers"`
	RightAnswer int      `json:"rightAnswer"`
}
