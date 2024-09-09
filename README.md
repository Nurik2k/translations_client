# Golang client for fun translations

fun_translation_client is a Go client library for accessing the https://api.funtranslations.com/

## Example:

```golang
package main

import (
	"fmt"
	"log"
	"time"
  
  	"github.com/BalamutDiana/fun_translations_client"
)

func main() {
	funtranslationsClient, err := funtranslations.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
	}

	lang := funtranslations.GetLanguagesList()
	text := "How was your day?"

	translation, err := funtranslationsClient.GetTranslation(lang.Shakespeare, text)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(translation)

}
```
