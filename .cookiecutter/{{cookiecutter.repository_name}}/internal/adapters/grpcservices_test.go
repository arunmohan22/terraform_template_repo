// {{cookiecutter.__copyright}}

package adapters

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewGRPCServices_No_Logger(t *testing.T) {
	services,err := NewGRPCServices(nil)
	require.Error(t, err, "logger not provided")
	require.Nil(t,services)
}
