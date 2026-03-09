// {{cookiecutter.__copyright}}

package rest

import (
	"net/http"
)

// ReadinessCheck is used to perform readiness checks for the REST API.
type ReadinessCheck struct{}

// NewReadinessCheck creates a new ReadinessCheck instance.
func NewReadinessCheck() *ReadinessCheck {
	return &ReadinessCheck{}
}

// ServeHTTP implements the readiness check for the REST API.
func (*ReadinessCheck) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNoContent)
}