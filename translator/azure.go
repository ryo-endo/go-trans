package translator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type transResponse []struct {
	Translations []struct {
		Text string `json:"text"`
		To   string `json:"to"`
	} `json:"translations"`
}

type errorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type sendBody struct {
	Text string `json:"Text"`
}

func NewAzure(key string) Translator {
	t := new(azure)
	t.key = key

	return t
}

type azure struct {
	key string
}

func (t *azure) Trans(s string, from string, to string) (string, error) {
	body, err := t.makeBody(s)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://api.cognitive.microsofttranslator.com/translate?api-version=3.0&from=%s&to=%s", from, to)

	resp, err := t.callTranslateApi(url, body, t.key)
	if err != nil {
		return "", fmt.Errorf("Cannot connect to the api server.\n%s", err.Error())
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		var result errorResponse
		if err := json.Unmarshal(bytes, &result); err != nil {
			return "", fmt.Errorf("The server returned an invalid JSON response.\n%s", err.Error())
		}
		return "", fmt.Errorf("The server returned an error respons.\ncode=%d\nmessage=%s", result.Error.Code, result.Error.Message)
	}

	var result transResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		return "", fmt.Errorf("The server returned an invalid JSON response.\n%s", err.Error())
	}

	out := result[0].Translations[0].Text
	return out, nil
}

func (t *azure) makeBody(s string) (string, error) {
	bodyText, err := json.Marshal([]sendBody{{s}})
	if err != nil {
		return "", err
	}

	return string(bodyText), nil
}

func (t *azure) callTranslateApi(url string, body string, key string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header = map[string][]string{
		"Content-Type":              {"application/json; charset=UTF-8"},
		"Ocp-Apim-Subscription-Key": {key},
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
