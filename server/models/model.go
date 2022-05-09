package models

type Question struct {
	Question        string    `json:"question"`
	Answers         [3]string `json:"answers"`
	CorrectAnswerID int       `json:"-"`
}

// UsersAnsweredCorrectly Number of users correctly answered this amount of questions.
var UsersAnsweredCorrectly [4]int

// AnsweredUsers has total amount of users taken the quiz.
var AnsweredUsers = 0
