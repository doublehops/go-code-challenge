package device

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"codechallenge.local/internal/pkg/types/api"
)

var validate *validator.Validate

type Handler struct {
	// Add a logger here.
	// Add db handler
}

// New Instantiate new handler instance.
func New() *Handler {

	fmt.Println("Setting up device handler")

	return &Handler{}
}

// CheckDevice - Handler for POST payload
func (h *Handler) CheckDevice(c *gin.Context) error {

	var device api.Device
	if err := c.ShouldBindJSON(&device); err != nil {
		return fmt.Errorf("could not parse payload: %s", err.Error())
	}

	validate = validator.New()
	err := validate.Struct(device)
	if err != nil {
		return err
	}

	return nil
}
