package clip

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type SearchResults struct {
	XMLEntry      xml.Name       `xml:"feed"`
	TotalResults  int            `xml:"totalResults"`
	ItemsPerPage  int            `xml:"itemsPerPage"`
	SearchResults []SearchResult `xml:"entry"`
}

type SearchResult struct {
	XMLEntry   xml.Name `xml:"entry"`
	ArxivID    string   `xml:"id"`
	Published  string   `xml:"published"`
	Title      string   `xml:"title"`
	Summary    string   `xml:"summary"`
	Authors    []Author `xml:"author"`
	Doi        string   `xml:"doi,omitempty"`
	JournalRef string   `xml:"journal_ref,omitempty"`
}

type Author struct {
	XMLEntry xml.Name `xml:"author"`
	Name     string   `xml:"name"`
}

func printParsedArxivXML(ro SearchResults) {
	fmt.Println("Showing ", ro.ItemsPerPage, "of ", ro.TotalResults, "results")
	for i := 0; i < len(ro.SearchResults); i++ {
		fmt.Println("*" + ro.SearchResults[i].Title)
		fmt.Printf(" %-15s:", "authors")
		for j := 0; j < len(ro.SearchResults[i].Authors); j++ {
			fmt.Printf(" %s,", ro.SearchResults[i].Authors[j].Name)
		}
		fmt.Printf("\n")
		fmt.Printf(" %-15s: %s\n", "published", ro.SearchResults[i].Published)
		if ro.SearchResults[i].Doi != "" {
			fmt.Printf(" %-15s: %s\n", "doi", ro.SearchResults[i].Doi)
		}
		if ro.SearchResults[i].JournalRef != "" {
			fmt.Printf(" %-15s: %s\n", "journal", ro.SearchResults[i].JournalRef)
		}
		fmt.Printf(" %-15s: %s\n", "arxivID", ro.SearchResults[i].ArxivID)
		// fmt.Printf("%-3s%s\n", ":", strings.TrimSpace(ro.SearchResults[i].Summary))
		fmt.Printf("\n")
	}
}

func SearchArxiv(opts Options) {
	var base_url, filter, query, offset, results string
	base_url = "http://export.arxiv.org/api/query?search_query="
	filter = opts.Search.Filter + ":"
	if opts.Search.Filter != "all" {
		filter = filter + opts.Search.Prefix + "+AND+"
	}
	query = strings.Replace(opts.Search.String, " ", "+AND+", -1)
	offset = "&start=" + strconv.Itoa(opts.Search.Offset)
	results = "&max_results=" + strconv.Itoa(opts.Search.Results)

	arxivSearchQueryURL := base_url + filter + query + offset + results
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
	printParsedArxivXML(responseObject)
}
