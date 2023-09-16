package handler

import (
	"net/http"

	"github.com/MYK12397/go-messenger/internal/core/domain"
	"github.com/MYK12397/go-messenger/internal/core/services"
	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	srv services.MessengerService
}

func NewHTTPHandler(MessengerService services.MessengerService) *HTTPHandler {
	return &HTTPHandler{
		srv: MessengerService,
	}
}

func (h *HTTPHandler) SaveMessage(c echo.Context) error {
	var message domain.Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Error": err.Error(),
		})

	}

	err := h.srv.SaveMessage(message)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})

	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "New Message Created Successfully",
	})
}

func (h *HTTPHandler) ReadMessage(c echo.Context) error {
	id := c.Param("id")
	message, err := h.srv.ReadMessage(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})

	}
	return c.JSON(http.StatusOK, message)
}

func (h *HTTPHandler) ReadMessages(c echo.Context) error {

	message, err := h.srv.ReadMessages()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": message,
	})

}
func (h *HTTPHandler) DeleteMessage(c echo.Context) error {
	id := c.Param("id")
	err := h.srv.DeleteMessage(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Message deleted successfully",
	})
}
