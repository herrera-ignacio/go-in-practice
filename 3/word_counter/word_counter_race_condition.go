package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	// wait group to monitor a group of goroutines
	var wg sync.WaitGroup

	w := newWords()

	/**
	Open each file and tally the words. If we open multiple files, sometimes it will fail
	with a `concurrent map writes` error. This is because the words' map is
	being accessed by multiple goroutines at the same time.
	*/
	for _, filename := range os.Args[1:] {
		wg.Add(1)

		go func(file string) {
			if err := tallyWords(file, w); err != nil {
				fmt.Println(err.Error())
			}
		}(filename)
	}
	wg.Wait()

	fmt.Println("Words that appear more than once:")

	for word, count := range w.found {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
}

type words struct {
	found map[string]int
}

/**
Creates a new words instance
*/
func newWords() *words {
	return &words{found: map[string]int{}}
}

/**
Tracks the number of times you've seen this word
*/
func (w *words) add(word string, n int) {
	count, ok := w.found[word]

	if !ok {
		/**
		If the word isn't already tracked, add it.
		Otherwise, increment the count.
		*/
		w.found[word] = n
		return
	}

	w.found[word] = count + n
}

/**
Open a file, parse its contents, and count the words that appear.
Copy function does all the copying for you
*/
func tallyWords(filename string, dict *words) error {
	file, err := os.Open(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	/**
	Scanner is a useful tool for parsing files like this
	*/
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}

	return scanner.Err()
}
