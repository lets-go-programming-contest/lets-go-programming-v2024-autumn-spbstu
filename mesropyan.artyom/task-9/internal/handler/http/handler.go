package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type service interface {
	Get(name string) ([]byte, error)
	Update(name string, number string) error
	Delete(name string) error
	Upload(name string, number string) error
}

type handler struct {
	service service
}

func NewHandler(s service, r *mux.Router) *mux.Router {
	h := handler{
		service: s,
	}
	r.HandleFunc("/upload", h.upload).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/delete", h.delete).Methods(http.MethodOptions, http.MethodDelete)
	r.HandleFunc("/update", h.update).Methods(http.MethodOptions, http.MethodPatch)
	r.HandleFunc("/get", h.get).Methods(http.MethodOptions, http.MethodGet)

	return r
}

func (h *handler) upload(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	number := r.FormValue("number")

	if err := h.service.Upload(name, number); err != nil {
		http.Error(w, fmt.Errorf("error while uploading contact: %w", err).Error(), http.StatusInternalServerError)
		return
	}
}

func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	if err := h.service.Delete(name); err != nil {
		http.Error(w, fmt.Errorf("contact does not exist: %w", err).Error(), http.StatusNotFound)
	}
}

func (h *handler) get(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	data, err := h.service.Get(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	w.Write(data)
}

func (h *handler) update(w http.ResponseWriter, r *http.Request) {

}
