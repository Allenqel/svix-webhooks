package svix

import (
	"github.com/svix/svix-libs/go/internal/openapi"
)

// Error provides access to the body, status, and error on returned errors.
type Error struct {
	status int
	body   []byte
	error  string
}

// Error returns non-empty string if there was an error.
func (e Error) Error() string {
	return e.error
}

// Body returns the raw bytes of the response.
func (e Error) Body() []byte {
	return e.body
}

// Status returns the HTTP status of the error.
func (e Error) Status() int {
	return e.status
}

// a simple function convert openapi errors to exposed svix.Error
func wrapError(err error, status int) error {
	if openapiError, ok := err.(openapi.GenericOpenAPIError); ok {
		return &Error{
			status: status,
			body:   openapiError.Body(),
			error:  openapiError.Error(),
		}
	}
	return err
}