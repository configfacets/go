# Configfacets - Go Client Library

## Overview

As applications scale and integrate with multiple systems, managing configurations becomes increasingly complex. Configfacets simplifies this with a Low-Code/No-Code configuration management system using plain JSON—no custom verbs, no complicated syntax. This Go client library facilitates seamless interaction with the Configfacets API, enabling efficient retrieval and management of configuration data.

Our key features are...

**Repositories & Versioning:**
Design configurations as modular, reusable components, store them in a centralized repository, and maintain full version control for better organization and tracking.

**Reusability:**
Add provider and community-contributed repositories as dependencies, reuse configuration templates, pass in customizable values to effortlessly set up and manage your application configurations.

**Collaboration:**
Invite users and teams to repository with precise role-based permissions—Admin, Collaborator, or Viewer—to control access and streamline contributions.

**REST APIs:**
Expose configurations through REST API endpoints. Generate static and dynamic configurations by leveraging facet filters and runtime configuration values in the request context.

**Organization Management:**
Our hierarchical design simplifies managing multi-level organizational structures, team hierarchies, roles, and responsibilities.

## Usage

### Installation

```sh
go get -u github.com/configfacets/go@latest
```

### Fetching Configuration Data

Create a new configuration client and fetch data:

```go
package main

import (
	"fmt"
	"log"

	"github.com/configfacets/go/configfacets"
)

func main() {
	source := "https://configfacets.com/apis/repos/configfacets/core-concepts/appconfigs/resources/collections/feature-flags/exec?format=json"
	sourceType := "url"
	apiKey := "<your_api_key>"
	postBody := map[string]interface{}{
		"facets": []string{"env:prod", "country:CA"},
	}

	config := configfacets.NewConfiguration(source, sourceType, apiKey, postBody)
	err := config.Fetch()
	if err != nil {
		log.Fatalf("Error fetching config: %v", err)
	}

	// Retrieve a nested value
	value := config.GetValue("theme.is_dark_mode_enabled")
	fmt.Printf("Is dark mode enabled: %v\n", value)
}
```

## API Reference

### `NewConfiguration(source, sourceType, apiKey string, postBody map[string]interface{}) *Configuration`

Initializes a new Configfacets configuration client.

### `func (c *Configuration) Fetch() error`

Fetches the configuration data from the API.

### `func (c *Configuration) GetValue(keyPath string) interface{}`

Retrieves the value for the specified key path. If the key is not found, it returns `nil` and logs a warning.

## Contributing

We welcome contributions!, feel free to connect with us in our [Discord community](https://discord.gg/zWj3Rzud5s)
