package handlers

import (
	"net/http"

	"github.com/trqvel/web-calc/backend/internal/services"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	svc services.Service
}

func NewHandler(s services.Service) *Handler {
	return &Handler{svc: s}
}

func (h *Handler) GetCalculations(c echo.Context) error {
	calculations, err := h.svc.ReadAllCalculations()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get calculations"})
	}

	return c.JSON(http.StatusOK, calculations)
}

func (h *Handler) PostCalculations(c echo.Context) error {
	var r services.Request
	err := c.Bind(&r)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	calc, err := h.svc.CreateCalculation(r.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not create calculation"})
	}

	return c.JSON(http.StatusCreated, calc)
}

func (h *Handler) PatchCalculations(c echo.Context) error {
	id := c.Param("id")

	var r services.Request
	err := c.Bind(&r)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	updatedCalc, err := h.svc.UpdateCalculation(id, r.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid expression"})
	}

	return c.JSON(http.StatusOK, updatedCalc)
}

func (h *Handler) DeleteCalculations(c echo.Context) error {
	id := c.Param("id")

	if err := h.svc.DeleteCalculation(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete calculation"})
	}

	return c.NoContent(http.StatusNoContent)
}
