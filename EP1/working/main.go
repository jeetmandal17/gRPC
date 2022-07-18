package main

import (
	"gRPC/EP1/handlers"
	"log"
	"net/http"
	"os"
)


func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	//Assign the new hello handler
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodBye(l)
	ch := handlers.NewCalculator(l)

	sm := http.NewServeMux()

	sm.Handle("/", hh)
	sm.Handle("/g", gh)
	sm.Handle("/c", ch)
	 
	http.ListenAndServe(":9090", sm)  

}