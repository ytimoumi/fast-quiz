package cmd

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// apiStartURL is the starting fragment of API's URL.
var apiStartURL = "http://localhost:9444/v1"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "A CLI interface to the quiz API.",
	Long:  "A CLI interface to the quiz API.",
}

// FetchQuestion fetches the question ID from API
func FetchQuestion(questionID string) (io.ReadCloser, error) {
	if questionIDnum, err := strconv.Atoi(questionID); err != nil || questionIDnum <= 0 {
		return nil, errors.New("question ID is not a valid ID")
	}

	questionURL := apiStartURL + "/questions/" + questionID
	resp, err := http.Get(questionURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP status is not OK")
	}

	return resp.Body, nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here we will define the flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for the application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
