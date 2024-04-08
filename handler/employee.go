package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"restapi-mongo/model"
	"restapi-mongo/repository"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeHandler struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func (hdl *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid body", err)
		res.Error = err.Error()
		return
	}

	emp.ID = uuid.NewString()

	repo := repository.EmployeeRepo{MongoCollection: hdl.MongoCollection}

	insertID, err := repo.InsertEmployee(&emp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("insert error")
		res.Error = err.Error()
		return
	}

	res.Data = emp.ID
	w.WriteHeader(http.StatusOK)

	log.Println("employee inserted with id ", insertID, emp)
}
func (hdl *EmployeeHandler) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	empID := mux.Vars(r)["id"]
	log.Println("employee id ", empID)

	repo := repository.EmployeeRepo{MongoCollection: hdl.MongoCollection}
	emp, err := repo.FindEmployeeByID(empID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error:", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp
	w.WriteHeader(http.StatusOK)
}
func (hdl *EmployeeHandler) GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: hdl.MongoCollection}

	emp, err := repo.FindAllEmployees()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error:", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp
	w.WriteHeader(http.StatusOK)
}
func (hdl *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	empID := mux.Vars(r)["id"]
	log.Println("employee id ", empID)

	if empID == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid employee id")
		res.Error = "invalid employee id"
		return
	}

	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid body", err)
		res.Error = err.Error()
		return
	}

	emp.ID = empID

	repo := repository.EmployeeRepo{MongoCollection: hdl.MongoCollection}

	count, err := repo.UpdateEmployeeByID(empID, &emp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error: ", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)
}
func (hdl *EmployeeHandler) DeleteEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	empID := mux.Vars(r)["id"]
	log.Println("employee id ", empID)

	if empID == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid employee id")
		res.Error = "invalid employee id"
		return
	}

	repo := repository.EmployeeRepo{MongoCollection: hdl.MongoCollection}

	count, err := repo.DeleteEmployeeByID(empID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error: ", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)
}
func (hdl *EmployeeHandler) DeleteAllEmployee(w http.ResponseWriter, r *http.Request) {}
