package auth

import (
	"github.com/arvinpaundra/go-eduworld/internal/adapters/request"
	"github.com/arvinpaundra/go-eduworld/internal/controllers/contracts"
	"github.com/arvinpaundra/go-eduworld/internal/services/auth"
	"github.com/arvinpaundra/go-eduworld/pkg/constants"
	"github.com/arvinpaundra/go-eduworld/pkg/helpers"
	"github.com/gofiber/fiber/v2"
)

type authController struct {
	authService auth.AuthService
}

func NewAuthController(authService auth.AuthService) contracts.AuthController {
	return &authController{
		authService: authService,
	}
}

func (a *authController) HandlerRegister(c *fiber.Ctx) error {
	var input request.Register

	_ = c.BodyParser(&input)

	if err := helpers.ValidateRequest(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.BadRequest("invalid request body", err))
	}

	if err := a.authService.Register(c.Context(), &input); err != nil {
		switch err {
		case constants.ErrUsernameAlreadyTaken:
			return c.Status(fiber.StatusBadRequest).JSON(helpers.BadRequest("invalid request body", map[string]string{
				"username": err.Error(),
			}))
		case constants.ErrInterestNotFound:
			return c.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(helpers.InternalServerError(err.Error()))
		}
	}

	return c.Status(fiber.StatusCreated).JSON(helpers.SuccessCreated("success register", nil))
}

func (a *authController) HandlerLogin(c *fiber.Ctx) error {
	var input request.Login

	_ = c.BodyParser(&input)

	if err := helpers.ValidateRequest(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.BadRequest("invalid request body", err))
	}

	// set ip address
	input.IPAddress = c.IP()

	response, err := a.authService.Login(c.Context(), &input)

	if err != nil {
		switch err {
		case constants.ErrUserNotFound:
			return c.Status(fiber.StatusUnprocessableEntity).JSON(helpers.UnprocessableEntity(constants.ErrCredentialInvalid.Error()))
		case constants.ErrPasswordIncorrect:
			return c.Status(fiber.StatusUnprocessableEntity).JSON(helpers.UnprocessableEntity(constants.ErrCredentialInvalid.Error()))
		case constants.ErrMenteeNotFound:
			return c.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(helpers.InternalServerError(err.Error()))
		}
	}

	return c.Status(fiber.StatusOK).JSON(helpers.SuccessOK("login success", response))
}

func (a *authController) HandlerLogout(c *fiber.Ctx) error {
	userId := c.Query("user_id")

	err := a.authService.Logout(c.Context(), userId)

	if err != nil {
		switch err {
		case constants.ErrUserNotFound:
			return c.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(helpers.InternalServerError(err.Error()))
		}
	}

	return c.Status(fiber.StatusOK).JSON(helpers.SuccessOK("logout success", nil))
}
