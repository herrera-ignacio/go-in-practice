package main

import (
	"errors"
	"io"
)

type File interface {
	Load(string) (io.ReadCloser, error)
	Save(string, io.ReadSeeker) error
}

type LocalFile interface {
	Base string
}

var (
	ErrFileNotFound   = errors.New("File not found")
	ErrCanontLoadFile = errors.New("Unable to load file")
	ErrCannotSaveFile = errors.New("Unable to save file")
)

func (l LocalFile)  Load(path string) (io.ReadCloser, error) {
	p := filepath.Join(l.Base, path)
	var oerr error
	o, err := os.Open(p)

	if err != nil && os.IsNotExist(err) {
		log.Printf("Unable to find %s", path)
		oerr = ErrFileNotFound
	} else if err != nil {
		log.Printf("Error loading file %s, err: %s", path, err)
		oerr = ErrCannotLoadFile
	}

	return o, oerr
}
