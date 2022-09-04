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
		digit := r.FormValue("digit")
		base := r.FormValue("base")
		convertBase := r.FormValue("convarsion-base")

		d, _ := strconv.Atoi(digit)
		b, _ := strconv.Atoi(base)
		cb, _ := strconv.Atoi(convertBase)

		result,_ := fromAnyBasetoAnyBase(b, d, cb)
	
	// 	// this is where i do base calculation
		fmt.Fprintf(w, "Answer = %v base %v\n", result, cb)
		
	
	}
func startServer() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/base", ConverterHandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal(err)
	}

	
}
