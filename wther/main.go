package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/joho/godotenv"

	"wther/internal/weather"
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

	flag.Parse()
	args := flag.Args()

	var location string
	if len(args) == 0 {
		// TODO setup configuration reading, for now this is defaulted to Montreal
		location = "Montreal"
	} else {
		location = strings.TrimSpace(strings.Join(args, " "))
	}

	queryOptions := weather.ForecastQueryOptions{
		Location: location,
		Days:     1, // TODO experiment and update default, and also take flag
	}

	if err := loadApiKey(); err != nil {
		fmt.Printf("an error has occurred: %v\n", err)
		os.Exit(1)
	}

	if _, err := tea.NewProgram(tui.NewModel(queryOptions)).Run(); err != nil {
		fmt.Printf("an error has occurred: %v\n", err)
		os.Exit(1)
	}
}
