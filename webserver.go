package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// convertToBaseHandler handles the "POST /convert-to-base" endpoint and
// converts the provided value to the base specified.
func convertToBaseHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v\n", err)
		return
	}

	// Getting the form values from base.html using their input name.
	digit := r.FormValue("digit")
	base := r.FormValue("base")
	convertBase := r.FormValue("conversion-base")

	// Converting the form values to  integers.
	d, err := strconv.Atoi(digit)
	if err != nil {
		fmt.Fprintf(w, "Failed to convert digit: %v\n", err)
		return
	}

	b, err := strconv.Atoi(base)
	if err != nil {
		fmt.Fprintf(w, "Failed to convert base: %v\n", err)
		return
	}

	cb, err := strconv.Atoi(convertBase)
	if err != nil {
		fmt.Fprintf(w, "Failed to convert conversion-base: %v\n", err)
		return
	}

	// Base calculation is done here.
	result, _ := convertToBase(int8(b), int64(d), cb)
	fmt.Fprintf(w, "Result = %v base %v\n", result, cb)
}

// startServer starts the default http server.
func startServer() {
	// Routing to index.html in the static file. Only the ./static directory was
	// specified but it wil automatically route to index.html as long it's in
	// the directory
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// Routing to /base e.g(localhost:8080/convert-to-base) which is handled by
	http.HandleFunc("/convert-to-base", convertToBaseHandler)

	// Listening and serving the html file on port 8080
	fmt.Println("Starting server on port 8080. Visit Base Converter on your browser at http://localhost:8080.")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
