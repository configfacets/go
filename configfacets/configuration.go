package configfacets

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Configuration represents the configuration client
 type Configuration struct {
	Source     string
	SourceType string
	APIKey     string
	PostBody   map[string]interface{}
	Response   map[string]interface{}
}

// NewConfiguration initializes the configuration client
func NewConfiguration(source, sourceType, apiKey string, postBody map[string]interface{}) *Configuration {
	return &Configuration{
		Source:     source,
		SourceType: sourceType,
		APIKey:     apiKey,
		PostBody:   postBody,
	}
}

// Fetch retrieves the configuration from the source (API or file)
func (c *Configuration) Fetch() error {
	if c.Source == "" {
		return errors.New("missing source")
	}

	if c.SourceType == "file" {
		data, err := ioutil.ReadFile(c.Source)
		if err != nil {
			return err
		}

		var result map[string]interface{}
		if err := json.Unmarshal(data, &result); err != nil {
			return err
		}

		c.Response = result
		return nil
	} else if c.SourceType == "url" {
		return c.fetchFromURL()
	}

	return errors.New("invalid source type")
}

func (c *Configuration) fetchFromURL() error {
	jsonData, err := json.Marshal(c.PostBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.Source, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.APIKey != "" {
		req.Header.Set("X-APIKEY", c.APIKey)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}

	c.Response = result
	return nil
}

func (c *Configuration) GetValue(keyPath string) interface{} {
	if c.Response == nil {
		fmt.Println("[ERROR] Response is nil, did you call Fetch()? Returning nil.")
		return nil
	}

	value, err := getByPath(c.Response, keyPath)
	if err != nil {
		fmt.Printf("Warning: Key '%s' not found. Returning nil.\n", keyPath)
		return nil
	}
	return value
}