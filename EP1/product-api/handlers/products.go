package handlers

import (
	"context"
	"encoding/json"
	"gRPC/EP1/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type KeyProduct struct{}

type Products struct{
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products{
	return &Products{l}
}

// PUT Request
func (p Products) UpdateProduct(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handled PUT Requests")

	//It stores the extracted variable into a collection("Vars()")
	//and the data can be extracted in the form of MAP
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert ID", http.StatusBadRequest)
		return
	}

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Printf("New details Prod: %#v", prod)
	
	err = data.UpdateProduct(id, &prod)
	if err != nil {
		http.Error(rw,"Cannot Update the Product informaition",http.StatusInternalServerError)
	}
	rw.Write([]byte("Succesfully updated Product in the Database"))
}

// POST Request
func (p Products) AddProducts(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handled POST Requests")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Printf("Prod: %#v", prod)
	data.AddProducts(&prod)
	rw.Write([]byte("Succesfully added Product in the Database"))
}

// GET Request
func (p Products) GetProducts(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handled Get Requests")

	lp := data.GetProducts()	
	info, err := json.Marshal(lp)
	if err != nil{
		http.Error(rw, "unable to marshal the product metadata", http.StatusInternalServerError )
	}

	rw.Write(info)
}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler{

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request){
		
		prod := data.Product{}
		err := prod.FromJSON(r.Body)

		if err != nil {
			http.Error(rw, "Unable to unmarshal the JSON", http.StatusBadRequest) 
			return
		}

		// Validate the product input 
		err = prod.Validate()
		if err != nil{
			p.l.Println(" [ERROR] validating Product", err)
			http.Error(rw, "Product fields are not correct", http.StatusBadRequest)
			return 
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})	 
}