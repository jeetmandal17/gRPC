package main

import (
	"context"
	"gRPC/EP1/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)


func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	//Assign new http handlers from here
	/* THESE ARE THE TESTING HANDLERS */
	// hh := handlers.NewHello(l)
	// gh := handlers.NewGoodBye(l)
	// ch := handlers.NewCalculator(l)

	ph := handlers.NewProduct(l)
 
	sm := http.NewServeMux()

	sm.Handle("/", ph)
	
	s := http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 20*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}


	go func(){
		err := s.ListenAndServe()  
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Received terminate, going graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}