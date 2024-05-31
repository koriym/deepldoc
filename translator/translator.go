package translator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// Structure of DeepL API response
type DeepLResponse struct {
	Translations []struct {
		Text string `json:"text"`
	} `json:"translations"`
}

func TranslateTextWithExclusions(text, targetLang string) (string, error) {
	re := regexp.MustCompile("(?s)(`.*?`|```.*?```)")
	parts := re.Split(text, -1)

	for i, part := range parts {
		if !strings.HasPrefix(part, "`") {
			result, err := Translate(part, targetLang)
			if err != nil {
				return "", err
			}
			parts[i] = result
		}
	}

	matches := re.FindAllString(text, -1)

	if len(matches) != 0 {
		for i, match := range matches {
			parts = insert(parts, match, 2*i+1)
		}
	}

	return strings.Join(parts, ""), nil
}

// insert function to add back original code blocks
func insert(slice []string, element string, index int) []string {
	slice = append(slice, "")
	copy(slice[index+1:], slice[index:])
	slice[index] = element
	return slice
}

// Translation function
func Translate(text string, targetLang string) (string, error) {
	apiKey := os.Getenv("DEEPL_API_KEY") // Load the DeepL API key from environment variables
	if apiKey == "" {
		fmt.Println("API key is empty. Please set the DEEPL_API_KEY environment variable.")
		os.Exit(1)
	}
	url := "https://api-free.deepl.com/v2/translate"
	// Create request body
	reqBody, err := json.Marshal(map[string]interface{}{
		"text":        []string{text},
		"target_lang": targetLang,
	})
	if err != nil {
		return "", err
	}
	// Create HTTP POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "DeepL-Auth-Key "+apiKey)
	// Create HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// Check status code
	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}
	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// Parse response
	var deeplResp DeepLResponse
	err = json.Unmarshal(body, &deeplResp)
	if err != nil {
		return "", err
	}
	if len(deeplResp.Translations) > 0 {
		return deeplResp.Translations[0].Text, nil
	}
	return "", fmt.Errorf("translation not found")
}
