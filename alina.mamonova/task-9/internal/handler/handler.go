package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

type service interface {
	Get(id string) ([]byte, error)
	Update(id string, name string, number string) error
	Delete(id string) error
	Upload(name string, number string) error
	GetAll() ([]byte, error)
}

type handler struct {
	service service
}

func NewHandler(s service, r *mux.Router) *mux.Router {
	h := handler{
		service: s,
	}
	r.HandleFunc("/contacts", h.upload).Methods(http.MethodPost)
	r.HandleFunc("/contacts", h.getAll).Methods(http.MethodGet)
	r.HandleFunc("/contacts/{id}", h.delete).Methods(http.MethodDelete)
	r.HandleFunc("/contacts/{id}", h.update).Methods(http.MethodPut)
	r.HandleFunc("/contacts/{id}", h.get).Methods(http.MethodGet)

	return r
}

func (h *handler) upload(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	number := r.FormValue("number")

	if name == "" || number == "" {
		http.Error(w, "Name and number are required", http.StatusBadRequest)
		return
	}

	if err := h.service.Upload(name, number); err != nil {
		code := getMappedStatusCode(errorsUploadMap, err)
		http.Error(w, err.Error(), code)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.service.Delete(id); err != nil {
		code := getMappedStatusCode(errorsDeleteMap, err)
		http.Error(w, err.Error(), code)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	data, err := h.service.Get(id)
	if err != nil {
		code := getMappedStatusCode(errorsGetMap, err)
		http.Error(w, err.Error(), code)
		return
	}

	w.Write(data)
}

func (h *handler) update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	name := r.FormValue("name")
	number := r.FormValue("number")

	if name == "" || number == "" {
		http.Error(w, "Name and number are required", http.StatusBadRequest)
		return
	}

	if err := h.service.Update(id, name, number); err != nil {
		code := getMappedStatusCode(errorsUpdateMap, err)
		http.Error(w, err.Error(), code)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) getAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
