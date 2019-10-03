FROM golang:1.7.4-alpine3.5

WORKDIR /src

COPY src/golang-ci/main.go ./
COPY src/golang-ci/main_test.go ./

EXPOSE 8080
CMD ["go", "run", "main.go"]
CMD ["go", "test"]