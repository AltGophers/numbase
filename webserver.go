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
	digits := r.FormValue("digit")
	base := r.FormValue("base")
	convertBase := r.FormValue("conversion-base")

	currentBase, err := strconv.Atoi(base)
	if err != nil {
		fmt.Fprintf(w, "Failed to convert base: %v\n", err)
		return
	}

	desiredBase, err := strconv.Atoi(convertBase)
	if err != nil {
		fmt.Fprintf(w, "Failed to convert conversion-base: %v\n", err)
		return
	}

	if currentBase < 2 || currentBase > 16 || desiredBase < 2 || desiredBase > 16 {
		fmt.Fprintf(w, "invalid input: This calculator supports only bases between 2 - 16")
		return
	}

	result, err := convertToBase(int8(currentBase), digits, int8(desiredBase))
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprintf(w, "%v base %v = %v base %v\n", digits, currentBase, result, desiredBase)
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
