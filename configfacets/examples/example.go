package main

import (
	"fmt"
	"log"

	"github.com/configfacets/go/configfacets"
)

func main() {
	apiUrl := "https://configfacets.com/apis/repos/configfacets/core-concepts/appconfigs/resources/collections/feature-flags/exec?format=json"
	apiKey := "<your_api_key>"
	postBody := map[string]interface{}{
		"facets": []string{"env:prod", "country:CA"},
	}

	config := configfacets.NewConfiguration(apiUrl, apiKey, postBody)
	err := config.Fetch()
	if err != nil {
		log.Fatalf("Error fetching config: %v", err)
	}

	// Retrieve a nested value
	value := config.GetValue("theme.is_dark_mode_enabled")
	fmt.Printf("Is dark mode enabled: %v\n", value)
}
