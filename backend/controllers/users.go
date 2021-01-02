package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/conduit-golang/backend/commands"
	"github.com/conduit-golang/backend/helpers"
	"github.com/conduit-golang/backend/model"
	"github.com/conduit-golang/backend/services"
)

// Register a new user
func Register(ctx *fiber.Ctx) error {
	var body commands.Register
	helpers.ParseBody(ctx, &body)
	bodyUser := body.User
	user := model.User{
		Username:     bodyUser.Username,
		Email:        bodyUser.Email,
		PasswordHash: bodyUser.Password,
	}
	fmt.Printf("%+v\n", user)
	services.CreateUser(user)
	return ctx.Status(fiber.StatusCreated).JSON(body)
}

// Login a user
func Login(ctx *fiber.Ctx) error {
	type request struct {
		Name string `json:"name"`
	}
	var body request
	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse json",
		})
		return err
	}
	user := model.User{}
	return ctx.Status(fiber.StatusCreated).JSON(user)
}
