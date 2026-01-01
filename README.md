# healthchecker

This is a very simple Golang program that takes a URL at compile time,
and simply performs health check on it when called.

This is intended to add health checks to hardened container images that do not ship
the use-space utililties such as `curl` or `wget`.

Because an example is worth a thousand words:

```shell
╭─user@example ~/healthchecker
╰─$ ─$ ./build.sh https://btefrance.fr
+ export CGO_ENABLED=0
+ CGO_ENABLED=0
+ go build -trimpath -gcflags '-l -B' '-ldflags=-w -s -X '\''github.com/BTE-France/healthchecker/config.UrlToCheck=https://btefrance.fr'\''' -o healthcheck github.com/BTE-France/healthchecker/cmd/healthcheck
++ realpath healthcheck
+ /home/smyler/Code/BteFrance/healthchecker/healthcheck -checkConfig
time=2026-01-01T16:41:26.744+01:00 level=INFO msg="configuration is valid" app=healthchecker urlToCheck=https://btefrance.fr checker=http
╭─user@example ~/healthchecker
╰─$ ./healthcheck; echo "exit code: $?"
time=2026-01-01T16:42:30.514+01:00 level=INFO msg="healthcheck finished" app=healthchecker urlToCheck=https://btefrance.fr checker=http success=true
exit code: 0


╭─user@example ~/healthchecker
╰─$ ./build.sh https://btefrance.fr/invalid/uri
+ export CGO_ENABLED=0
+ CGO_ENABLED=0
+ go build -trimpath -gcflags '-l -B' '-ldflags=-w -s -X '\''github.com/BTE-France/healthchecker/config.UrlToCheck=https://btefrance.fr/invalid/uri'\''' -o healthcheck github.com/BTE-France/healthchecker/cmd/healthcheck
++ realpath healthcheck
+ /home/smyler/Code/BteFrance/healthchecker/healthcheck -checkConfig
time=2026-01-01T16:44:28.568+01:00 level=INFO msg="configuration is valid" app=healthchecker urlToCheck=https://btefrance.fr/invalid/uri checker=htt
╭─user@example ~/healthchecker
╰─$ ./healthcheck; echo "exit code: $?"
time=2026-01-01T16:44:33.023+01:00 level=ERROR msg="healthcheck finished" app=healthchecker urlToCheck=https://btefrance.fr/invalid/uri checker=http success=false error="server did not return a success: 404 Not Found (404)"
exit code: 1
```
