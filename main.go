package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	flagWords = flag.Int("words", 5, "Number of words in a phrase")
	flagCount = flag.Int("count", 10, "Number of phrases to generate")
)

func rolls() string {
	return fmt.Sprintf("%d%d%d%d%d",
		rand.Intn(6)+1,
		rand.Intn(6)+1,
		rand.Intn(6)+1,
		rand.Intn(6)+1,
		rand.Intn(6)+1,
	)
}

func pickWord() string {
	return wordlist[rolls()]
}

func main() {
	flag.Parse()
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < *flagCount; i++ {
		words := make([]string, *flagWords)
		for j := range words {
			words[j] = pickWord()
		}
		fmt.Println(strings.Join(words, " "))
	}
}
