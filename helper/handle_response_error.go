package helper

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func HandleErrorResponse(c echo.Context, err error) error {

	// Cek error validasi
	if validationErr, ok := err.(*ValidationErrors); ok {

		return c.JSON(validationErr.StatusCode,
			NewErrorResponse(http.StatusBadRequest, validationErr.Messages))
	}

	// Cek error timeout
	if err == context.DeadlineExceeded {
		return c.JSON(http.StatusRequestTimeout,
			NewErrorResponse(http.StatusRequestTimeout, []string{"request timed out"}))
	}

	if err == gorm.ErrRecordNotFound {
		return c.JSON(http.StatusNotFound,
			NewErrorResponse(http.StatusNotFound, []string{"record not found"}))
	}

	// Untuk error lainnya
	return c.JSON(http.StatusInternalServerError,
		NewErrorResponse(http.StatusInternalServerError, []string{"internal server error"}))
}
