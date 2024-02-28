package usercontroller

import (
	"go-fiber-basic/datamodels"
	userservice "go-fiber-basic/services/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type registerUserRequest struct {
	Username        string  `json:"username" validate:"required"`
	Password        string  `json:"password" validate:"required"`
	ConfirmPassword string  `json:"confirm_password" validate:"required,eqfield=Password"`
	Email           string  `json:"email" validate:"required,email"`
	Cash            float64 `json:"cash" validate:"required"`
}

func RegisterUser(c *fiber.Ctx) error {
	req := new(registerUserRequest)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	v := validator.New()

	err = v.Struct(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	user := &datamodels.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Cash:     req.Cash,
	}

	err = userservice.RegisterUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON("Create User Success !")

}
