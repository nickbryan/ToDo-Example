package rest_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/nickbryan/todo/pkg/transport/http/rest"
)

func TestCreate(t *testing.T) {
	repo := emptyMockRepo()

	srv := rest.NewServer()
	srv.RegisterRoutesFor(&rest.TodoHandler{
		TodoRepository: repo,
	})

	body := strings.NewReader(`{
	"content": "My new todo item2",
	"done": false
}`)
	req := httptest.NewRequest(http.MethodPost, "/todos", body)
	w := httptest.NewRecorder()

	srv.ServeHTTP(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("incorrect status code: got %d, want: %d", resp.StatusCode, http.StatusCreated)
	}

	if len(repo.items) != 1 {
		t.Errorf("no items have been created")
	}
}
