package handlers

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	products := []Product{
		{
			ID:          1,
			Name:        "Boxing Padwork Plan",
			Description: "4 week boxing padwork training plan",
			Price:       19.99,
		},
		{
			ID:          2,
			Name:        "Fat Loss Meal Guide",
			Description: "Simple meal guide for fat loss",
			Price:       14.99,
		},
	}

	_ = json.NewEncoder(w).Encode(products)
}
