package borrowers

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

func NewBorrowerHandler(business business.Business) Handler {
	return &handler{
		business: business,
	}
}

func (h *handler) CreateHandler(c *fiber.Ctx) error {

	var (
		Entity = "CreateBorrower"
	)
	availError := common.DefaultAvailableErrors()
	errorCust := availError.CustomeError(common.AvailableErrors{
		{
			Code:    http.StatusNotFound,
			Err:     presentations.ErrBorrowersNotFound,
			Message: presentations.ErrBorrowersNotFound.Error(),
		},
		{
			Code:    http.StatusConflict,
			Err:     presentations.ErrBorrowersAlreadyExist,
			Message: presentations.ErrBorrowersAlreadyExist.Error(),
		},
	})

	var payload entity.Borrower
	if err := c.BodyParser(&payload); err != nil {
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", err).
			JSON(c, fiber.StatusBadRequest)
	}

	res, err := h.business.Borrowers.Create(c.UserContext(), &payload)
	if err != nil {
		fmt.Println("err ", err)
		err := errorCust.GetError(err)
		return response.NewResponse(Entity).
			Errors("Failed to create borrowers", err.Message).
			JSON(c, err.Code)
	}

	return response.NewResponse(Entity).
		Success("Created Borrowers successfully", res).
		JSON(c, fiber.StatusCreated)
}

func (h *handler) ListHandler(c *fiber.Ctx) error {

	var (
		Entity = "ListBorrower"
	)
	availError := common.DefaultAvailableErrors()
	errorCust := availError.CustomeError(common.AvailableErrors{
		{
			Code:    http.StatusNotFound,
			Err:     presentations.ErrBorrowersNotFound,
			Message: presentations.ErrBorrowersNotFound.Error(),
		},
		{
			Code:    http.StatusConflict,
			Err:     presentations.ErrBorrowersAlreadyExist,
			Message: presentations.ErrBorrowersAlreadyExist.Error(),
		},
	})

	q := c.Queries()

	m := meta.NewParams(q)

	res, err := h.business.Borrowers.List(c.UserContext(), &m)
	if err != nil {
		fmt.Println("err", err)
		err := errorCust.GetError(err)
		return response.NewResponse(Entity).
			Errors("Failed to list borrowers", err.Message).
			JSON(c, err.Code)
	}

	return response.NewResponse(Entity).
		SuccessWithMeta("List Borrowers successfully", res, m).
		JSON(c, fiber.StatusOK)
}
