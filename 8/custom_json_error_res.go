package main

type Error struct {
	HTTPCode int    `json:"-"`
	Code     int    `json:"code,omitempty"`
	Message  string `json:"message"`
}

func JSONError(w http.ResponseWriter, e Error) {
	// Wraps Error struct in anonymous struct with error property
	data := struct {
		Err Error `json:"error"`
	}{e}

	// Converts error data to JSON and handles an error if one exists
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Sets the response MIME type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Makes sure the HTTP status code is properly set for the error
	w.WriteHeader(e.HTTPCode)

	// Writes the JSON body as output
	fmt.Fprint(w, string(b))

	func displayError(w http.ResponseWriter, r *http.Request) {
		// Creates an instance of Error to use for the response error
		e := Error{
			HTTPCode: http.StatusForbidden,
			Code : 123,
			Message: "An Error Ocurred"
		}

		JSONError(w, e)
	}
}

func main() {
	http.HandleFunc("/", displayError)
	http.ListenAndServe(":8080", nil)
}
