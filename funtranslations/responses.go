package funtranslations

import "fmt"

type assetsResponse struct {
	Contents ContentsData `json:"contents"`
	Success  SuccessData  `json:"success"`
}

type ContentsData struct {
	Translated  string `json:"translated"`
	Text        string `json:"text"`
	Translation string `json:"translation"`
}

func (c ContentsData) GetText() string {
	return fmt.Sprintf("Translation from english to %s\n Source text: %s\n Translation: %s",
		c.Translation, c.Text, c.Translated)
}

type SuccessData struct {
	Total int `json:"total"`
}

type assetsError struct {
	Contents ErrorData `json:"error"`
}

type ErrorData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ErrorData) GetText() string {
	return e.Message
}
