#!/usr/bin/env bash

set -euo pipefail

if [[ "$#" -lt 1 ]] || [[ "$#" -gt 2 ]]
then
  echo "Usage: $0 <URL to probe> [binary output path]" >&2
  exit 2
fi

rootGoModule="github.com/BTE-France/healthchecker"
urlToCheck="$1"
outputPath="${2:-healthcheck}"

set -x  # Show execution to caller starting there

export CGO_ENABLED=0

# Build the health checker program
go build \
  -trimpath \
  -gcflags "-l -B" \
  -ldflags="-w -s -X '${rootGoModule}/config.UrlToCheck=${urlToCheck}'" \
  -o "${outputPath}" \
  "${rootGoModule}/cmd/healthcheck"

# Execute the built binary to verify its embedded configuration
$(realpath "$outputPath") -checkConfig

