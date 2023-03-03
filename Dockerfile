FROM golang:alpine AS build
WORKDIR /go/src/app
COPY . .
RUN go build -o /go/bin/app cmd/api/main.go

FROM scratch
COPY --from=build /go/bin/app /go/bin/app
EXPOSE 3000
ENTRYPOINT ["/go/bin/app"]