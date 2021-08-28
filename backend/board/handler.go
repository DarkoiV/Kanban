package board

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Board handler
type handler struct {
    l *log.Logger
    db *gorm.DB
}

// Create board handler
func NewHandler (l *log.Logger, db *gorm.DB) *handler {
    l.Println("Auto migrating database")
    db.AutoMigrate(&board{}, &taskList{}, &task{})
    return &handler{l, db}
}

func (bh handler) RegisterRoutes(route *mux.Router) {
    route.Use(bh.middleware)

    // Board route
    route.HandleFunc("/new", bh.CreateBoard).Methods("POST")
    boardRoute := route.PathPrefix("/{boardID:[0-9]+}").Subrouter()

    boardRoute.HandleFunc("", bh.GetBoard).Methods("GET")
    boardRoute.HandleFunc("", bh.UpdateOrder).Methods("PATCH")
    boardRoute.HandleFunc("", bh.DeleteBoard).Methods("DELETE")

    // List route
    boardRoute.HandleFunc("/new", bh.CreateList).Methods("POST")
    listRoute := boardRoute.PathPrefix("/{listID:[0-9]+}").Subrouter()
    listRoute.HandleFunc("", bh.DeleteList).Methods("DELETE")

    // Task route
    listRoute.HandleFunc("/new", bh.PostTask).Methods("POST")
    listRoute.HandleFunc("/{taskID:[0-9]+}", bh.UpdateTask).Methods("PUT")
    listRoute.HandleFunc("/{taskID:[0-9]+}", bh.DeleteTask).Methods("DELETE")
}

func (bh handler) middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, rq *http.Request) {
        // Log request
        bh.l.Println(rq.Method, "on", rq.URL)

        // Set header
        rw.Header().Set("Content-Type", "application/json")

        next.ServeHTTP(rw, rq)
    })
}
