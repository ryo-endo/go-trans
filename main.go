package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

func main() {
	if err := Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func Run(args []string) error {
	from := flag.String("from", "en", `Set the language code. "en" "ja" "vn"`)
	to := flag.String("to", "ja", `Set the language code. "en" "ja" "vn"`)
	flag.Parse()

	if len(args) <= 1 {
		return fmt.Errorf(`usage: trans "Hello world!"`)
	}

	key := os.Getenv("TRANS_API_KEY")
	if key == "" {
		return fmt.Errorf("You may need to set TRANS_API_KEY.\n# export TRANS_API_KEY=your-api-key")
	}

	input := flag.Arg(0)
	out, err := trans(input, *from, *to, key)
	if err != nil {
		return err
	}

	fmt.Println(out)

	return nil
}

func trans(s string, from string, to string, key string) (string, error) {
	url := fmt.Sprintf("https://api.cognitive.microsofttranslator.com/translate?api-version=3.0&from=%s&to=%s", from, to)
	body := strings.NewReader(fmt.Sprintf("[{'Text':'%s'}]", s))

	resp, err := callTranslateApi(url, body, key)
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

func callTranslateApi(url string, body *strings.Reader, key string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header = map[string][]string{
		"Content-Type":              {"application/json"},
		"Ocp-Apim-Subscription-Key": {key},
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
