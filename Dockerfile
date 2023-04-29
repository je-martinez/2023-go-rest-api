FROM golang:alpine AS base

WORKDIR /app
COPY . .

RUN go build -o ./bin/main .

FROM alpine AS final

WORKDIR /app
COPY --from=base /app/bin ./bin
COPY --from=base /app/config/config.yml ./config/

CMD [ "./bin/main" ]