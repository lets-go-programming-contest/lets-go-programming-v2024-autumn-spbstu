package http

import (
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
		code := getMappedStatusCode(errorsUploadMap, err)
		http.Error(w, err.Error(), code)
		return
	}
}

func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if err := h.service.Delete(name); err != nil {
		code := getMappedStatusCode(errorsDeleteMap, err)
		http.Error(w, err.Error(), code)
	}
}

func (h *handler) get(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	data, err := h.service.Get(name)
	if err != nil {
		code := getMappedStatusCode(errorsGetMap, err)
		http.Error(w, err.Error(), code)
		return
	}
	w.Write(data)
}

func (h *handler) update(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	number := r.FormValue("number")

	if err := h.service.Update(name, number); err != nil {
		code := getMappedStatusCode(errorsUpdateMap, err)
		http.Error(w, err.Error(), code)
		return
	}
}
