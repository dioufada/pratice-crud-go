package controllers

import (
	"encoding/json"
	models "example/web-service-gin/Models"
	"example/web-service-gin/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewProduct models.Product

func GetProducts(w http.ResponseWriter, r *http.Request) {
	newProducts := models.GetAllProducts()
	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(newProducts)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	// json.NewEncoder(w).Encode(newProducts)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	productId := vars["productId"]

	ID, err := strconv.ParseInt(productId, 0, 0)

	if err != nil {
		fmt.Println("Erreur de parsement")
	}
	productDetails, _ := models.GetProductById(ID)
	res, _ := json.Marshal((productDetails))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	// json.NewEncoder(w).Encode(productId)
	fmt.Println("GetProduitById")
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	CreateProduct := &models.Product{}
	utils.ParseBody(r, CreateProduct)
	b := CreateProduct.CreateProduct()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	// json.NewEncoder(w).Encode(CreateProduct)
	fmt.Println("CreateProduit")
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]
	ID, err := strconv.ParseInt(productId, 0, 0)
	if err != nil {
		fmt.Println("Erreur de parsement")
	}

	product := models.DeleteProduct(ID)
	res, _ := json.Marshal(product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	json.NewEncoder(w).Encode("Product Deleted Successfully!")
	fmt.Println("DeleteProduit")
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var updateProduct = &models.Product{}
	utils.ParseBody(r, updateProduct)
	vars := mux.Vars(r)
	productId := vars["productId"]
	ID, err := strconv.ParseInt(productId, 0, 0)
	if err != nil {
		fmt.Println("Erreur de parsement")
	}
	productDetails, db := models.GetProductById(ID)

	if updateProduct.Name != "" {
		productDetails.Name = updateProduct.Name
	}

	if updateProduct.Price != 0 {
		productDetails.Price = updateProduct.Price
	}

	if updateProduct.Price != 0 {
		productDetails.Category = updateProduct.Category
	}

	if updateProduct.Quantity != 0 {
		productDetails.Quantity = updateProduct.Quantity
	}

	db.Save(&productDetails)
	res, _ := json.Marshal(productDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	json.NewEncoder(w).Encode(productDetails)
	fmt.Println("UpdateProduit")
}
