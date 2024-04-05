package repository

import (
	"context"
	"log"
	"os"
	"restapi-mongo/model"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_URI")

	mongoTestClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))

	if err != nil {
		log.Fatal("error while connecting mongodb", err)
	}

	log.Println("mongodb successfully connected")

	err = mongoTestClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed", err)
	}

	log.Println("ping success")

	return mongoTestClient
}

func TestMongoOperation(t *testing.T) {
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	// dummy data
	emp1 := uuid.New().String()
	// emp2 := uuid.New().String()

	// connect to collection
	coll := mongoTestClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("DB_COLLECTION"))

	empRepo := EmployeeRepo{MongoCollection: coll}

	// insert employee 1 data
	t.Run("Insert employee 1", func(t *testing.T) {
		emp := model.Employee{
			Name:       "test1",
			Department: "department_test1",
			ID:         emp1,
		}

		result, err := empRepo.InsertEmployee(&emp)

		if err != nil {
			t.Fatal(err)
		}

		t.Log("Insert emplyee 1 success", result)

	})

	t.Run("Get employee 1", func(t *testing.T) {
		result, err := empRepo.FindEmployeeByID(emp1)

		if err != nil {
			t.Fatal(err)
		}

		t.Log("emp 1", result)
	})

	t.Run("Get all employees", func(t *testing.T) {
		result, err := empRepo.FindAllEmployees()

		if err != nil {
			t.Fatal(err)
		}

		t.Log("all employees", result)
	})

	t.Run("Update employee 1 name", func(t *testing.T) {
		emp := model.Employee{
			Name:       "testupdate1",
			Department: "testdepartment1",
			ID:         emp1,
		}
		result, err := empRepo.UpdateEmployeeByID(emp1, &emp)
		if err != nil {
			t.Fatal(err)
		}

		t.Log("employee updated", result)
	})

	t.Run("Get employee 1 after update", func(t *testing.T) {
		result, err := empRepo.FindEmployeeByID(emp1)
		if err != nil {
			t.Fatal(err)
		}

		t.Log("emp 1", result.Name)
	})

	t.Run("Delete employee 1", func(t *testing.T) {
		result, err := empRepo.DeleteEmployeeByID(emp1)

		if err != nil {
			t.Fatal(err)
		}

		t.Log("delete count", result)
	})

	t.Run("Get all employees after delete", func(t *testing.T) {
		result, err := empRepo.FindAllEmployees()
		if err != nil {
			t.Fatal(err)
		}

		t.Log("employees", result)
	})

	t.Run("Delete all employees", func(t *testing.T) {
		result, err := empRepo.DeleteAllEmployees()

		if err != nil {
			t.Fatal(err)
		}

		t.Log("delete count", result)
	})
}
