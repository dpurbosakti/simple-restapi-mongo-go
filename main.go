package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"restapi-mongo/handler"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load .env file", err)
	}

	log.Println(".env file loaded")

	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatal("connection error", err)
	}

	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed", err)
	}

	log.Println("mongo connected")
}

func main() {
	defer mongoClient.Disconnect(context.Background())

	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_COLLECTION"))

	empHandler := handler.EmployeeHandler{MongoCollection: coll}
	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler).Methods(http.MethodGet)
	r.HandleFunc("/employee", empHandler.CreateEmployee).Methods(http.MethodPost)
	r.HandleFunc("/employee/{id}", empHandler.GetEmployeeByID).Methods(http.MethodGet)
	r.HandleFunc("/employee", empHandler.GetAllEmployee).Methods(http.MethodGet)
	r.HandleFunc("/employee/{id}", empHandler.UpdateEmployee).Methods(http.MethodPut)
	r.HandleFunc("/employee/{id}", empHandler.DeleteEmployeeByID).Methods(http.MethodDelete)
	r.HandleFunc("/employee", empHandler.DeleteAllEmployee).Methods(http.MethodDelete)

	log.Println("server is running on 4444")
	http.ListenAndServe(":4444", r)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("running..."))
}
