package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"gofr.dev/pkg/gofr"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/rider/handler"
	"github.com/rider/migrations"
	"github.com/rider/service"
	"github.com/rider/store"
)

func main() {
	// Load environment variables from configs/.env file
	envPath := filepath.Join("configs", ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file from %s", envPath)
	}

	// Get database connection details from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// connection string
	connStr := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " host=" + dbHost + " port=" + dbPort + " sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	a := gofr.New()
	a.Migrate(migrations.All())

	str := store.New(db)
	svc := service.New(str)
	h := handler.New(svc)

	// Create a new mux router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/rider/{riderid}/availability", h.UpdateRiderAvailability).Methods(http.MethodPut)
	router.HandleFunc("/rider/{riderid}/location", h.UpdateRiderLocation).Methods("PUT")
	router.HandleFunc("/riders/nearby", h.GetNearbyRiders).Methods("GET")
	router.HandleFunc("/rider", h.RegisterRiders).Methods("POST")
	router.HandleFunc("/rider/{id}", h.UpdateRiderDetails).Methods("PUT")
	router.HandleFunc("/rider/{id}", h.GetRiderDetails).Methods("GET")

	log.Println("Service initialized at port:8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
