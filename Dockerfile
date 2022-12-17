FROM golang:1.19 as build
WORKDIR /webserver
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o webserver

FROM scratch
COPY --from=build /webserver/webserver .
ENTRYPOINT ["./webserver"]