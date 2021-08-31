package board

import (
	"io/ioutil"
	"net/http"
)

// (POST) Create new task
func (bh handler) PostTask(rw http.ResponseWriter, rq *http.Request) {

}

// (PUT) update task
func (bh handler) UpdateTask(rw http.ResponseWriter, rq *http.Request) {
    body, _ := ioutil.ReadAll(rq.Body)
    rw.Write(body)
}

// (DELETE) task
func (bh handler) DeleteTask(rw http.ResponseWriter, rq *http.Request) {

}
