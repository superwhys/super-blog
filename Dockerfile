FROM bitnami/git as src
ARG REPO=https://superwhys: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@github.com/superwhys/superBlog.git
ARG BRANCH=main
WORKDIR /
RUN git clone --depth=1 -b $BRANCH $REPO app

# builder
FROM golang:alpine3.18 as builder

WORKDIR /app

COPY --from=src /app /app

ARG SERVICE_NAME

ENV GO111MODULE=auto
ENV GOPROXY=https://goproxy.cn,direct
RUN cd /app && \
	go install github.com/gobuffalo/packr/v2/packr2@v2.8.3 && \
	go mod tidy && \
	packr2 && \
	packr2 clean && \
	go build -o ./server-go && \
	chmod +x server-go 

# runner
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/server-go /app/server-go
COPY ./config.yaml /app/config.yaml
COPY ./blog-posts /app/blog-posts

ENTRYPOINT ["/app/server-go", "-f", "/app/config.yaml"]

