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

func  (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	lp := data.GetProducts()	
	info, err := json.Marshal(lp)
	if err != nil{
		http.Error(rw, "unable to marshal the product metadata", http.StatusInternalServerError )
	}

	rw.Write(info)
}