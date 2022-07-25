package handlers

import (
	"encoding/json"
	"gRPC/EP1/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)


type Products struct{
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products{
	return &Products{l}
}


//This is the standard procedure to direct REST requests
func  (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	
	/* GET request to the server */
	if r.Method == http.MethodGet{
		p.l.Println("GET")

		p.getProducts(rw,r)
		return
	}

	/* POST request to the server */
	if r.Method == http.MethodPost{
		p.l.Println("POST")
		p.addProducts(rw,r)
		return
	}

	/* PUT requests to the server */
	if r.Method == http.MethodPut {

		p.l.Println("PUT", r.URL.Path )
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path,-1)

		if len(g) != 1 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "String conversion failed", http.StatusFailedDependency)
			return
		}

		p.l.Println("Got id", id) 

		p.updateProduct(id, rw, r)
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handled PUT Requests")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal the JSON", http.StatusBadRequest) 
	}

	p.l.Printf("New details Prod: %#v", prod)
	
	err = data.UpdateProduct(id, prod)
	if err != nil {
		http.Error(rw,"Cannot Update the Product informaition",http.StatusInternalServerError)
	}
	rw.Write([]byte("Succesfully updated Product in the Database"))
}

func (p *Products) addProducts(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handled POST Requests")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal the JSON", http.StatusBadRequest) 
	}

	p.l.Printf("Prod: %#v", prod)
	data.AddProducts(prod)
	rw.Write([]byte("Succesfully added Product in the Database"))
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handled Get Requests")

	lp := data.GetProducts()	
	info, err := json.Marshal(lp)
	if err != nil{
		http.Error(rw, "unable to marshal the product metadata", http.StatusInternalServerError )
	}

	rw.Write(info)
}