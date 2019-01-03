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
		case "arxiv":
			fmt.Println("begin search...")
			// clp.SearchArxiv(opts)
		default:
			errorExit("please provide a valid source for searching")
		}
	}
}
