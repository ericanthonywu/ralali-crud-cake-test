package Controller

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"ralali-crud-cake-test/Constant"
	"ralali-crud-cake-test/Model"
	"ralali-crud-cake-test/Model/Database"
	"ralali-crud-cake-test/Services"
	"ralali-crud-cake-test/Utils"
	"strconv"
)

type (
	ICakeControllerHandler interface {
		GetCake(c *fiber.Ctx) (err error)
		GetCakeById(c *fiber.Ctx) (err error)
		AddCake(c *fiber.Ctx) (err error)
		UpdateCake(c *fiber.Ctx) (err error)
		DeleteCake(c *fiber.Ctx) (err error)
	}

	CakeControllerHandler struct {
		service Services.ICakeServicesHandler
	}
)

func CakeControllerControllerProvider(service Services.ICakeServicesHandler) *CakeControllerHandler {
	return &CakeControllerHandler{
		service: service,
	}
}

func (h *CakeControllerHandler) GetCake(c *fiber.Ctx) (err error) {
	var (
		pagination = c.QueryInt("pagination", 1)
		serviceErr *Model.ServiceErrorDto
		response   []Database.Cakes
	)

	if pagination < 1 {
		return c.Status(http.StatusBadRequest).JSON(Model.ErrorResponse("pagination must be greater than 0", nil))
	}

	if response, serviceErr = h.service.GetCake(uint(pagination)); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Model.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Model.SuccessResponse("cake data", response))
}

func (h *CakeControllerHandler) GetCakeById(c *fiber.Ctx) (err error) {
	var (
		id         int
		serviceErr *Model.ServiceErrorDto
		response   Database.Cakes
	)

	if id, err = c.ParamsInt("id"); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Model.ErrorResponse("id must be an integer", nil))
	}

	if response, serviceErr = h.service.GetCakeById(uint64(id)); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Model.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Model.SuccessResponse("cake data with id "+strconv.Itoa(id), response))
}

func (h *CakeControllerHandler) AddCake(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Model.ServiceErrorDto
		request    Model.CakeRequestDto
		errs       []Model.ValidationError
	)

	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Model.ErrorResponse(Constant.FailedBindError, err))
	}

	if errs = Utils.Validate(request); errs != nil {
		return c.Status(http.StatusBadRequest).JSON(Model.ErrorResponse(Constant.ValidationError, errs))
	}

	if serviceErr = h.service.AddCake(request); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Model.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusCreated).JSON(Model.SuccessResponse("cake inserted", nil))
}

func (h *CakeControllerHandler) UpdateCake(c *fiber.Ctx) (err error) {
	var (
		id         int
		serviceErr *Model.ServiceErrorDto
		request    Model.CakeRequestDto
		errs       []Model.ValidationError
	)

	if id, err = c.ParamsInt("id"); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Model.ErrorResponse("id must be an integer", nil))
	}

	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Model.ErrorResponse(Constant.FailedBindError, err))
	}

	if errs = Utils.Validate(request); errs != nil {
		return c.Status(http.StatusBadRequest).JSON(Model.ErrorResponse(Constant.ValidationError, errs))
	}

	if serviceErr = h.service.UpdateCake(uint64(id), request); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Model.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Model.SuccessResponse("cake updated", nil))
}

func (h *CakeControllerHandler) DeleteCake(c *fiber.Ctx) (err error) {
	var (
		id         int
		serviceErr *Model.ServiceErrorDto
	)

	if id, err = c.ParamsInt("id"); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Model.ErrorResponse("id must be an integer", nil))
	}

	if serviceErr = h.service.DeleteCake(uint64(id)); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Model.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Model.SuccessResponse("cake deleted", nil))
}
