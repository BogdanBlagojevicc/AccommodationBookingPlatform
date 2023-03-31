package handler

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"user-service/model"
	"user-service/repository"
)

type KeyProduct struct{}
type UserHandler struct {
	Logger *log.Logger
	Repo   *repository.UserRepository
}

func NewUserHandler(l *log.Logger, r *repository.UserRepository) *UserHandler {
	return &UserHandler{l, r}
}

func (u *UserHandler) DatabaseName(ctx context.Context) {
	dbs, err := u.Repo.Cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(dbs)
}

//func (u *UserHandler) GetUserById(rw http.ResponseWriter, h *http.Request) {
//	vars := mux.Vars(h)
//	id := vars["id"]
//
//	patient, err := u.Repo.GetById(id)
//	if err != nil {
//		u.Logger.Print("Database exception: ", err)
//	}
//
//	if patient == nil {
//		http.Error(rw, "Patient with given id not found", http.StatusNotFound)
//		u.Logger.Printf("Patient with id: '%s' not found", id)
//		return
//	}
//
//	err = patient.ToJSON(rw)
//	if err != nil {
//		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
//		p.logger.Fatal("Unable to convert to json :", err)
//		return
//	}
//}

func (u *UserHandler) MiddlewareUserDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		user := &model.User{}
		err := user.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.Logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, user)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (u *UserHandler) PostUser(rw http.ResponseWriter, h *http.Request) {
	user := h.Context().Value(KeyProduct{}).(*model.User)
	//newUser := model.User{ID: primitive.NewObjectID(), FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, Password: user.Password}
	user.ID = primitive.NewObjectID()
	u.Repo.Insert(user)
	rw.WriteHeader(http.StatusCreated)
}

func (u *UserHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		u.Logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
