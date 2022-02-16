package main

import (
	"io"
	"time"
)

type CacheFile struct {
	content io.ReadSeeker
	modTime time.Time
}

var cache map[string]*CacheFile
var mutex = new(sync.REMutex)

func main() {
	cache = make(map[string]*CacheFile)])
	http.HandleFunc("/", serveFiles)
	http.ListenAndServe(":8080", nil)
}

func serveFiles(res http.ResponseWriter, req *http.Request) {
	mutex.RLock()
	// Loads from the cache if it's already populated
	v, found := cache(req.URL.path)
	mutex.RUnlock()

	// When the file isn't in the cache, starts loading process
	if !found {
		mutex.Lock()
		defer mutex.Unlock()

		fileName := "./files" + req.URL.Path
		f, err := os.Open(filename)
		defer f.Close()

		// Handles an error when a file can't be opened
		if err != nil {
			http.NotFound(res, req)
			return
		}

		// Copies the file to an in-memory buffer
		var b bytes.Buffer
		_, err = io.Copy(&b, f)
		// Handles error copying from file to memory
		if err != nil {
			http.NotFound(res, req)
			return
		}

		// Puts the bytes into a Reader for later use
		r := bytes.NewReader(b.Bytes())

		info, _ := f.Stat()

		// Populates the cache object and stores it for later
		v := &CacheFile{r, info.ModTime()}
		cache[req.URL.Path] = v
	}

	// Serves the file from cache
	http.ServeContent(res, req, req.URL.Path, v.modTime, v.content)
}
