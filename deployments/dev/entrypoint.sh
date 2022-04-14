#!/bin/sh
set -e

cd /go/src/sb

echo "Going to run environment ${ENVIRONMENT}"

if [ "$ENVIRONMENT" = "test" ]; then
    # shellcheck disable=SC2046
    go test $(go list ./... | grep -v /vendor/) -v
else
    reflex -v --start-service --regex='(\.go$|go\.mod|\.js$|\.html$)' -- sh -c 'go run /go/src/sb/cmd/app'
fi