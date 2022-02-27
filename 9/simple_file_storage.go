package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type File interface {
	Load(string) (io.ReadCloser, error)
	Save(string, io.ReadSeeker) error
}

type LocalFile struct {
	Base string
}

func (l LocalFile) Load(path string) (io.ReadCloser, error) {
	p := filepath.Join(l.Base, path)
	return os.Open(p)
}

func (l LocalFile) Save(path string, body io.ReadSeeker) error {
	/**
	Saving files requires making sure the local directory exists for the requested path
	and copying the contents to saved file location
	*/
	p := filepath.Join(l.Base, path)
	d := filepath.Dir(p)
	err := os.MkdirAll(d, os.ModeDir|os.ModePerm)

	if err != nil {
		return err
	}

	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, body)
	return err
}

func main() {
	content := "test"
	body := bytes.NewReader([]byte(content))

	// Retrieves a store implementing the File interface
	store, err := fileStore()
	if err != nil {
		panic(err)
	}

	// Saves the example content to the file store
	fmt.Println("Storing content...")
	err = store.Save("foo/bar", body)
	if err != nil {
		panic(err)
	}

	// Retrieves and prints out content from the file store
	fmt.Println("Retrieving content...")
	c, err := store.Load("foo/bar")
	if err != nil {
		panic(err)
	}

	o, err := ioutil.ReadAll(c)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(o))

}
