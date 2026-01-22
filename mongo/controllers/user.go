package controllers

import (
	"net/http"

	"encoding/json"

	"github.com/ishansaini194/Projects/mongo/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if id == "" || !bson.IsObjectIdHex(id) {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	oid := bson.ObjectIdHex(id)

	session := uc.session.Copy()
	defer session.Close()

	u := models.User{}

	err := session.DB("mongo").C("users").FindId(oid).One(&u)

	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := models.UserResponse{
		Id:     u.Id.Hex(),
		Name:   u.Name,
		Gender: u.Gender,
		Age:    u.Age,
	}

	json.NewEncoder(w).Encode(resp)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()

	var u models.User

	// Decode request
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// Basic validation
	if u.Name == "" || u.Gender == "" || u.Age <= 0 {
		http.Error(w, "Missing or invalid fields", http.StatusBadRequest)
		return
	}

	// Create Mongo ID
	u.Id = bson.NewObjectId()

	// Copy session (CRITICAL)
	session := uc.session.Copy()
	defer session.Close()

	// Insert into DB
	if err := session.DB("mongo").C("users").Insert(&u); err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Prepare response
	resp := models.UserResponse{
		Id:     u.Id.Hex(),
		Name:   u.Name,
		Gender: u.Gender,
		Age:    u.Age,
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Encoding error", http.StatusInternalServerError)
	}
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if id == "" || !bson.IsObjectIdHex(id) {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	oid := bson.ObjectIdHex(id)

	// CRITICAL: copy session
	session := uc.session.Copy()
	defer session.Close()

	err := session.DB("mongo").
		C("users").
		RemoveId(oid)

	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "user deleted successfully",
		"id":      oid.Hex(),
	})
}
