package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	validator "github.com/go-playground/validator/v10"
)

type Product struct{
	ID 			int			`json:"id"`
	Name		string		`json:"name" validate:"required"`
	Description	string		`json:"description"`
	Price		float32		`json:"price" validate:"gt=0"`
	SKU			string		`json:"sku" validate:"required,sku"`
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

func (p *Product) Validate() error{

	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool{

	// sKU is of format qwer-asdf-zxcv
	reg := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := reg.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1{
		return false
	}
	return true
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