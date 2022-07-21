package data

import "time"

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

func GetProducts() []*Product{
	return productList
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
		ID:				1,
		Name:			"Espresso",
		Description:	"Short and string coffee without milk",
		Price: 			1.99,
		SKU: 			"fjd34",
		CreatedON: 		time.Now().UTC().String(),
		UpdatedOn: 		time.Now().UTC().String(),
	},
}