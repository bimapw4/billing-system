package loan

import (
	"billing/internal/business"
	"billing/internal/common"
	"billing/internal/entity"
	"billing/internal/presentations"
	"billing/internal/response"
	"billing/pkg/meta"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	CreateHandler(c *fiber.Ctx) error
	ListHandler(c *fiber.Ctx) error
}

type handler struct {
	business business.Business
}

func NewLoanHandler(business business.Business) Handler {
	return &handler{
		business: business,
	}
}

func (h *handler) CreateHandler(c *fiber.Ctx) error {

	var (
		Entity = "CreateLoan"
	)
	availError := common.DefaultAvailableErrors()
	errorCust := availError.CustomeError(common.AvailableErrors{
		{
			Code:    http.StatusNotFound,
			Err:     presentations.ErrLoanNotFound,
			Message: presentations.ErrLoanNotFound.Error(),
		},
		{
			Code:    http.StatusConflict,
			Err:     presentations.ErrLoanAlreadyExist,
			Message: presentations.ErrLoanAlreadyExist.Error(),
		},
	})

	var payload entity.Loan
	if err := c.BodyParser(&payload); err != nil {
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", err).
			JSON(c, fiber.StatusBadRequest)
	}

	res, err := h.business.Loan.Create(c.UserContext(), &payload)
	if err != nil {
		fmt.Println("err ", err)
		err := errorCust.GetError(err)
		return response.NewResponse(Entity).
			Errors("Failed to create loan", err.Message).
			JSON(c, err.Code)
	}

	return response.NewResponse(Entity).
		Success("Created Loan successfully", res).
		JSON(c, fiber.StatusCreated)
}

func (h *handler) ListHandler(c *fiber.Ctx) error {

	var (
		Entity = "ListLoan"
	)
	availError := common.DefaultAvailableErrors()
	errorCust := availError.CustomeError(common.AvailableErrors{
		{
			Code:    http.StatusNotFound,
			Err:     presentations.ErrLoanNotFound,
			Message: presentations.ErrLoanNotFound.Error(),
		},
		{
			Code:    http.StatusConflict,
			Err:     presentations.ErrLoanAlreadyExist,
			Message: presentations.ErrLoanAlreadyExist.Error(),
		},
	})

	q := c.Queries()

	m := meta.NewParams(q)

	res, err := h.business.Loan.List(c.UserContext(), &m)
	if err != nil {
		fmt.Println("err", err)
		err := errorCust.GetError(err)
		return response.NewResponse(Entity).
			Errors("Failed to list loan", err.Message).
			JSON(c, err.Code)
	}

	return response.NewResponse(Entity).
		SuccessWithMeta("List Loan successfully", res, m).
		JSON(c, fiber.StatusOK)
}
