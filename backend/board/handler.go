package board

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Board handler
type handler struct {
	l  *log.Logger
	db *gorm.DB
}

// Create board handler
func NewHandler(l *log.Logger, db *gorm.DB) *handler {
	l.Println("Auto migrating database")
	db.AutoMigrate(&board{}, &taskList{}, &task{})
	return &handler{l, db}
}

// Register board routes
func (bh handler) RegisterRoutes(route *mux.Router) {
	route.Use(bh.middleware)

	// Board methods
	route.HandleFunc("", bh.CreateBoard).Methods("POST")
	route.HandleFunc("/list", bh.GetBoards).Methods("GET")
	boardRoute := route.PathPrefix("/{boardID:[0-9]+}").Subrouter()

	boardRoute.HandleFunc("", bh.GetBoard).Methods("GET")
	boardRoute.HandleFunc("", bh.DeleteBoard).Methods("DELETE")
	boardRoute.HandleFunc("", bh.CreateList).Methods("POST")

	// List methods
	listRoute := boardRoute.PathPrefix("/{listID:[0-9]+}").Subrouter()
	listRoute.HandleFunc("", bh.UpdateList).Methods("PATCH")
	listRoute.HandleFunc("", bh.DeleteList).Methods("DELETE")

	// Task methods
	listRoute.HandleFunc("", bh.PostTask).Methods("POST")
	listRoute.HandleFunc("/{taskID:[0-9]+}", bh.UpdateTask).Methods("PUT")
	listRoute.HandleFunc("/{taskID:[0-9]+}", bh.DeleteTask).Methods("DELETE")
}

// Middleware func
func (bh handler) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, rq *http.Request) {
		// Log request
		bh.l.Println(rq.Method, "on", rq.URL)

		// Set header
		rw.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(rw, rq)
	})
}
