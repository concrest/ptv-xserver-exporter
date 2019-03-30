package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	logLevelEnvVarName           = "LOG_LEVEL"
	portEnvVarName               = "HTTP_PORT"
	metricsAPIURLEnvVarName      = "METRICS_API_URL"
	incDebugHandlersEnvVarName   = "INCLUDE_DEBUG_HANDLERS"
	httpLoggingEnabledEnvVarName = "HTTP_LOGGING_ENABLED"
)

// Environment contains the environment variables for the app
type Environment struct {
	// LogLevel is the logging level to output to console
	// Supported values: Debug, Info, Warn, Error.  Default is Info
	LogLevel string
	// Port is the HTTP port to start the server on.  Default is 9562
	Port string
	// MetricsAPIURL must be set to point to a PTV xServer metrics endpoint
	MetricsAPIURL string

	// IncDebugHandlers specifies whether to add pprof HTTP endpoints.
	// Use 0: false, 1: true. Default is false.
	IncDebugHandlersValue string
	IncDebugHandlers      bool

	// HttpLoggingEnabledValue - 0 or 1 for logging HTTP values
	HTTPLoggingEnabledValue string
	HTTPLoggingEnabled      bool

	// VersionInfo contains the version details set at build time
	VersionInfo *VersionInfo
}

// NewEnvironment creates an Environment pointer
func NewEnvironment(version *VersionInfo) *Environment {
	env := &Environment{
		VersionInfo:        version,
		LogLevel:           "Info",
		Port:               "9562", // Claimed on https://github.com/prometheus/prometheus/wiki/Default-port-allocations
		IncDebugHandlers:   false,
		HTTPLoggingEnabled: false,
	}

	if level, hasLevel := os.LookupEnv(logLevelEnvVarName); hasLevel {
		env.LogLevel = level
	}

	if port, hasPort := os.LookupEnv(portEnvVarName); hasPort {
		env.Port = port
	}

	if incDebugHandlers, hasIncDebug := os.LookupEnv(incDebugHandlersEnvVarName); hasIncDebug {
		env.IncDebugHandlersValue = incDebugHandlers
		env.IncDebugHandlers = incDebugHandlers == "1"
	}

	if httpLogging, hasHTTPLogging := os.LookupEnv(httpLoggingEnabledEnvVarName); hasHTTPLogging {
		env.HTTPLoggingEnabledValue = httpLogging
		env.HTTPLoggingEnabled = httpLogging == "1"
	}

	env.MetricsAPIURL = os.Getenv(metricsAPIURLEnvVarName)

	return env
}

// LogVariables dumps the current variables to log
func (e *Environment) LogVariables() {
	log.WithFields(log.Fields{
		logLevelEnvVarName:           e.LogLevel,
		portEnvVarName:               e.Port,
		metricsAPIURLEnvVarName:      e.MetricsAPIURL,
		incDebugHandlersEnvVarName:   e.IncDebugHandlersValue,
		httpLoggingEnabledEnvVarName: e.HTTPLoggingEnabledValue,
		"VersionInfo":                e.VersionInfo,
	}).Info("Environment variables")
}
