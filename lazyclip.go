package main

import (
	"encoding/json"
	_ "encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// http://export.arxiv.org/api/query?search_query=<query:terms>&start=0&max_results=10
// http://api.crossref.org/works/<DOI>/transform/application/x-bibtex or json
type ArticleDetails struct {
	Doi                 string       `json:"DOI"`
	Url                 string       `json:"URL"`
	JournalIssue        JournalIssue `json:"journal-issue"`
	Title               string       `json:"title"`
	Volume              string       `json:"volume"`
	ContainerTitle      string       `json:"container-title"`
	ContainerTitleShort string       `json:"container-title-short"`
	Issued              string       `json:"issued"`
	Authors             []Author     `json:"author"`
	Pages               string       `json:"page"`
	Issn                []string     `json:"ISSN"`
}
type JournalIssue struct {
	Issue          string    `json:"issue"`
	PublishedPrint DateParts `json:"published-print"`
}
type DateParts struct {
	DateParts [][]int `json:"date-parts"`
}

type Author struct {
	Given  string `json:"given"`
	Family string `json:"family"`
}

func print_json(ro ArticleDetails) {
	fmt.Printf("%-15s:%s\n", "doi ", ro.Doi)
	fmt.Printf("%-15s:%s\n", "url ", ro.Url)
	fmt.Printf("%-15s:%d\n", "year ", ro.JournalIssue.PublishedPrint.DateParts[0][0])
	fmt.Printf("%-15s:%s\n", "volume ", ro.Volume)
	fmt.Printf("%-15s:%s\n", "issue ", ro.JournalIssue.Issue)
	fmt.Printf("%-15s:%s\n", "pages ", ro.Pages)
	fmt.Printf("%-15s:%s\n", "title ", ro.Title)
	fmt.Printf("%-15s:%s\n", "journal ", ro.ContainerTitle)
	fmt.Printf("%-15s:%s\n", "journal-short ", ro.ContainerTitleShort)
	fmt.Printf("%-15s:", "authors ")
	for i := 0; i < len(ro.Authors); i++ {
		fmt.Printf("%s %s, ", ro.Authors[i].Given, ro.Authors[i].Family)
	}
	fmt.Printf("\n")
	fmt.Printf("%-15s:%s\n", "issn ", ro.Issn)
}

var doi *string
var search_term *string

func init_flags() {
	doi = flag.String("d", "", " DOI of the paper")
	search_term = flag.String("s", "", " String to be searched in double quotes")

}

func doi_lookup(doi *string) {
	crossref_url_json := "http://api.crossref.org/works/" + *doi + "/transform/application/json"
	response, err := http.Get(crossref_url_json)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject ArticleDetails
	json.Unmarshal(responseData, &responseObject)
	print_json(responseObject)

}
func arxive_search(search_term *string) {
	query := strings.Replace(*search_term, " ", ":", -1)
	arxive_url_xml := "http://export.arxiv.org/api/query?search_query=" + query + "&start=0&max_results=2"
	response, err := http.Get(arxive_url_xml)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}

func lazyclip_usage() {
	fmt.Printf("Usage of %s:\n\n", os.Args[0])
	fmt.Printf("    lazyclip <OPTIONS> [ARGUMENTS]\n\n")
	fmt.Printf("OPTIONS: \n\n")
	flag.PrintDefaults()
	fmt.Printf("\n")

}

func main() {
	init_flags()
	flag.Parse()

	if flag.Lookup("d") != nil && *doi != "" {
		doi_lookup(doi)
	} else if flag.Lookup("s") != nil && *search_term != "" {
		arxive_search(search_term)
	} else {
		lazyclip_usage()
	}
}
