FROM golang:1.13.5-alpine AS build
COPY . .
RUN apk add git
RUN go get -v -u github.com/lib/pq && go get -v -u github.com/caarlos0/env && go get -v -u github.com/prometheus/client_golang/prometheus && go get -v -u github.com/prometheus/client_golang/prometheus/promhttp
RUN go build  ./app.go
FROM alpine:3.10.3 as production
COPY --from=build /go/app .
CMD [ "./app" ]
