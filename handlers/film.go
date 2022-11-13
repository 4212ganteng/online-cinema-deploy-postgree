package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"online-cinema/dto"
	"online-cinema/dto/film"
	"online-cinema/models"
	"online-cinema/repositories"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gorilla/mux"
)

// Cloudinary
// Declare Context Background, Cloud Name, API Key, API Secret here ...
var ctx = context.Background()
var CLOUD_NAME = os.Getenv("CLOUD_NAME")
var API_KEY = os.Getenv("API_KEY")
var API_SECRET = os.Getenv("API_SECRET")

type handlerFilm struct {
	FilmRepository repositories.FilmRepository
}

func HandlerFilm(filmRepository repositories.FilmRepository) *handlerFilm {
	return &handlerFilm{filmRepository}
}

//get film

func (h *handlerFilm) GetFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	film, err := h.FilmRepository.GetFilm()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// add manipulation path file on this below code ...
	for i, p := range film {
		imagePath := os.Getenv("PATH_FILE") + p.Image
		film[i].Image = imagePath
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: film}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFilm) GetFilmId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	film, err := h.FilmRepository.GetfilmID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response := dto.ErrResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Add manipulation path file on this below code ...
	film.Image = os.Getenv("PATH_FILE") + film.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: film}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFilm) CreateFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get image filepath after cloudinary
	dataContex := r.Context().Value("dataFile")
	filepath := dataContex.(string)

	price, _ := strconv.Atoi(r.FormValue("price"))
	category, err := strconv.Atoi(r.FormValue("category_id"))

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "dumbmerch_file_b40"})

	if err != nil {
		fmt.Println(err.Error())
	}

	// Upload file to Cloudinary here ...

	Field := models.Film{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Price:       price,
		FilmUrl:     r.FormValue("filmUrl"),
		// Image:    filename, // Modify store file URL to database from resp.SecureURL here ...
		Image:      resp.SecureURL,
		CategoryID: category,
	}

	film, err := h.FilmRepository.CreateFilm(Field)

	fmt.Println("ini data film", film)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: film}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFilm) UpdateFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataUpload := r.Context().Value("dataFile")
	filename := dataUpload.(string)

	price, _ := strconv.Atoi(r.FormValue("price"))
	category, _ := strconv.Atoi(r.FormValue("category_id"))

	request := film.CreateFilmRequest{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Price:       price,
		FilmUrl:     r.FormValue("filmUrl"),
		Image:       os.Getenv("PATH_FILE") + filename,
		CategoryID:  category,
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	film := models.Film{}

	film.ID = id

	if request.Title != "" {
		film.Title = request.Title
	}

	if request.Price != 0 {
		film.Price = request.Price
	}

	if filename != "" {
		film.Image = request.Image
	}

	if request.Description != "" {
		film.Description = request.Description
	}
	if request.FilmUrl != "" {
		film.FilmUrl = request.FilmUrl
	}
	if request.CategoryID != 0 {
		film.CategoryID = request.CategoryID
	}

	data, err := h.FilmRepository.UpdateFilm(film, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFilm) DeleteFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	film := models.Film{}

	deletedFilm, err := h.FilmRepository.DeleteFilm(film, id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrResult{Status: "Failed delete", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: deletedFilm}
	json.NewEncoder(w).Encode(response)
}
