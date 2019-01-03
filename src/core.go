package clip

import (
	"fmt"
	"os"
)

func Run(opts Options, revision string) {
	if opts.Help || len(os.Args) < 2 {
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
}
