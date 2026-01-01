package main

import (
	"flag"
	"fmt"
	"healthchecker/config"
	"healthchecker/internal"
	"log/slog"
	"os"
	"strings"
)

type cliArgs struct {
	doConfigCheck bool
	logLevel      slog.Level
	logHandler    slog.Handler
}

func parseCliArgs() (*cliArgs, error) {
	args := &cliArgs{}

	var logLevelArg string
	var logFormatArg string
	flag.BoolVar(&args.doConfigCheck, "checkConfig", false, "If provided, checks the configuration and exits, instead of performing the actual check")
	flag.StringVar(&logLevelArg, "logLevel", "info", "The log level")
	flag.StringVar(&logFormatArg, "logFormat", "text", "The log format")
	flag.Parse()

	var firstErr error = nil
	setErr := func(err error) {
		if firstErr == nil {
			firstErr = err
		}
	}

	var logLevel slog.LevelVar
	if err := logLevel.UnmarshalText(([]byte)(logLevelArg)); err != nil {
		setErr(err)
	}
	logHandlerOption := slog.HandlerOptions{
		Level: &logLevel,
	}

	switch strings.ToLower(logFormatArg) {
	case "json":
		args.logHandler = slog.NewJSONHandler(os.Stdout, &logHandlerOption)
	case "text":
		args.logHandler = slog.NewTextHandler(os.Stdout, &logHandlerOption)
	default:
		setErr(fmt.Errorf("unknown log format: %s", logFormatArg))
		args.logHandler = slog.NewTextHandler(os.Stdout, &logHandlerOption)
	}

	return args, firstErr
}

func main() {

	args, err := parseCliArgs()
	logger := slog.New(args.logHandler).With(
		"app", "healthchecker",
		"urlToCheck", config.UrlToCheck,
	)
	slog.SetDefault(logger)
	slog.SetLogLoggerLevel(args.logLevel)
	if err != nil {
		logger.Error("invalid CLI arguments", "error", err)
		os.Exit(2)
	}

	healthchecker, err := internal.GetHealthCheckerForUrl(config.UrlToCheck)
	if err != nil {
		logger.Error("invalid configuration", "error", err)
		os.Exit(3)
	}

	logger = logger.With("checker", healthchecker.Name())

	if args.doConfigCheck {
		logger.Info("configuration is valid")
		os.Exit(0)
	}

	logger.Debug("healthcheck starting")
	err = healthchecker.Check()
	if err != nil {
		logger.Error(
			"healthcheck finished",
			"success", false,
			"error", err,
		)
		os.Exit(1)
	}
	logger.Info(
		"healthcheck finished",
		"success", true,
	)
	os.Exit(0)

}
