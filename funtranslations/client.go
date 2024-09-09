package funtranslations

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Languages struct {
	Pirate      string
	Shakespeare string
	Yoda        string
}

func GetLanguagesList() *Languages {
	return &Languages{
		Pirate:      "pirate",
		Shakespeare: "shakespeare",
		Yoda:        "yoda",
	}
}

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}

	return &Client{
		client: &http.Client{
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
			Timeout: timeout,
		},
	}, nil
}

func (c Client) GetTranslation(language, text string) (string, error) {
	resp, err := marshalRequest(c, language, text)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	available, err := checkAvailability(resp)
	if err != nil {
		return "", err
	}

	if available {
		return unmarshalResponse(resp, assetsResponse{})
	} else {
		return unmarshalResponse(resp, assetsError{})
	}
}

func checkAvailability(resp *http.Response) (bool, error) {
	switch resp.StatusCode {
	case 200:
		return true, nil
	case 429:
		return false, nil
	default:
		return false, errors.New("something went wrong, check the request or try again later")
	}
}

func unmarshalResponse(resp *http.Response, data interface{}) (string, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if d, ok := data.(assetsResponse); ok {
		err = json.Unmarshal(body, &d)
		if err != nil {
			return "", err
		}
		return d.Contents.GetText(), nil
	} else if e, ok := data.(assetsError); ok {
		err = json.Unmarshal(body, &e)
		if err != nil {
			return "", err
		}
		return e.Contents.GetText(), nil
	} else {
		return "", errors.New("wrong data type")
	}

}

func marshalRequest(c Client, language, text string) (*http.Response, error) {
	values := map[string]string{"text": text}
	json_data, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.funtranslations.com/translate/%s.json", language)
	return c.client.Post(url, "application/json", bytes.NewBuffer(json_data))
}
