package handler

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeHandler struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func (hdl *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request)     {}
func (hdl *EmployeeHandler) GetEmployeeByID(w http.ResponseWriter, r *http.Request)    {}
func (hdl *EmployeeHandler) GetAllEmployee(w http.ResponseWriter, r *http.Request)     {}
func (hdl *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request)     {}
func (hdl *EmployeeHandler) DeleteEmployeeByID(w http.ResponseWriter, r *http.Request) {}
func (hdl *EmployeeHandler) DeleteAllEmployee(w http.ResponseWriter, r *http.Request)  {}
