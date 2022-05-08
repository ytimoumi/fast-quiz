package models

type Question struct {
	Question        string    `json:"question"`
	Answers         [3]string `json:"answers"`
	CorrectAnswerID int       `json:"-"`
}
