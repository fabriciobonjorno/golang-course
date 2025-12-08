package handlers

import (
	"apis/internal/dto"
	"apis/internal/entity"
	"apis/internal/infra/database"
	entityPkg "apis/pkg/entity"
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// Implementation for creating a product goes here route
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	// Further processing to save the product using h.ProductDB
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		http.Error(w, "Error creating product entity", http.StatusInternalServerError)
		return
	}
	err = h.ProductDB.Create(p)
	if err != nil {
		http.Error(w, "Error saving product to database", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a product goes here route
	// Assume we get the product ID from the URL
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}
	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a product goes here route
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	err = h.ProductDB.Update(&product)
	if err != nil {
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a product goes here route
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}
	_, err := h.ProductDB.FindByID(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	err = h.ProductDB.Delete(id)
	if err != nil {
		http.Error(w, "Error deleting product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// FindAll products handler can be added similarly if needed
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	// Convert string to int if necessary
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0 // default page
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10 // default limit
	}
	sort := r.URL.Query().Get("sort")
	// For simplicity, with pagination logic here
	products, total, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(limitInt)))
	response := map[string]interface{}{
		"data":        products,
		"page":        pageInt,
		"limit":       limitInt,
		"total_items": total,
		"total_pages": totalPages,
		"has_next":    pageInt < totalPages,
		"has_prev":    pageInt > 1,
		"next_page": func() int {
			if pageInt < totalPages {
				return pageInt + 1
			}
			return 0
		}(),
		"prev_page": func() int {
			if pageInt > 1 {
				return pageInt - 1
			}
			return 0
		}(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
