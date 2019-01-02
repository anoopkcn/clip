package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
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

type SearchResults struct {
	XMLEntry      xml.Name       `xml:"feed"`
	TotalResults  int            `xml:"totalResults"`
	ItemsPerPage  int            `xml:"itemsPerPage"`
	SearchResults []SearchResult `xml:"entry"`
}

type SearchResult struct {
	XMLEntry   xml.Name `xml:"entry"`
	ArxiveID   string   `xml:"id"`
	Published  string   `xml:"published"`
	Title      string   `xml:"title"`
	Summary    string   `xml:"summary"`
	Authors    []Autho  `xml:"author"`
	Doi        string   `xml:"doi"`
	JournalRef string   `xml:"journal_ref"`
}

type Autho struct {
	XMLEntry xml.Name `xml:"author"`
	Name     string   `xml:"name"`
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

func print_xml(ro SearchResults) {
	fmt.Println("Showing ", ro.ItemsPerPage, "of ", ro.TotalResults, "results")
	fmt.Printf("\n")
	for i := 0; i < len(ro.SearchResults); i++ {
		fmt.Println("*", ro.SearchResults[i].Title, "(", ro.SearchResults[i].Published, ")")
		fmt.Printf("%-3s|", "")
		for j := 0; j < len(ro.SearchResults[i].Authors); j++ {
			fmt.Printf("%s, ", ro.SearchResults[i].Authors[j].Name)
		}
		fmt.Printf("\n")
		fmt.Printf("%-3s|%s\n", "", ro.SearchResults[i].Doi)
		fmt.Printf("%-3s|%s\n", "", ro.SearchResults[i].JournalRef)
		fmt.Printf("%-3s|%s\n", "", ro.SearchResults[i].ArxiveID)
		// fmt.Printf("%-3s%s\n", ":", strings.TrimSpace(ro.SearchResults[i].Summary))
		fmt.Printf("\n")
	}
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
func arxive_search(search_term, search_prefix, search_prefix_value *string, search_offset, search_count *int) {
	base_url := "http://export.arxiv.org/api/query?search_query="
	query := strings.Replace(*search_term, " ", "+AND+", -1)
	prefix := *search_prefix + ":"
	offset := "&start=" + strconv.Itoa(*search_offset)
	count := "&max_results=" + strconv.Itoa(*search_count)

	arxive_search_query_url := base_url + prefix + query + offset + count
	response, err := http.Get(arxive_search_query_url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(responseData))
	var responseObject SearchResults
	xml.Unmarshal(responseData, &responseObject)
	print_xml(responseObject)
}

func lazyclip_usage() {
	fmt.Printf("Usage of %s:\n\n", os.Args[0])
	fmt.Printf("    lazyclip <OPTIONS> [ARGUMENTS]\n\n")
	fmt.Printf("OPTIONS: \n\n")
	flag.PrintDefaults()
	fmt.Printf("\n")

}

var doi *string
var search_term *string
var search_prefix *string
var search_prefix_value *string
var search_count *int
var search_offset *int

func init_flags() {
	doi = flag.String("d", "", " DOI of the paper")
	search_term = flag.String("s", "", " String to be searched, in double quotes")
	search_prefix = flag.String("t", "all", " Subject Category to be searched, in double quotes/without space")
	search_count = flag.Int("c", 5, "Number of results, used in conjunction with -c")
	search_offset = flag.Int("o", 0, "Search offset, used in conjunction with -c")
}

func main() {
	init_flags()
	flag.Parse()
	if flag.Lookup("d") != nil && *doi != "" {
		doi_lookup(doi)
	} else if flag.Lookup("s") != nil && *search_term != "" {
		arxive_search(search_term, search_prefix, search_prefix_value, search_offset, search_count)
	} else {
		lazyclip_usage()
	}
}
