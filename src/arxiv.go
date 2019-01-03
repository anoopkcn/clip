package clip

import (
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

func printArxivXML(ro SearchResults) {
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

func SearchArxiv(opts Options) {
	var base_url, prefix, query, offset, results string
	base_url = "http://export.arxiv.org/api/query?search_query="
	prefix = opts.Search.Prefix + ":"
	if flag.Lookup("prefix") != nil && prefix != "all" {
		prefix = prefix + opts.Search.PrefixValue + "+AND+"
	}
	query = strings.Replace(opts.Search.String, " ", "+AND+", -1)
	offset = "&start=" + strconv.Itoa(opts.Search.Start)
	results = "&max_results=" + strconv.Itoa(opts.Search.Results)

	arxivSearchQueryURL := base_url + prefix + query + offset + results
	response, err := http.Get(arxivSearchQueryURL)
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
	printArxivXML(responseObject)
}
