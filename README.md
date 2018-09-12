#  Flix
This repository serves to be a reference project for Go API services.

# Docker stack
`docker stack` is a built-in docker command that deploys a 'stack', or group of containers that operate together. For local development, this can be very useful.

## Build Postgres Container
```bash
cd ./build
docker build -t flix-db:1.0.0 -f Dockerfile.db .
```

## Build
```bash
cd ./build
docker stack deploy flix -c stack.yml
psql -h localhost -p 5432 -U postgres -f schema.sql flix_development
```

## Run
```bash
go build -o flix ./cmd/flix/main.go && ./flix
```

## Shutdown
```bash
cd ./build
docker stack rm flix
```

## Example Usage
```bash
# Create
curl -X PUT localhost:8100/movie?title="Blade"

# Read
curl localhost:8100/movie?id=1

# Delete
curl -X DELETE localhost:8100/movie?id=1
```

