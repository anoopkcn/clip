package main

import "github.com/strivetobelazy/clip/src"

var revision string

func main() {
	clip.Run(clip.ParseOptions(), revision)
}
