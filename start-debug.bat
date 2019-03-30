@echo off
SET LOG_LEVEL=Debug
SET PORT=9562
SET INCLUDE_DEBUG_HANDLERS=1
SET HTTP_LOGGING_ENABLED=1
SET METRICS_API_URL=http://localhost:5000/example-xlocate.json

ptv-xserver-exporter.exe