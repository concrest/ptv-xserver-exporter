# Test Data Server

Build a docker image to host the JSON files:

```
$ docker build -t fake-ptv-server .
```

Run a container, exposing port 5000:

```
$ docker run --rm -d -p 5000:80 fake-ptv-server
```

Then run the exporter with environment variable METRICS_API_URL=http://localhost:5000/the-json-filename-here (see start-debug.bat for example)