@echo off
SET LOG_LEVEL=Debug
SET PORT=8080
SET INCLUDE_DEBUG_HANDLERS=1
SET HTTP_LOGGING_ENABLED=1
SET METRICS_API_URL=http://server:50020/xlocate/some/path

ptv-xserver-exporter.exe