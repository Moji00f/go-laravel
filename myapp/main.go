package main

import (
	"myapp/handlers"

	"github.com/Moji00f/celeritas"
)

type application struct {
	App      *celeritas.Celeritas
	Handlers *handlers.Handlers
}

func main() {

	c := initApplication()
	c.App.ListenAndServer()
}
