package user

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func GetUserHandler(service *UserService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := r.PathValue("id")
		id, err := strconv.Atoi(data)
		if err != nil {
			http.Error(w, "id is not valid", http.StatusBadRequest)
			return
		}

		user, err := service.Get(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if user == nil {
			http.NotFound(w, r)
			return
		}

		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			http.Error(w, "error encode json", http.StatusInternalServerError)
		}
	})
}

func CreateUserHandler(service *UserService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userData, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, "error read body", http.StatusBadRequest)
			return
		}

		var userReq userRequest

		err = json.Unmarshal(userData, &userReq)
		if err != nil {
			http.Error(w, "error decoding json :"+err.Error(), http.StatusInternalServerError)
			return
		}

		id, err := service.Create(r.Context(), &userReq)
		if err != nil {
			http.Error(w, "error creating user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(strconv.Itoa(id)))
	})
}

func UpdateUserHandler(service *UserService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func DeleteUserHandler(service *UserService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
