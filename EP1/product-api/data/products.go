package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct{
	ID 			int			`json:"id"`
	Name		string		`json:"name"`
	Description	string		`json:"description"`
	Price		float32		`json:"price"`
	SKU			string		`json:"sku"`
	CreatedON	string		`json:"-"`
	UpdatedOn	string		`json:"-"`
	DeletedOn	string		`json:"- "`
}

func UpdateProduct(id int, p *Product) error{
	idx, err := findProduct(id)

	if err != nil {
		return err
	}

	p.ID = id
	productList[idx] = p
	return nil
}

var ErrItemNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (int, error){

	for index,item := range productList{
		if item.ID == id{
			return index, nil
		}
	}

	return -1, ErrItemNotFound
}

func GetProducts() []*Product{
	return productList
}

func AddProducts(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int{
	return productList[len(productList)-1].ID + 1
}

func (p *Product) FromJSON(r io.Reader)  error{
	e := json.NewDecoder(r)
	return e.Decode(p) 
}

var productList = []*Product{
	{
		ID:				1,
		Name:			"Latte",
		Description:	"Forthy milky coffee",
		Price: 			2.45,
		SKU: 			"abc323",
		CreatedON: 		time.Now().UTC().String(),
		UpdatedOn: 		time.Now().UTC().String(),
	},
	{
		ID:				2,
		Name:			"Espresso",
		Description:	"Short and string coffee without milk",
		Price: 			1.99,
		SKU: 			"fjd34",
		CreatedON: 		time.Now().UTC().String(),
		UpdatedOn: 		time.Now().UTC().String(),
	},
}