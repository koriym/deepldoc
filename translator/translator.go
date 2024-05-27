package translator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// DeepL APIレスポンスの構造体
type DeepLResponse struct {
	Translations []struct {
		Text string `json:"text"`
	} `json:"translations"`
}

// 翻訳関数
func Translate(text string, targetLang string) (string, error) {
	apiKey := os.Getenv("DEEPL_API_KEY") // 環境変数からDeepL APIキーを読み込む
	url := "https://api-free.deepl.com/v2/translate"
	// リクエストボディの作成
	reqBody, err := json.Marshal(map[string]interface{}{
		"text":        []string{text},
		"target_lang": targetLang,
	})
	if err != nil {
		return "", err
	}
	// HTTP POSTリクエストの作成
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "DeepL-Auth-Key "+apiKey)
	// HTTPクライアントの作成
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// ステータスコードをチェック
	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}
	// レスポンスボディの読み取り
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// レスポンスのパース
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
