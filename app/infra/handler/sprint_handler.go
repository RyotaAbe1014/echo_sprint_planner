package handler

import (
	"echo_sprint_planner/app/domains/services"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type SprintCreateRequest struct {
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	UpdatedBy string    `json:"updated_by"`
}

type SprintUpdateRequest struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	UpdatedBy string    `json:"updated_by"`
}

type SprintDeleteRequest struct {
	ID uuid.UUID `json:"id"`
}

type ISprintHandler interface {
	SprintCreate(c echo.Context) error
	SprintList(c echo.Context) error
	SprintUpdate(c echo.Context) error
	SprintDelete(c echo.Context) error
}

type sprintHandler struct {
	ss services.ISprintService
}

// constructorを使用して、controllerの構造体を生成
func NewSprintHandler(ss services.ISprintService) ISprintHandler {
	return &sprintHandler{ss}
}

func (sh *sprintHandler) SprintCreate(c echo.Context) error {
	req := new(SprintCreateRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request format"})
	}
	err := sh.ss.SprintCreate(req.Name, req.StartDate, req.EndDate, req.UpdatedBy)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "success")
}

func (sh *sprintHandler) SprintList(c echo.Context) error {
	sprintList, err := sh.ss.SprintList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, sprintList)
}

func (sh *sprintHandler) SprintUpdate(c echo.Context) error {
	req := new(SprintUpdateRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request format"})
	}
	err := sh.ss.SprintUpdate(req.ID, req.Name, req.StartDate, req.EndDate, req.UpdatedBy)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "success")
}

func (sh *sprintHandler) SprintDelete(c echo.Context) error {
	req := new(SprintDeleteRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request format"})
	}
	err := sh.ss.SprintDelete(req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "success")
}
