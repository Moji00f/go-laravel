package main

import (
	"fmt"
	"myapp/data"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	// middleware must come before any routes

	// a.App.Routes.Use(a.App.SessionLoad)

	// add routes here
	mux := a.App.Routes
	mux.Group(func(r chi.Router) {
		r.Use(a.App.SessionLoad)
		r.Get("/", a.Handlers.Home)
		r.Get("/go-page", a.Handlers.GoPage)
		r.Get("/jet-page", a.Handlers.JetPage)
		r.Get("/sessions", a.Handlers.SesstionTest)
		r.Get("/users/login", a.Handlers.UserLogin)
		r.Post("/users/login", a.Handlers.PostUserLogin)
		r.Get("/users/logout", a.Handlers.Logout)
	})

	// a.App.Routes.Get("/", a.Handlers.Home)
	// a.App.Routes.Get("/go-page", a.Handlers.GoPage)
	// a.App.Routes.Get("/jet-page", a.Handlers.JetPage)
	// a.App.Routes.Get("/sessions", a.Handlers.SesstionTest)

	// a.App.Routes.Get("/test-database", func(w http.ResponseWriter, r *http.Request) {
	// 	query := "select id, first_name from users where id = 1"
	// 	row := a.App.DB.Pool.QueryRowContext(r.Context(), query)

	// 	var id int
	// 	var name string

	// 	err := row.Scan(&id, &name)
	// 	if err != nil {
	// 		a.App.ErrorLog.Println(err)
	// 		return
	// 	}

	// 	fmt.Fprintf(w, "%d %s", id, name)

	// })

	// a.App.Routes.Get("/create-user", func(w http.ResponseWriter, r *http.Request) {
	// 	u := data.User{
	// 		FirstName: "Mojtaba",
	// 		LastName:  "Ahmadi",
	// 		Email:     "ssss@gmail.com",
	// 		Active:    1,
	// 		Password:  "password",
	// 	}
	// 	id, err := a.Models.Users.Insert(u)
	// 	if err != nil {
	// 		a.App.ErrorLog.Println(err)
	// 		return
	// 	}

	// 	fmt.Fprintf(w, "%d: %s", id, u.FirstName)
	// })
	var userInsertionFlag = false
	mux.Get("/create-user", func(w http.ResponseWriter, r *http.Request) {
		if userInsertionFlag {
			a.App.ErrorLog.Println("User insertion already in progress")
			// Handle the situation accordingly
			return
		}
		u := data.User{
			FirstName: "Yas",
			LastName:  "Fad",
			Email:     "Fad@gmail.com",
			Active:    1,
			Password:  "password",
		}

		// Ensure only one insertion attempt per request
		userInsertionFlag = true
		defer func() {
			userInsertionFlag = false
		}()

		id, err := a.Models.Users.Insert(u)
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		fmt.Fprintf(w, "%d: %s\n", id, u.FirstName)
	})

	mux.Get("/get-all-users", func(w http.ResponseWriter, r *http.Request) {
		users, err := a.Models.Users.GetAll()
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		for _, x := range users {
			fmt.Fprint(w, x.LastName)
		}
	})

	mux.Get("/get-user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		u, err := a.Models.Users.Get(id)
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		fmt.Fprintf(w, "%s %s %s", u.FirstName, u.LastName, u.Email)
	})   

	mux.Get("/update-user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		u, err := a.Models.Users.Get(id)
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		u.LastName = a.App.RandomString(10)
		err = u.Update(*u)
		if err != nil { 
			a.App.ErrorLog.Println(err)
			return
		}

		fmt.Fprintf(w, "updated last name to %s", u.LastName)

	})

	// static routes
	fileserver := http.FileServer(http.Dir("./public"))
	// a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileserver))
	mux.Handle("/public/*", http.StripPrefix("/public", fileserver))

	// return a.App.Routes
	return mux
}
