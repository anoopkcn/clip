package clip

import (
	"flag"
	"os"
	"strconv"
)

const usage = `usage: clip  [-version] [-help] <command> [<args>]

These are common clip commands and their arguments used in various situations:
`

// Options stores the values of command-line options
type Options struct {
	Version     bool
	Help        bool
	SearchBegin bool
	Search
}
type Search struct {
	Source  string
	String  string
	Match   string
	Offset  int
	Results int
	Filter  string
	Prefix  string
}

func help(code int) {
	os.Stderr.WriteString(usage)
	os.Exit(code)
}

func errorExit(msg string) {
	os.Stderr.WriteString(msg + "\n")
	os.Exit(exitError)
}

func atoi(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		errorExit("not a valid integer: " + str)
	}
	return num
}

func atof(str string) float64 {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		errorExit("not a valid number: " + str)
	}
	return num
}

func ParseOptions() Options {
	version := flag.Bool("version", false, "version information of clip")
	help := flag.Bool("help", false, "basic usage instructions")
	// search command group
	searchBegin := false
	searchCommand := flag.NewFlagSet("search", flag.ExitOnError)
	searchSource := searchCommand.String("source", "arxiv", "online repository to be searched")
	searchString := searchCommand.String("string", "", "search string, phrases, doi, etc,.")
	searchMatch := searchCommand.String("match", "phrase", "what  a given search string should match,[phrase|doi|title|author|issn]")
	searchFilter := searchCommand.String("filter", "all", "search filter type")
	searchPrefix := searchCommand.String("prefix", "", "value of the search filter")
	searchStart := searchCommand.Int("offset", 0, "offset for search results")
	searchResults := searchCommand.Int("results", 5, "number of results to be returned")
	// get command group

	switch os.Args[1] {
	case "search":
		searchCommand.Parse(os.Args[2:])
		if searchCommand.Parsed() {
			// Required Flags
			if *searchString != "" {
				searchBegin = true
			}
		}
	default:
		flag.Parse()
	}
	return Options{
		Version:     *version,
		Help:        *help,
		SearchBegin: searchBegin,
		Search: Search{
			Source:  *searchSource,
			String:  *searchString,
			Match:   *searchMatch,
			Filter:  *searchFilter,
			Prefix:  *searchPrefix,
			Offset:  *searchStart,
			Results: *searchResults,
		},
	}
}
