package main

import (
	"dewetour/database"
	"dewetour/pkg/mysql"
	"dewetour/routes"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)
// Test
func main() {

	// Database
	mysql.DatabaseInit()

	// Routing
	database.RunMigration()
	r := mux.NewRouter()
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	// Untuk mengakses file upload
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads", http.FileServer(http.Dir("./uploads"))))

	// Set untuk Header, Method dan Origin
	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	var port = os.Getenv("PORT")

	fmt.Println("server running localhost:" + port)
	http.ListenAndServe(":"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
}
