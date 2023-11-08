FROM golang:1.21.4 AS gobuilder

WORKDIR /app

COPY . ./
RUN go build -o freecaster cmd/main.go

#CMD ["go", "run", "cmd/main.go"]

FROM ubuntu:23.10

COPY --from=gobuilder /app/freecaster ./
CMD ["./freecaster"]