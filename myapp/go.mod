module myapp

go 1.19

replace github.com/Moji00f/celeritas => ../celeritas

require (
	github.com/Moji00f/celeritas v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi/v5 v5.0.12
)

require github.com/joho/godotenv v1.5.1 // indirect
