package handlers

import (
	"encoding/json"
	"net/http"
	"online-cinema/dto"
	"online-cinema/repositories"
	"strconv"

	"github.com/gorilla/mux"
)

type handler struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handler {
	return &handler{UserRepository}
}

//get users

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.UserRepository.GetUser()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: users}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) GetuserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.UserRepository.GetUserID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response := dto.ErrResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: user}
	json.NewEncoder(w).Encode(response)
}

// func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])
// 	var request auth.AuthRequest
// 	request = auth.AuthRequest{
// 		FullName: r.FormValue("fullName"),
// 		Email:    r.FormValue("email"),
// 	}
// 	userModel := models.User{}
// 	if request.FullName != "" {
// 		userModel.FullName = request.FullName
// 	}

// 	if request.Email != "" {
// 		userModel.Email = request.Email
// 	}
// 	user, err := h.UserRepository.UpdateUser(userModel, id)

// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrResult{Status: "Failed", Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Status: "Succses", Data: user}
// 	json.NewEncoder(w).Encode(response)
// }

// func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])
// 	user := models.User{}

// 	deletedUser, err := h.UserRepository.DeleteUser(user, id)

// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrResult{Status: "Failed delete", Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Status: "Success", Data: deletedUser}
// 	json.NewEncoder(w).Encode(response)
// }
