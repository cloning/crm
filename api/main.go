package main

import (
	"github.com/cloning/crm/core"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {

	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Charset: "UTF-8",
	}))

	m.Post("/auth/login", binding.Json(AuthRequest{}), func(request AuthRequest, r render.Render) {
		r.JSON(200, core.Auth_Login(request.Email, request.Password))
	})

	m.Run()
}
