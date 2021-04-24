package device

import (
	"fmt"

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
func (h *Handler) CheckDevice(d *api.DeviceCheckDetails) error {

	validate = validator.New()
	err := validate.Struct(d)
	if err != nil {
		return err
	}

	return nil
}
