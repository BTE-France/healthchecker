# healthchecker

This is a very simple Golang program that takes a URL at compile time,
and simply performs health check on it when called.

This is intended to add health checks to hardened container images that do not ship
the use-space utililties such as `curl` or `wget`.

Because an example is worth a thousand words:

```shell
╭─user@example ~/healthchecker
╰─$ go build \
    -ldflags="-X 'healthchecker/config.UrlToCheck=https://btefrance.fr'" \
    -o healthcheck \ 
    healthchecker/cmd/healthcheck
╭─user@example ~/healthchecker
╰─$ ./healthcheck; echo "status: $?"
time=2026-01-01T14:48:50.360+01:00 level=INFO msg="healthcheck finished" app=healthchecker urlToCheck=https://btefrance.fr checker=http success=true
status: 0


╭─user@example ~/healthchecker
╰─$ go build \
    -ldflags="-X 'healthchecker/config.UrlToCheck=https://btefrance.fr/invalid/uri'" \
    -o healthcheck \ 
    healthchecker/cmd/healthcheck
╭─user@example ~/healthchecker
╰─$ ./healthcheck; echo "status: $?"
time=2026-01-01T14:51:29.921+01:00 level=ERROR msg="healthcheck finished" app=healthchecker urlToCheck=https://btefrance.fr/invalid/uri checker=http success=false error="server did not return a success: 404 Not Found (404)"
status: 1
```
