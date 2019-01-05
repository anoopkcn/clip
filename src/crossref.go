package clip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// http://api.crossref.org/works/<DOI>/transform/application/x-bibtex or json
type ArticleDetails struct {
	Doi                 string           `json:"DOI"`
	Url                 string           `json:"URL"`
	JournalIssue        JournalIssue     `json:"journal-issue"`
	Title               string           `json:"title"`
	Volume              string           `json:"volume"`
	ContainerTitle      string           `json:"container-title"`
	ContainerTitleShort string           `json:"container-title-short"`
	Issued              string           `json:"issued"`
	Authors             []CrossrefAuthor `json:"author"`
	Pages               string           `json:"page"`
	Issn                []string         `json:"ISSN"`
}
type JournalIssue struct {
	Issue          string    `json:"issue"`
	PublishedPrint DateParts `json:"published-print"`
}
type DateParts struct {
	DateParts [][]int `json:"date-parts"`
}

type CrossrefAuthor struct {
	Given  string `json:"given"`
	Family string `json:"family"`
}

func printParsedCrossrefJSON(ro ArticleDetails) {
	fmt.Printf("%s%s\n", "*", ro.Title)
	fmt.Printf("%-20s:%s\n", "  doi ", ro.Doi)
	fmt.Printf("%-20s:%s\n", "  url ", ro.Url)
	fmt.Printf("%-20s:%d\n", "  year ", ro.JournalIssue.PublishedPrint.DateParts[0][0])
	fmt.Printf("%-20s:%s\n", "  volume ", ro.Volume)
	fmt.Printf("%-20s:%s\n", "  issue ", ro.JournalIssue.Issue)
	fmt.Printf("%-20s:%s\n", "  pages ", ro.Pages)
	fmt.Printf("%-20s:%s\n", "  journal ", ro.ContainerTitle)
	fmt.Printf("%-20s:%s\n", "  journal-s ", ro.ContainerTitleShort)
	fmt.Printf("%-20s:", "  authors ")
	for i := 0; i < len(ro.Authors); i++ {
		fmt.Printf("%s %s, ", ro.Authors[i].Given, ro.Authors[i].Family)
	}
	fmt.Printf("\n")
	fmt.Printf("%-20s:%s\n", "  issn ", ro.Issn)
}

func SearchCrossref(opts Options) {
	crossrefSearchQueryURL := "http://api.crossref.org/works/" + opts.Search.String + "/transform/application/json"
	response, err := http.Get(crossrefSearchQueryURL)
	if err != nil {
		// fmt.Print(err.Error())
		errorExit("search failed. source down or no such host")
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject ArticleDetails
	json.Unmarshal(responseData, &responseObject)
	printParsedCrossrefJSON(responseObject)

}
