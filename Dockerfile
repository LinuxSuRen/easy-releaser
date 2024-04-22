ARG REGISTRY=172.11.0.6:30002/
FROM ${REGISTRY}docker.io/library/golang:1.22.2 as builder
ARG GOPROXY=direct
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o bin/releaser .

FROM ${REGISTRY}docker.io/library/alpine:3.18.4
COPY --from=builder /app/bin/releaser /usr/local/bin/releaser

CMD ["releaser"]
