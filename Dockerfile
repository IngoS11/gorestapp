FROM golang:1.20.3-alpine as build

WORKDIR /src
COPY . .
RUN go build -o /bin/gorestapp .

FROM scratch as bin

WORKDIR /app
COPY --from=build /bin/gorestapp /app

CMD [ "/app/gorestapp" ]
