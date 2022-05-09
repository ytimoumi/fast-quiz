package cmd

import (
	"github.com/spf13/cobra"
	"io"
)

// Question 's structure
type Question struct {
	Question string    `json:"question"`
	Answers  [3]string `json:"answers"`
}

// Answer 's structure
type Answer struct {
	AnswerID int `json:"answerId"`
}

// Answers 's structure
type Answers [3]Answer

// answerCmd represents the answer command
var answerCmd = &cobra.Command{
	// todo

}

// Send our answers to the API for processing
func sendAnswer(questionID string, answers Answers) (io.ReadCloser, error) {
	// todo
}

// init :)
func init() {
	rootCmd.AddCommand(answerCmd)

	answerCmd.Flags().StringP("id", "i", "", "Question ID")
}
