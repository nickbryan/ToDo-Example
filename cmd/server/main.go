package main

import (
	"fmt"
	"os"

	"github.com/nickbryan/todo/pkg/datastore/inmemory"

	"github.com/nickbryan/todo/pkg/transport/http/rest"
)

func main() {
	s := rest.NewServer()

	s.RegisterRoutesFor(&rest.TodoHandler{
		TodoRepository: inmemory.NewTodoRepository(),
	})

	if err := s.Start(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
