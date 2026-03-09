#!/bin/sh
set -u

# source vault secrets if present before starting the service
for file in $(find /vault/secrets -maxdepth 1 -type f); do
    echo "sourcing file: ${file}"
    source $file
done

# Execute the passed in parameters
$@

# Get exit code of the passed in parameters
mainExit=$?

# Workaround for https://github.com/istio/istio/issues/6324 if a job
# This is never reached if the above doesn't terminate (i.e. not a job)
curl -sfI -X POST http://127.0.0.1:15020/quitquitquit

# Exit with main job's exit code
exit $mainExit
