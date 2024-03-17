package controller

import (
	"errors"
	"service-user/helpers"
	"service-user/model"

	"service-user/config"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type WebResponse struct {
	Code   int
	Status string
	Data   interface{}
}

func Register(c *fiber.Ctx) error {
	var requestBody model.User
	pool := config.GetPostgresDatabase()

	conn, err := pool.Acquire(c.Context())
	if err != nil {
		return err
	}

	defer conn.Release()

	requestBody.Id = uuid.New().String()

	ctx, cancel := config.NewPostgresContext()
	defer cancel()

	c.BodyParser(&requestBody)

	_, err = conn.Exec(ctx, "INSERT INTO service_users (id, email, password) VALUES ($1, $2, $3)",
		requestBody.Id, requestBody.Email, helpers.HashPassword([]byte(requestBody.Password)))
	if err != nil {
		return err
	}

	return c.JSON(WebResponse{
		Code:   201,
		Status: "OK",
		Data:   requestBody.Email,
	})
}

func Login(c *fiber.Ctx) error {
	var requestBody model.User
	var result model.User

	pool := config.GetPostgresDatabase()
	conn, err := pool.Acquire(c.Context())
	if err != nil {
		return err
	}

	defer conn.Release()

	c.BodyParser(&requestBody)

	ERR := pool.QueryRow(c.Context(), "SELECT email, password FROM service_users WHERE email = $1", requestBody.Email).
		Scan(&result.Email, &result.Password)
	if ERR != nil {
		return c.JSON(WebResponse{
			Code:   401,
			Status: "BAD_REQUEST",
			Data:   ERR.Error(),
		})
	}

	checkPassword := helpers.ComparePassword([]byte(result.Password), []byte(requestBody.Password))
	if !checkPassword {
		return c.JSON(WebResponse{
			Code:   401,
			Status: "BAD_REQUEST",
			Data:   errors.New("invalid password").Error(),
		})
	}

	access_token := helpers.SignToken(requestBody.Email)

	return c.JSON(struct {
		Code        int
		Status      string
		AccessToken string
		Data        interface{}
	}{
		Code:        200,
		Status:      "OK",
		AccessToken: access_token,
		Data:        result,
	})
}

func Auth(c *fiber.Ctx) error {
	return c.JSON("OK")
}
