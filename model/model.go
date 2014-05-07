package model

// Question represents a question json response.
type Question struct {
	Phrase      string    `json:"phrase"`
	Answers     [4]string `json:"answers"`
	RightAnswer int       `json:"rightAnswer"`
}
