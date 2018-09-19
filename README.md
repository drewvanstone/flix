#  Flix
This repository serves to be a reference project for Go API services.

## Run
```bash
docker-compose -f build/docker-compose.yml up --build --force-recreate
go build -o flix ./cmd/flix/main.go && ./flix
```

## Example Usage
```bash
# Create
curl -X PUT localhost:8080/movie?title="Blade"

# Delete
curl -X DELETE localhost:8080/movie?id=1

# Get single movie
curl -s -X GET localhost:8080/movie?id=1 | jq .

# Get all movies
curl -s -X GET localhost:8080/movies | jq .
```

