FROM golang:alpine AS base

WORKDIR /app
COPY . .

RUN go build -o ./cmd/api/server ./cmd/api/main.go

FROM alpine AS final

WORKDIR /app
COPY --from=base /app/cmd/api/server ./cmd/api/
COPY --from=base /app/config/config.yml ./config/

CMD [ "./cmd/api/server" ]