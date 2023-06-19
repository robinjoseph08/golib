# golib

[![Version](https://img.shields.io/badge/version-v0.1.0-green.svg)](https://github.com/robinjoseph08/golib/tags)
[![GoDoc](https://pkg.go.dev/badge/github.com/robinjoseph08/golib)](https://pkg.go.dev/github.com/robinjoseph08/golib)

This is a Go module of some of the packages that I use in almost every one of my Go projects, so instead of copying them
over from project to project, I decided to just put them up here so that I can reference them directly. That way, any
changes I need to make, I can just make it in one place.

## Packages

The included packages are:

- [`errutils`](./errutils) - A package to provide some utilities when dealing with errors.
- [`logger`](./logger) - A standard logger that uses [zerolog](https://github.com/rs/zerolog) under the hood, but
  provides an interface that is a bit more ergonomic.
- [`signals`](./signals) - A standard way to capture shutdown signals, so that you can gracefully shut your application
  down.
- [Echo](https://echo.labstack.com/) v4 packages (v3 is also supported by removing `v4` from the path)
  - [`echo/v4/health`](./echo/v4/health) - A basic health check endpoint for Echo servers.
  - [`echo/v4/middleware/logger`](./echo/v4/middleware/logger) - A logging middleware that logs requests for Echo
    servers.
  - [`echo/v4/middleware/recovery`](./echo/v4/middleware/recovery) - A recovery middleware to save from panics for Echo
    servers.
  - [`echo/v4/test`](./echo/v4/test) - A package that could be useful for testing Echo servers.

## Release

To release a new version, run the following (replacing `v1.0.0` with the appropriate new version number):

```sh
make release tag=v1.0.0
```
