# builder
FROM golang:alpine3.18 as builder

WORKDIR /app

COPY ./ /app

ENV GO111MODULE=auto
ENV GOPROXY=https://goproxy.cn,direct
RUN cd /app && \
	go install github.com/gobuffalo/packr/v2/packr2@v2.8.3 && \
	go mod tidy && \
	packr2 && \
	go build -o ./server-go && \
	chmod +x server-go 

# runner
FROM alpine:latest
ARG VENKIT_SERVICE 
ENV VENKIT_SERVICE=$VENKIT_SERVICE

WORKDIR /app
COPY --from=builder /app/server-go /app/server-go
COPY --from=builder /app/server-go /usr/local/bin/server-go

ENTRYPOINT ["server-go", "-f", "config/config.yaml"]

