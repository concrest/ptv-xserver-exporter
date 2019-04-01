# PTV xServer Prometheus Exporter

Prometheus Exporter for PTV's xServer suite of GIS services. Converts metrics exposed as JSON by each service to Prometheus format.  Intent is to support the following services:

  * xMap
  * xLocate
  * xRoute
  * xMapMatch

Other PTV xServer products may work if the JSON format is the same.

Project Status: *Alpha* - Minimal functionality, minimal build and release process

## Planned Roadmap

  * Alpha - Minimal functionality, minimal build and release process
  * Beta - Feature complete, build and release process complete
  * Final - Major version 1, ready for production

## Configuration

The only required configuration setting is the PTV xServer metrics API you want to monitor. The exporter only supports 1 metrics API, so when you wish to monitor multiple xServer instances you will need an equal number of exporters.  e.g. if you have an xMap and an xLocate server, you will need 2 exporters - one configured to connect to the xMap metrics API and one to connect to the xLocate metrics API.

See [Administrator Guide - Surveillance and Monitoring](https://xserver.ptvgroup.com/fileadmin/files/PTV-COMPONENTS/DeveloperZone/Documents/xServer_public/manual/Default.htm#Administrators_Guide/DSC_SurveillanceAndMonitoring.htm%3FTocPath%3DAdministrator's%2520Guide%7CAdministration%7C_____3) for details on PTV monitoring.  This exporter works with the API described in the *Use the status report* section:

```
For automated monitoring there are status reports available under //<servername>:<port>/<service>/pages/moduleCommand.jsp?status=json which can be automatically retrieved and parsed.
```

This exporter calls this endpoint, parses the JSON, and exposes the same data as Prometheus data types each time Prometheus scrapes the /metrics endpoint.

Examples:

  * xMap server on hosted on a server call MyMapServer and port 50010: http://MyMapServer:50010/xmap/pages/moduleCommand.jsp?status=json
  * xRoute server on hosted on a server call MyRouteServer and port 50030: http://MyRouteServer:50030/xroute/pages/moduleCommand.jsp?status=json

Set environment variable *METRICS_API_URL* to one of these values.  The exporter will call this URL each time your Prometheus server scrapes the /metrics endpoint.

### Prometheus Setup

The exporter runs on port 9562 by default with a standard /metrics endpoint.  Add this scrape target to your Prometheus setup either directly to the prometheus.yml file, or by using your existing service discovery mechanisms.

See [Prometheus Configuration Guide](https://prometheus.io/docs/prometheus/latest/configuration/configuration/) for more information.

### Environment Variables

Check env.go for full details.  Summary:

  * METRICS_API_URL - Mandatory - PTV xServer metrics API (1 API only). Example: http://servername:50010/xmap/pages/moduleCommand.jsp?status=json
  * LOG_LEVEL - Supported values: Debug, Info, Warn, Error.  Default is Info
  * PORT - Default is 9562
  * INCLUDE_DEBUG_HANDLERS - 0 or 1. Whether to add pprof HTTP endpoints. Default is 0
  * HTTP_LOGGING_ENABLED - 0 or 1. Whether to log HTTP calls at Debug level. Default is 0

## Building from source

Using Go version: go1.12.1 windows/amd64.  Currently developed only on Windows 10.

  * build-debug.bat - Builds a local version with hard coded version numbers
  * start-debug.bat - Executes with full debug logging. Edit this file to set your METRICS_API_URL address

Linux and Docker support will be added later.