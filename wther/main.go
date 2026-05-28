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

// is set at build time using -ldflags, is overridden by .env in loadApiKey()
var VersionDefaultApiKey string

// overrides VersionDefaultApiKey
func loadApiKey() error {

	_ = godotenv.Load()

	if envKey, ok := os.LookupEnv("WEATHER_API_KEY"); ok && envKey != "" {
		return nil
	}

	// if the key is not available in a .env, ensure that it was set at build time
	if VersionDefaultApiKey != "" {
		os.Setenv("WEATHER_API_KEY", VersionDefaultApiKey)
		return nil
	}

	return errors.New("WEATHER_API_KEY not found in environment, .env, or binary build flags")
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
