// {{cookiecutter.__copyright}}

package adapters

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.hpe.com/cloud/go-gadgets/x/config"
	mocks "{{cookiecutter.repoURL}}/mocks/adapters"
)

func TestProvideArgs(t *testing.T) {
	args := ProvideArgs()
	assert.IsType(t, Args{}, args)
}

func TestProvideFlags_HappyPath(t *testing.T) {
	args := []string{"--config", "config.yaml"}
	flags, err := ProvideFlags(args)
	assert.NoError(t, err)

	value, err := flags.GetString("config")
	assert.Nil(t, err)
	assert.Equal(t, "config.yaml", value)
}

func TestProvideFlags_NoFlags(t *testing.T) {
	args := []string{}
	flags, err := ProvideFlags(args)

	assert.NoError(t, err)
	value, _ := flags.GetString("config")
	assert.Equal(t, "", value)
}

func TestProvideFlags_UnknownFlag(t *testing.T) {
	args := []string{"--unknown", "config.yaml"}
	flags, err := ProvideFlags(args)

	assert.Error(t, err)
	assert.Equal(t, "parsing flags: unknown flag: --unknown", err.Error())
	assert.Nil(t, flags)
}

func TestProvideConfig_NoConfigFile(t *testing.T) {
	flags := &mocks.Flags{}

	flags.On("GetString", "config").Return("", nil)
	koanfConfig, err := ProvideConfig(flags)
	assert.NoError(t, err)
	assert.IsType(t, &config.KoanfConfig{}, koanfConfig)
}

func TestProvideConfig_UnknownFlag(t *testing.T) {
	flags := &mocks.Flags{}

	flags.On("GetString", "config").Return("", errors.New("unknown flag"))
	koanfConfig, err := ProvideConfig(flags)
	assert.Error(t, err)
	assert.Nil(t, koanfConfig)
}

func TestEnvVarConverterFromKeyMap(t *testing.T) {
	var testCases = []struct {
		name      string
		envVar    string
		configKey string
	}{
		{
			name:      "loglevel",
			envVar:    EnvVarLogLevel,
			configKey: EnvVarToConfigKeys[EnvVarLogLevel],
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, _ := envVarConverter(tc.envVar, tc.configKey)
			assert.Equal(t, tc.configKey, result)
		})
	}
}

func TestValueConverterValues(t *testing.T) {
	var testCases = []struct {
		name           string
		key            string
		value          string
		convertedValue any
	}{
		{
			name:           "string value",
			key:            "unknown",
			value:          "10",
			convertedValue: "10",
		},
		{%- if cookiecutter.rest == "True" %}
		{
			name:           "rest server port",
			key:            EnvVarToConfigKeys[EnvVarRestPort],
			value:          "8080",
			convertedValue: int64(8080),
		},
		{
			name:           "rest server port error",
			key:            EnvVarToConfigKeys[EnvVarRestPort],
			value:          "int",
			convertedValue: 0,
		},
		{%- endif %}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := valueConverter(tc.key, tc.value)
			assert.Equal(t, tc.convertedValue, result)
		})
	}
}

