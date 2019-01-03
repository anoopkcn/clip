package clip

import (
	"fmt"
)

func Run(revision string) {
	fmt.Printf("%s (%s)\n", version, revision)
}
