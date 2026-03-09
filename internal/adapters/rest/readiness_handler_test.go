// (C) Copyright 2023 Hewlett Packard Enterprise Development LP

package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	mockReadinessCheckPath = "/readyz"
)

func Test_NewReadinessCheck(t *testing.T) {
	// no arrange necessary

	// act
	handler := NewReadinessCheck()

	// assert
	assert.NotNil(t, handler)
}

func Test_ReadinessCheck_ServeHTTP(t *testing.T) {
	// arrange
	handler := NewReadinessCheck()

	writer := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, mockReadinessCheckPath, http.NoBody)

	// act
	handler.ServeHTTP(writer, request)

	// assert
	response := writer.Result()
	defer response.Body.Close()

	assert.Equal(t, http.StatusNoContent, response.StatusCode)
}
