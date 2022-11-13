package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"online-cinema/dto"
	"online-cinema/dto/auth"
	"online-cinema/models"
	bcryptpkg "online-cinema/pkg/bcrypt"
	jwtauth "online-cinema/pkg/jwt"
	"online-cinema/repositories"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type handleAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handleAuth {
	return &handleAuth{AuthRepository}
}

func (h *handleAuth) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	request := new(auth.AuthRequest)

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrResult{Status: "failed register", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	//	check if email already exist
	userExist, emailErr := h.AuthRepository.Login(request.Email)
	if emailErr == nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrResult{Status: "Failed", Message: "Email" + userExist.Email + "already exist!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	//	hashed password
	hashedpass, _ := bcryptpkg.HashPassword(request.Password)
	authField := models.User{
		Email:    request.Email,
		FullName: request.FullName,
		Phone: request.Phone,
		Password: hashedpass,
		Role:     "user",
	}

	user, err := h.AuthRepository.Register(authField)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: convertResponse(user)}
	json.NewEncoder(w).Encode(response)
}

func (h *handleAuth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	request := new(auth.AuthRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	fieldLogin := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	user, err := h.AuthRepository.Login(fieldLogin.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrResult{Status: "Failed", Message: "Email not registered!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	//	cek password
	if isPasswordMatch := bcryptpkg.CheckPasswordHash(fieldLogin.Password, user.Password); !isPasswordMatch {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrResult{Status: "Failed", Message: "wrong password!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	//	create jwt token
	generateToken := jwt.MapClaims{}

	generateToken["id"] = user.ID
	generateToken["exp"] = time.Now().Add(time.Hour * 3).Unix()

	token, err := jwtauth.CreateToken(&generateToken)
	if err != nil {
		log.Println(err)
		fmt.Println("Unauthorize")
		return
	}

	AuthResponse := auth.AuthResponse{
		ID:       user.ID,
		Email:    user.Email,
		FullName: user.FullName,
		Role: user.Role,
		Token:    token,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Status: "Success", Data: AuthResponse}
	json.NewEncoder(w).Encode(response)
}

// cek auth

func (h *handleAuth) CheckAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// Check User by Id
	user, err := h.AuthRepository.Getuser(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrResult{Status: "Failed", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	CheckAuthResponse := auth.CheckAuthResponse{
		Id:    user.ID,
		Name:  user.FullName,
		Email: user.Email,
		Role: user.Role,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Status: "Success", Data: CheckAuthResponse}
	json.NewEncoder(w).Encode(response)
}

func convertResponse(u models.User) map[string]models.User {

	return map[string]models.User{
		"user": {
			FullName: u.FullName,
			Password: u.Password,
			Email:    u.Email,
			ID:       u.ID,
		},
	}
}
