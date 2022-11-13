package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"online-cinema/dto"
	"online-cinema/dto/categories"
	"online-cinema/models"
	"online-cinema/repositories"
	"strconv"
)

type handlercategory struct {
	CategoryRepository repositories.CategoryRepository
}

func HandlerCategory(categoryRepository repositories.CategoryRepository) *handlercategory {
	return &handlercategory{categoryRepository}
}

//get users

func (h *handlercategory) GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	catgeory, err := h.CategoryRepository.GetCategory()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: catgeory}
	json.NewEncoder(w).Encode(response)
}

func (h *handlercategory) GetCategoryId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	category, err := h.CategoryRepository.GetCategoriID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response := dto.ErrResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: category}
	json.NewEncoder(w).Encode(response)
}

func (h *handlercategory) CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	request := new(categories.CategoryRequest)

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrResult{Status: "failed register", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	Field := models.Categorie{
		Name: request.Name,
	}

	category, err := h.CategoryRepository.CreateCategory(Field)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: category}
	json.NewEncoder(w).Encode(response)
}

func (h *handlercategory) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var request categories.CategoryRequest
	request = categories.CategoryRequest{
		Name: r.FormValue("name"),
	}
	categoryModel := models.Categorie{}
	if request.Name != "" {
		categoryModel.Name = request.Name
	}

	category, err := h.CategoryRepository.UpdateCategorie(categoryModel, id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Succses", Data: category}
	json.NewEncoder(w).Encode(response)
}

func (h *handlercategory) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	category := models.Categorie{}

	deletedCategory, err := h.CategoryRepository.DeleteCategory(category, id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrResult{Status: "Failed delete", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: deletedCategory}
	json.NewEncoder(w).Encode(response)
}
