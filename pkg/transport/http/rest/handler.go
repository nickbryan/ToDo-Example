package rest

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"github.com/gorilla/mux"

	"github.com/nickbryan/todo/pkg"
)

type TodoHandler struct {
	TodoRepository pkg.TodoRepository
}

func (h *TodoHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/todos", h.handleList()).Methods(http.MethodGet)
	r.HandleFunc("/todos", h.handleCreate()).Methods(http.MethodPost)
	r.HandleFunc("/todos/{id}", h.handleUpdate()).Methods(http.MethodPut)
}

func (h *TodoHandler) handleList() http.HandlerFunc {
	type item struct {
		ID      string `json:"id"`
		Content string `json:"content"`
		Done    bool   `json:"bool"`
	}

	type response struct {
		List []item `json:"list"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		resp := &response{List: make([]item, 0)}

		for _, i := range h.TodoRepository.List() {
			resp.List = append(resp.List, item{
				ID:      i.ID().String(),
				Content: i.Content(),
				Done:    i.Done(),
			})
		}

		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (h *TodoHandler) handleCreate() http.HandlerFunc {
	type request struct {
		Content string `json:"content"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req request

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		item := pkg.NewTodoItem(req.Content)
		if err := h.TodoRepository.Create(item); err != nil {
			// TODO: error response
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func (h *TodoHandler) handleUpdate() http.HandlerFunc {
	type request struct {
		Content string `json:"content"`
		Done    bool   `json:"done"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req request

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		reqVars := mux.Vars(r)

		id, err := uuid.Parse(reqVars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		item, err := h.TodoRepository.Find(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		item.MarkDone(req.Done)
		item.UpdateContent(req.Content)

		if err = h.TodoRepository.Update(&item); err != nil {
			// TODO: error response
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
