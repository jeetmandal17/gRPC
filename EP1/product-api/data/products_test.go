package data

import "testing"

func TestChecksValidation(t *testing.T){
	p := &Product{
		Name: "Jeet",
		Price: 10,
		SKU: "abcd-efgh-ijkl",
	}

	err := p.Validate()

	if err != nil{
		t.Fatal(err)
	}
}