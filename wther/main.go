package main

import (
	"errors"
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/joho/godotenv"

	"wther/tui"
)

const url = "https://charm.sh/"

func loadApiKey() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	if _, ok := os.LookupEnv("WEATHER_API_KEY"); !ok {
		return errors.New("WEATHER_API_KEY not set")
	}

	return nil

}

func main() {
	if err := loadApiKey(); err != nil {
		fmt.Printf("an error has occurred: %v\n", err)
		os.Exit(1)
	}

	if _, err := tea.NewProgram(tui.NewModel()).Run(); err != nil {
		fmt.Printf("an error has occurred: %v\n", err)
		os.Exit(1)
	}
}
