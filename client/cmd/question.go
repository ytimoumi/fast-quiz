package cmd

import (
	"github.com/spf13/cobra"
)

// getquestionCmd represents the getquestion command
var getquestionCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(getquestionCmd)

	getquestionCmd.Flags().StringP("id", "i", "", "Question ID")
}
