FROM golang:1.21.4 AS gobuilder

WORKDIR /app

COPY . ./
RUN go build -o freecaster cmd/main.go

FROM ubuntu:23.10

# Update SSL certs so that the FreeStuff API cert signer can be trusted
RUN apt-get update
RUN apt-get install ca-certificates -y

COPY --from=gobuilder /app/freecaster ./
CMD ["./freecaster"]