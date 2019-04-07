# Test Data Server

For development, use the JSON files in this directory as fake data sources so you don't need access to a real PTV service.

An easy way to do this is to build a local Docker image to host the JSON files.  The `Dockerfile` in this directory uses `nginx:alpine` to create a basic web server with the JSON files as the content.

Create a fake PTV server image:

```console
docker build -t fake-ptv-server .
```

Run a container, exposing port 5000:

```console
docker run --rm -d -p 5000:80 fake-ptv-server
```

Then run the exporter with environment variable METRICS_API_URL=http://localhost:5000/the-json-filename-here (see start-debug.bat for example)