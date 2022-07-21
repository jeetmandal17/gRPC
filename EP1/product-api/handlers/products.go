package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"gRPC/EP1/data"
)


type Products struct{
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products{
	return &Products{l}
}


//This is the standard procedure to direct REST requests
func  (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		p.getProducts(rw,r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request){
	lp := data.GetProducts()	
	info, err := json.Marshal(lp)
	if err != nil{
		http.Error(rw, "unable to marshal the product metadata", http.StatusInternalServerError )
	}

	rw.Write(info)
}