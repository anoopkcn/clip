package clip

import (
	"fmt"
	"os"
)

func Run(opts Options, revision string) {
	if opts.Help {
		help(exitOk)
	}
	if opts.Version {
		if len(revision) > 0 {
			fmt.Printf("%s (%s)\n", version, revision)
		} else {
			fmt.Println(version)
		}
		os.Exit(exitOk)
	}
	if opts.SearchBegin {
		switch opts.Search.Source {
		case "google":
			fmt.Println("Google: begin search...")
			SearchGoogle(opts)
		case "crossref":
			fmt.Println("Crossref: begin search...")
			SearchCrossref(opts)
		case "arxiv":
			fmt.Println("arXiv: begin search...")
			SearchArxiv(opts)
		default:
			errorExit("please provide a valid source for searching")
		}
	}
}
