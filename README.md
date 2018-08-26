#  Minerva
This repository serves to be a reference project for Go API services.

# Docker stack
`docker stack` is a built-in docker command that deploys a 'stack', or group of containers that operate together. For local development, this can be very useful.

## Build
```bash
cd ./build
docker stack deploy minerva -c stack.ym
psql -h localhost -p 5432 -U postgres -f schema.sql minerva_development
```

## Run
```bash
go build -o minerva ./cmd/minerva/main.go && ./minerva
```

## Shutdown
```bash
cd ./build
docker stack rm minerva
```

## Example Usage
```bash
# Create
curl -X PUT localhost:8100/task?task="MyNewTask"

# Read
curl localhost:8100/task?id=1

# Delete
curl -X DELETE localhost:8100/task?id=1
```

