package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var file string
	i := -1

	for i, file = range os.Args[1:] {
		/**
		For every file you add, you tell the wait group
		that you're waiting for one more compress operation
		*/
		wg.Add(1)

		/**
		This function calls compress and then notifies
		the wait group that it's done.
		You want each goroutine to be scheduled with that iteration's
		value of file, so you pass it as a function parameter.
		*/
		go func(filename string) {

			compress(filename)
			wg.Done()
		}(file)
	}

	/**
	The outer goroutine (main) waits until all the
	compressing goroutines have called wg.Done
	*/
	wg.Wait()

	fmt.Printf("Compressed %d file\n", i+1)
}

func compress(filename string) error {
	in, err := os.Open(filename)

	if err != nil {
		return err
	}

	defer in.Close()

	out, err := os.Create(filename + ".gz")

	if err != nil {
		return err
	}

	defer out.Close()

	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)

	gzout.Close()

	return err
}
