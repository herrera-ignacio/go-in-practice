package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	w := newWords()

	for _, filename := range os.Args[1:] {
		wg.Add(1)

		go func(file string) {
			if err := tallyWords(file, w); err != nil {
				fmt.Println(err.Error())
			}

			wg.Done()
		}(filename)
	}
	wg.Wait()

	fmt.Println("Words that appear more than once:")

	/**
	Locks and unlocks the map when you iterate at the end.
	Stricly speaking, this isn't necessary because you know
	that this section won't happen until all files are processed.
	*/
	w.Lock()
	for word, count := range w.found {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
	w.Unlock()
}

type words struct {
	sync.Mutex // the words struct now inherits the mutex lock
	found      map[string]int
}

func newWords() *words {
	return &words{found: map[string]int{}}
}

/**
Locks the object, modifies the map,
and then unlocks the object.
*/
func (w *words) add(word string, n int) {
	w.Lock()
	defer w.Unlock()

	count, ok := w.found[word]

	if !ok {
		w.found[word] = n
		return
	}

	w.found[word] = count + n
}

/**
Open a file, parse its contents, and count the words that appear.
*/
func tallyWords(filename string, w *words) error {
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
		w.add(word, 1)
	}

	return scanner.Err()
}
