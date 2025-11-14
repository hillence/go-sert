package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Desc  string `json:"desc"`
}

var Products []Product

func LoadProducts() {
	data, err := ioutil.ReadFile("data/products.json")
	if err != nil {
		log.Fatal("Ошибка чтения products.json:", err)
	}

	err = json.Unmarshal(data, &Products)
	if err != nil {
		log.Fatal("Ошибка JSON:", err)
	}
}

func GetProductByID(id string) *Product {
	for _, p := range Products {
		if p.ID == id {
			return &p
		}
	}
	return nil
}
