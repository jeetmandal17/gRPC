package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


func main() {


	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request){
		log.Println("Hello world!")

		d, err := ioutil.ReadAll(r.Body)
		if err != nil || d == nil {
			http.Error(rw, "Wrong name Passed", http.StatusBadRequest) 
			return
		}

		for i:=0; i<5; i++{
			fmt.Fprintf(rw, "Hello %s", d)
		}
		/*
		//This snippet is used to get data from user request.
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Wrong name Passed", http.StatusBadRequest) 
		}

		log.Printf("Data %s \n", d)
		*/
	})

	http.HandleFunc("/help", func(rw http.ResponseWriter, r*http.Request){
		log.Println("Hello Help")

		d, err := ioutil.ReadAll(r.Body)
		if err != nil || d == nil {
			http.Error(rw, "Help Menu Coming Soon", http.StatusContinue)
			return
		}

		for i:=0; i<5; i++{
			fmt.Fprintf(rw, "COMING SOON ! \n")
		}
	})

	http.ListenAndServe(":9090", nil) 

}