package front

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Front single page handler
type handler struct {
    l           *log.Logger
    staticPath  string
    indexPath   string
}

// Create single page handler
func NewHandler (l* log.Logger, staticPath string, indexPath string) *handler {
    return &handler{l, staticPath, indexPath}
}

// Handle serving page
func (fh handler) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
    fh.l.Println("Requested file", rq.URL)

    path, err := filepath.Abs(rq.URL.Path)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusBadRequest)
        return
    }

    // Look for file, if not found serve index
    path = filepath.Join(fh.staticPath, path)
    _, err = os.Stat(path)
    if os.IsNotExist(err) {
        http.ServeFile(rw, rq, filepath.Join(fh.staticPath, fh.indexPath))
        return
    } else if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }

    // Alternatively serve requested file
    http.FileServer(http.Dir(fh.staticPath)).ServeHTTP(rw, rq)
}
