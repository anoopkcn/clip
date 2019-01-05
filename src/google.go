package clip

import (
	// "fmt"
	"net/http"
	"strings"
)

func SearchGoogle(opts Options) {
	query := strings.Replace(opts.Search.String, " ", "+", -1)
	googleSearchQueryURL := "https://www.google.com/search?q=" + query + "&hl=lang_en&num=5"
	response, err := http.Get(googleSearchQueryURL)
	if err != nil {
		// fmt.Print(err.Error())
		errorExit("search failed. source down or no such host")
	}
	defer response.Body.Close()
	// responseData, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(responseData))
	// parse class=g> cite + class=st
	// var responseObject ArticleDetails
	// json.Unmarshal(responseData, &responseObject)

}
