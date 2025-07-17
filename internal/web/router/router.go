package router

import (
	"goProject/internal/service"
	"goProject/internal/web/handler"

	http "net/http"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

// Route represents a single HTTP route configuration
// Path: URL path for the route
// Method: HTTP method (GET, POST, PUT, DELETE)
// Handler: Function to handle the HTTP request

// Initialize sets up the HTTP router with all routes
// Creates a new service instance and handler
// Registers routes for:
// - Root path ("/") - Serves the main HTML template
// - /items - GET: List all items, POST: Create new item
// - /items/{id} - GET: Get item, PUT: Update item, DELETE: Remove item
// Returns configured ServeMux

// registerRoutes takes a ServeMux and slice of Routes
// Registers each route with its handler function
// Validates that incoming requests use the correct HTTP method
// Returns 405 Method Not Allowed for invalid methods

func Initialize() *http.ServeMux {
	// routing
	mux := http.NewServeMux()
	svc := service.New()
	h := handler.New(svc)
	mux.HandleFunc("/", h.HandleRoot)
	mux.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.HandleItems(w, r)
		case http.MethodPost:
			h.HandleItems(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/items/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.HandleItem(w, r)
		case http.MethodPut:
			h.HandleItem(w, r)
		case http.MethodDelete:
			h.HandleItem(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
