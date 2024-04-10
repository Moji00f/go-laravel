package main

import "github.com/Moji00f/celeritas"

type application struct {
	App *celeritas.Celeritas
}

func main() {

	c := initApplication()
	c.App.ListenAndServer()
}
