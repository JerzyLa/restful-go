# mirror_finder

## Run swagger documentation
```bash
docker run --rm -p 80:8080 -e SWAGGER_JSON=/app/openapi.json -v ${path_to_app}:/app swaggerapi/swagger-ui
```