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
	Version bool
	Help    bool
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
	Version := flag.Bool("version", false, "version information of clip")
	Help := flag.Bool("help", false, "basic usage instructions")
	flag.Parse()
	return Options{
		Version: *Version,
		Help:    *Help,
	}
}
