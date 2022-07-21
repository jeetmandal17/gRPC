package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)


type Calculator struct{
	l *log.Logger
}

func NewCalculator(l *log.Logger) *Calculator{
	return &Calculator{l}
}

func (c *Calculator) ServeHTTP(rw http.ResponseWriter, r *http.Request){

	c.l.Println("I am the Calculator!")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	n , err:= strconv.Atoi(string(data))

	switch n{
		case 1:
			fmt.Fprintf(rw, "Addtion -> %v", n)
		case 2:
			fmt.Fprintf(rw, "Substraction -> %v", n)
		case 3:
			fmt.Fprintf(rw, "Multiplication -> %v", n)
		case 4:
			fmt.Fprintf(rw, "Division -> %v", n)
	}
}