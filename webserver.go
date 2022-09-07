package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	)

	func ConverterHandler(w http.ResponseWriter, r *http.Request){
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w,"ParseForm() err: %v\n", err)
			return
		}
		fmt.Fprintf(w,"Post request successful\n")

		// getting the form values from base.html using their input name 
		digit := r.FormValue("digit")
		base := r.FormValue("base")
		convertBase := r.FormValue("convarsion-base")

		// converting the form values to  integers
		d, _ := strconv.Atoi(digit)
		b, _ := strconv.Atoi(base)
		cb, _ := strconv.Atoi(convertBase)

	//   base calculation is done here
		result,_:= fromAnyBasetoAnyBase(int8(b), int64(d), cb)

		fmt.Fprintf(w, "Result = %v base %v\n", result, cb)
	
	}

func startServer() {
	// routing to index.html in the static file.
	//note: only the ./static directory was specified but it wil automatically route to index.html as long it's in the directory
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)

  //routing to /base e.g(localhost:8080/base) which is handled by convertHandler to display result
	http.HandleFunc("/base", ConverterHandler)

	fmt.Println("Starting server at port 8080")
	fmt.Println("go to http://localhost:8080 on your browser")

	//listening and serving the html file on port 8080
	if err := http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal(err)
	}

}
