package translator
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Structure of DeepL API response
type DeepLResponse struct {
	Translations []struct {
		Text string `json:"text"`
	} `json:"translations"`
}

// Translation function
func Translate(text string, targetLang string) (string, error) {
	apiKey := os.Getenv("DEEPL_API_KEY") // Load the DeepL API key from environment variables
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
