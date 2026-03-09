// {{cookiecutter.__copyright}}
package adapters

import (
	"fmt"
	"os"
	{%- if cookiecutter.rest == "True" %}
	"strconv"
	{%- endif %}

	flag "github.com/spf13/pflag"
	"github.hpe.com/cloud/go-gadgets/x/config"
)
// Environment variable names
const (
	EnvVarLogLevel = "APP_LOGGING_LOGLEVEL"
	{%- if cookiecutter.grpc == "True" %}
	EnvVarGrpcPort = "APP_GRPC_PORT"
	{%- endif %}
	{%- if cookiecutter.rest == "True" %}
	EnvVarRestPort = "APP_REST_PORT"
	{%- endif %}
)

const (
	logLevelKey = "logging.loglevel"
	{%- if cookiecutter.grpc == "True" %}
	grpcPortKey = "grpc.port"
	{%- endif %}
	{%- if cookiecutter.rest == "True" %}
	restPortKey = "rest.port"
	{%- endif %}
)

// EnvVarToConfigKeys is the map of environment variable names to their corresponding config keys
var EnvVarToConfigKeys = map[string]string{
	EnvVarLogLevel : logLevelKey,
	{%- if cookiecutter.grpc == "True" %}
	EnvVarGrpcPort : grpcPortKey,
	{%- endif %}
	{%- if cookiecutter.rest == "True" %}
	EnvVarRestPort : restPortKey,
	{%- endif %}
}

type Args []string

// Flags is an interface that allows the mocking of flag.FlagSet
type Flags interface {
	GetString(name string) (string, error)
}

func ProvideArgs() Args {
	return os.Args[1:]
}

// ProvideFlags returns a Flags that contains all the Flags required for the program.
func ProvideFlags(args Args) (Flags, error) {
	f := flag.NewFlagSet("config", flag.ContinueOnError)
	f.String("config", "", "path to config file")
	if err := f.Parse(args); err != nil {
		return nil, fmt.Errorf("parsing flags: %w", err)
	}
	return f, nil
}

func ProvideConfig(flags Flags) (*config.KoanfConfig, error) {
	configFilePath, err := flags.GetString("config")
	if err != nil {
		return nil, fmt.Errorf("reading flags for config files: %w", err)
	}

	conf := config.NewKoanfConfig(".")
	if configFilePath != "" {
		conf.WithYAML(configFilePath, false)
	}

	return conf.
		WithEnvVars("APP_", envVarConverter).
		TreatInt64AsInt(true).
		Load()
}

// envVarConverter defines a env var conversion to normalise our env var keys
var envVarConverter = config.EnvConverterFromKeyMap(
	EnvVarToConfigKeys,
	valueConverter,
)

// valueConverter defines a value converter function to ensure that any config
// values are returned as the correct type, even when coming
// from a config map file or env var
func valueConverter(key, value string) any {
	switch key {
	{%- if cookiecutter.rest == "True" %}
	case EnvVarToConfigKeys[EnvVarRestPort]:
		intVal, err := strconv.ParseInt(value, 10, 0)
		if err != nil {
			return 0
		}
		return intVal
	{%- endif %}
	default:
		return value
	}
}
