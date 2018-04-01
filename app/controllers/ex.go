package controllers

import (
	"amai/app/models"

	"github.com/revel/revel"
)

type Ex struct {
	*revel.Controller
}

func (c Ex) Index(name string) revel.Result {
	edad := 19
	return c.Render(edad, name)
}

func (c Ex) Create(name string, mail string) revel.Result {
	user := models.User{Name: name, Mail: mail}
	Gdb.Create(&user)
	return c.RenderJSON(user)
}
