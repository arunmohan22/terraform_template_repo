// {{cookiecutter.__copyright}}

package drivers

import (
	"context"
	"fmt"

	"{{ cookiecutter.repoURL }}/internal/adapters"
	"github.com/go-chi/chi/v5"
	"github.hpe.com/cloud/go-gadgets/x/config"
	"github.hpe.com/cloud/go-gadgets/x/logging"
	"github.hpe.com/cloud/go-gadgets/x/restserver"
	"github.hpe.com/cloud/go-gadgets/x/restserver/chiserver"
)

func ProvideRESTServer(ctx context.Context, conf config.Config, logger logging.Logger, router chi.Router)(*restserver.Server, error){
	port, err := config.Get[int](conf, adapters.EnvVarToConfigKeys[adapters.EnvVarRestPort])
	if err != nil {
		return nil, fmt.Errorf("getting rest.port from config: %v", err)
	}
	server, err := chiserver.NewServerFromRouter(router, uint16(port))

	if err != nil{
		return nil, fmt.Errorf("provide rest server: %v", err)
	}

	return server, nil
}
