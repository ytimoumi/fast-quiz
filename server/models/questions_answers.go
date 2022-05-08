package models

// QuestionsMap is a list of questions to send to the client. DB emulation
var QuestionsMap = map[int][3]Question{
	1: {
		{
			Question:        "What is the capital city of Tunisia ?",
			Answers:         [3]string{"Sfax", "Tunis", "Kairouan"},
			CorrectAnswerID: 1,
		},
		{
			Question:        "How many countries are there in the region of Europe ?",
			Answers:         [3]string{"42", "43", "44"},
			CorrectAnswerID: 2,
		},
		{
			Question:        "How many elements are in the periodic table ?",
			Answers:         [3]string{"118", "117", "119"},
			CorrectAnswerID: 0,
		},
	},
	2: {
		{
			Question:        "What language is spoken in Brazil ?",
			Answers:         [3]string{"Portuguese", "English", "Brazilian"},
			CorrectAnswerID: 0,
		},
		{
			Question:        "How many keys are there on a piano ?",
			Answers:         [3]string{"86", "88", "90"},
			CorrectAnswerID: 1,
		},
		{
			Question:        "How many hearts does an octopus have ?",
			Answers:         [3]string{"3", "2", "1"},
			CorrectAnswerID: 0,
		},
	},
}
