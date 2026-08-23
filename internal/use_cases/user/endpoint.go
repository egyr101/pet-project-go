package user

import (
	"encoding/json"
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

func CreateUserHandler(service *UserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func UpdateUserHandler(service *UserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func DeleteUserHandler(service *UserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
