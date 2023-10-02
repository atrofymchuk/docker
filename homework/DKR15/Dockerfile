FROM golang:1.13.5-alpine
COPY . .
RUN apk add git
RUN go get -v -u github.com/lib/pq && go get -v -u github.com/caarlos0/env && go get -v -u github.com/prometheus/client_golang/prometheus && go get -v -u github.com/prometheus/client_golang/prometheus/promhttp
RUN go build  ./app.go
ARG ARG_NAME
ENV ENV_NAME=$ARG_NAME
FROM alpine:3.10.3
ARG ARG_NAME
RUN echo $ARG_NAME > ./file
COPY --from=0 /go/app .
CMD [ "./app" ]
